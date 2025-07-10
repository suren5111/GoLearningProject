一个基于Golang + Gin + GORM + Redis的学生信息管理工具

## 版本信息
Golang: 1.24.5

Redis: 7.4.2

MySQL: 9.2.0

## 项目目录结构
StudentService/

├── main.go

├── go.mod

├── go.sum

├── config/

│   └── config.go

├── controller/

│   └── student_controller.go

├── model/

│   └── student.go

├── service/

│   └── student_service.go

├── dao/

│   └── student_dao.go

├── cache/

│   └── redis_cache.go

├── database/

│   └── db.go

└── test/

    └── load_test.go

## MySQL表
| 字段名 | 类型 | 说明 |
| --- | --- | --- |
| `id` | VARCHAR(255) | 主键 |
| `name` | VARCHAR(255) | 学生姓名 |
| `age` | INT | 年龄 |
| `grade` | VARCHAR(255) | 年级 |


## 开始
安装依赖

```java
go mod tidy
```

修改数据库配置文件config.go建表后，启动项目

```java
go run main.go
```

服务将运行在http://localhost:8080

