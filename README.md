# 一个基于Golang + Gin + GORM + Redis的学生信息管理工具，通过Jwt实现接口鉴权，并使用Redis实现缓存。
测试接口时请先登录获取token，登录接口为/login,
POST http:······/login
{ "username"= "任意", "password"= "任意"}
请求头为Authorization: {token}

## 版本信息
Golang: 1.24.5

Redis: 7.4.2

MySQL: 9.2.0

## 项目目录结构
- `cache/`      // 缓存
    - `redis_cache.go`
- `config/`    // 配置
    - `config.go`
- `controller/`    // 控制器
    - `auth_controller.go`
    - `student_controller.go`
- `dao/`    // 数据访问
    - `student_dao.go`
- `data/`    //  数据
    - `student.json`
- `database/`    // 数据库
    - `db.go`
- `middleware/`    // 中间件
    - `auth.go`
- `model/`    // 类型
    - `student.go`
- `service/`    //  服务
    - `student_controller.go`
- `test/`    // 测试
    - `load_test.go`
- `utils/`    // 工具
    - `jwt.go`
- `go.mod`    // 依赖
- `go.sum`    
- `main.go`    // 入口
- `README.md`   
- `test2`    // Linux平台可执行文件

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

