# Kitex-dousheng

本项目的一些素材来源于字节跳动后端训练营，他们提供了一个抖声客户端，并且规定了可以实现的一些接口，包括用户登陆与注册、视频流、评论功能、点赞功能、关注功能等等。在此之前，我们曾经实现过mvc单体架构的后端项目，可以查看该仓库https://github.com/Wuhlan3/dousheng
之后我们打算将该项目改造为微服务架构，从而提高其负载均衡的能力，降低模块与模块之间的耦合性。最终，使用了字节跳动自主研发的kitex框架，结合Go语言HTTP框架Gin和Go语言ORM框架Gorm。

[![Go](https://img.shields.io/badge/Go-v1.16-blue)](https://go.dev/)
[![Kitex](https://img.shields.io/badge/Kitex-v0.3.0-green)](https://github.com/cloudwego/kitex)
[![Gin](https://img.shields.io/badge/Gin-v1.8.1-brightgreen)](https://github.com/gin-gonic/gin)
[![Gorm](https://img.shields.io/badge/Gorm-v1.23.8-blue)](https://gorm.io/)
[![ETCD](https://img.shields.io/badge/ETCD-v0.0.0-orange)](https://github.com/kitex-contrib/registry-etcd)
[![Jaeger](https://img.shields.io/badge/Jaeger-v2.30.0-blue)](https://github.com/jaegertracing/jaeger-client-go)
[![JWT](https://img.shields.io/badge/JWT-v3.2.0-green)](https://github.com/golang-jwt/jwt)
[![ffmpeg](https://img.shields.io/badge/ffmpeg-v3.1-orange)](https://github.com/u2takey/ffmpeg-go)
[![MySQL](https://img.shields.io/badge/MySQL-v5.7-green)](https://www.mysql.com/)
[![Redis](https://img.shields.io/badge/Redis-v6.2.6-orange)](https://redis.io/)


## 一、项目架构如下

<img src="https://wuhlan3-1307602190.cos.ap-guangzhou.myqcloud.com/img/kitexdousheng08092316.jpg" width="450px">

## 二、数据库E-R图

<img src="https://wuhlan3-1307602190.cos.ap-guangzhou.myqcloud.com/img/er.jpg" width="800px">

## 三、feed过程
feed即用户在刷视频过程中请求的接口，响应的是视频相关数据，这一部分应该是最频繁调用的且包括了几乎所有表的数据，所以该过程较复杂。
1. 用户会请求两个参数，分别是token和latest_time。其中token会经过JWT解析，得到用户的uid，latest_time表示限制返回视频的时间戳；
2. 由于需要限制返回的视频数量，且我们期望能够优先刷到最新投稿的视频，所以可以采用Redis中的ZSET数据结构来保存视频的序列号；
3. 为了减少视频信息的查询数据库次数，当我们获得视频序列号的时候，可以直接通过video:id在Redis中查询相应的视频信息。
其流程图如下：

<img src="https://wuhlan3-1307602190.cos.ap-guangzhou.myqcloud.com/img/dousheng_feed.jpg" width="700px">

经过测试，使用Redis后，响应速度从98ms提升到55ms。

## 四、运行方法
```
#运行http接口
cd cmd/api
sh run.sh

#运行user服务
cd cmd/user
sh build.sh
sh output/bootstrap.sh

#运行feed服务
cd cmd/feed
sh build.sh
sh output/bootstrap.sh

#运行publish服务
cd cmd/publish
sh build.sh
sh output/bootstrap.sh

#运行comment服务
cd cmd/comment
sh build.sh
sh output/bootstrap.sh

#运行favorite服务
cd cmd/favorite
sh build.sh
sh output/bootstrap.sh

#运行relation服务
cd cmd/relation
sh build.sh
sh output/bootstrap.sh
```

## 五、实现各模块后分别用Docker部署
以api模块为例：
``` yaml
FROM golang:latest
WORKDIR /kitexdousheng 
ADD . /kitexdousheng
ENV GOPROXY https://goproxy.cn
EXPOSE 8081
CMD go mod tidy
CMD cd /kitexdousheng/cmd/api && sh run.sh
```
- 然后docker build -t api_service . 生成镜像（api_service是镜像名）
- docker run -d --network host api_service 以共享网络方式运行镜像
- 接下来登录Docker hub账号push上去

## 六、运行结果
注册与登录、视频流功能如下：

![dousheng_result1](https://wuhlan3-1307602190.cos.ap-guangzhou.myqcloud.com/img/dousheng1.png)

点赞、关注、喜欢视频列表、评论等功能如下：

![dousheng_result2](https://wuhlan3-1307602190.cos.ap-guangzhou.myqcloud.com/img/dousheng2.png)

## 七、jeager使用方法
在浏览器访问http://127.0.0.1:16686/

![jeager](https://wuhlan3-1307602190.cos.ap-guangzhou.myqcloud.com/img/Snipaste_2022-08-15_11-56-57.jpg)

## 八、测试

### 8.1 单元测试
- 可以参考我们的[Apifox文档](https://www.apifox.cn/apidoc/shared-b0f7a1d0-d9b5-4f01-af65-8876a319fc0b)

- 另外，在每一些小的功能模块中，有相应的test文件。通过`go test .`可以来测试该模块是否正确。

### 8.2 性能测试

- 使用火焰图来查看全局的CPU、内存使用情况。具体过程可以参考[pprof性能分析](https://wuhlan3.gitee.io/wuhlan3/2022/07/31/pprof性能分析/)

- 通过jeager辅助调试和排查性能问题。

- 使用Apifox来进行压力测试。

## 九、项目亮点
- 对密码进行加密，使用jwt鉴权
- 将视频上传到腾讯云cos存储桶，便于管理，提供传输效率
- 使用ffmpeg截取视频的封面
- 使用jaeger进行链路追踪
- 添加cpu限制，增加熔断报警
- 使用Redis缓存，减少访问数据库的次数；

## 十、提高方向
- 考虑负载均衡等问题；
- 参数校验完善；
- 考虑消息队列的使用

## 十一、References
[1] https://www.cloudwego.io/zh/docs/kitex/getting-started/

[2] https://github.com/cloudwego/kitex-examples/tree/main/bizdemo/easy_note

[3] 之前做的MVC结构抖声项目: https://github.com/Wuhlan3/dousheng

