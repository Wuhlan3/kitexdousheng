# Kitex-dousheng
本项目的一些素材来源于字节跳动后端训练营，他们提供了一个抖声客户端，并且规定了可以实现的一些接口，包括用户登陆与注册、视频流、评论功能、点赞功能、关注功能等等。在此之前，我们曾经实现过mvc单体架构的后端项目，可以查看该仓库https://github.com/Wuhlan3/dousheng。
之后我们打算将该项目改造为微服务架构，从而提高其负载均衡的能力，降低模块与模块之间的耦合性。最终，使用了字节跳动自主研发的kitex框架，结合Go语言HTTP框架Gin和Go语言ORM框架Gorm，

## 一、项目架构如下
<img src="https://wuhlan3-1307602190.cos.ap-guangzhou.myqcloud.com/img/kitexdousheng.jpg" width="500px">

## 二、数据库E-R图
<img src="https://wuhlan3-1307602190.cos.ap-guangzhou.myqcloud.com/img/er.jpg" width="800px">

## 三、运行方法
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

## 四、实现各模块后分别用Docker部署
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

## 五、jeager使用方法
在浏览器访问http://127.0.0.1:16686/

## 六、项目亮点
- 对密码进行加密，使用jwt鉴权
- 将视频上传到腾讯云cos存储桶，便于管理，提供传输效率
- 使用ffmpeg截取视频的封面
- 使用jaeger进行链路追踪
- 添加cpu限流器，提高可用性

## 七、提高方向
- 使用Redis缓存，减少数据库操作；
- 考虑负载均衡等问题；
- 参数校验完善；
- 考虑消息队列的使用

## 八、References
[1] https://www.cloudwego.io/zh/docs/kitex/getting-started/

[2] https://github.com/cloudwego/kitex-examples/tree/main/bizdemo/easy_note

[3] 之前做的MVC结构抖声项目: https://github.com/Wuhlan3/dousheng
