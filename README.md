# About

此项目是基于[cladmin-vue](https://github.com/chenglongcl/cladmin-vue)二次开发后台管理前端功能。

## 系统环境

golang语言：go 1.19+

数据库：mysql 5.7

缓存：redis 5.0

# 说明

> 传送门：[前端系统地址](https://github.com/chenglongcl/cladmin-vue)

## 技术栈

gin + gorm/gen + casbin rabc

## 项目运行

1、clone项目源代码。

```
git clone  https://github.com/chenglongcl/cladmin.git
```

2、导入cmd/cladmin/db/mysql.sql数据库文件至你的数据库。

3、复制cmd/cladmin/conf/config_example.yaml配置文件并重命名为config.yaml，修改mysql、redis等配置。

4、启动服务端

```
cd cmd/cladmin
go run main.go
```

5、数据库操作：[Gen Guides | GORM - The fantastic ORM library for Golang, aims to be developer friendly.](https://gorm.io/zh_CN/gen/)