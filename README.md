# Kitex-dousheng
基于kitex+gin+gorm实现的青春版抖音项目

# 项目架构如下
<img src="https://wuhlan3-1307602190.cos.ap-guangzhou.myqcloud.com/img/UML.jpg" width="500px">

# 数据库E-R图
![dousheng](https://wuhlan3-1307602190.cos.ap-guangzhou.myqcloud.com/img/dousheng.svg)

# 运行方法
```
#运行http接口
cd cmd/api
go run main.go router.go
#运行user服务
cd cmd/user
sh build.sh
sh output/bootstrap.sh
```
# References
[1] https://www.cloudwego.io/zh/docs/kitex/getting-started/

[2] https://github.com/cloudwego/kitex-examples/tree/main/bizdemo/easy_note

[3] 之前做的MVC结构抖声项目: https://github.com/Wuhlan3/dousheng
