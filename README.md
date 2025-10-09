# toge

一个基于 Go 和 Gin 框架构建的企业级 RESTful API 服务，采用 Clean Architecture 设计模式，支持完整的用户管理、人员管理、时区处理等功能。

## 🚀 特性

### 核心功能
- **用户管理系统** - 完整的用户注册、登录、认证、CRUD操作
- **人员管理系统** - 人员信息的增删改查、搜索、分页
- **时区处理系统** - 时区转换、时间解析、格式化
- **系统监控** - 健康检查、性能监控、系统状态

### 技术特性
- **Clean Architecture** - 清晰的分层架构设计
- **依赖注入** - 使用 Google Wire 进行依赖管理
- **JWT认证** - 安全的用户认证和授权
- **数据库支持** - MySQL + GORM ORM
- **缓存支持** - Redis 缓存集成
- **日志系统** - 结构化日志、链路追踪
- **API文档** - 自动生成 Swagger 文档
- **容器化** - 完整的 Docker 部署方案
- **多环境** - 支持 dev/test/production 环境

## 📋 目录结构

```
toge-api/
├── cmd/                    # 应用程序入口
│   ├── main.go            # 主程序入口
│   ├── migrate/           # 数据库迁移工具
│   └── generate/          # 代码生成工具
├── config/                # 配置文件
│   ├── dev.yaml          # 开发环境配置
│   ├── test.yaml         # 测试环境配置
│   └── production.yaml   # 生产环境配置
├── internal/              # 内部业务逻辑
│   ├── app/              # 应用程序核心
│   ├── domain/           # 领域层（接口定义）
│   ├── handler/          # HTTP处理器
│   ├── service/          # 业务服务层
│   ├── repository/       # 数据访问层
│   ├── model/            # 数据模型
│   ├── middleware/       # 中间件
│   └── wire/             # 依赖注入配置
├── pkg/                   # 公共包
│   ├── config/           # 配置管理
│   ├── database/         # 数据库连接
│   ├── logger/           # 日志系统
│   ├── jwt/              # JWT认证
│   ├── redis/            # Redis连接
│   ├── response/         # 响应处理
│   ├── pagination/       # 分页处理
│   ├── timezone/         # 时区处理
│   └── password/         # 密码处理
├── docs/                  # API文档
├── Dockerfile            # Docker配置
├── Makefile              # 构建脚本
└── go.mod               # Go模块文件
```

## 🛠️ 技术栈

- **Go 1.25.0** - 编程语言
- **Gin** - Web框架
- **GORM** - ORM框架
- **MySQL** - 数据库
- **Redis** - 缓存
- **JWT** - 认证
- **Wire** - 依赖注入
- **Swagger** - API文档
- **Docker** - 容器化

## 🚀 快速开始

### 环境要求

- Go 1.25.0+
- MySQL 5.7+
- Redis 6.0+
- Docker (可选)

### 安装依赖

```bash
go mod download
```

### 配置数据库

1. 创建 MySQL 数据库
2. 修改 `config/dev.yaml` 中的数据库配置

### 运行数据库迁移

```bash
make migrate-up
```

### 启动服务

#### 开发环境
```bash
# 直接运行
go run cmd/main.go

# 或使用 Docker
make dev
```

#### 生产环境
```bash
make production
```

### 访问服务

- **API服务**: http://localhost:8080
- **API文档**: http://localhost:8080/swagger/index.html
- **健康检查**: http://localhost:8080/health

## 📚 API 文档

### 认证相关

#### 用户注册
```http
POST /auth/register
Content-Type: application/json

{
  "username": "john_doe",
  "email": "john@example.com",
  "password": "123456",
  "nickname": "John Doe"
}
```

#### 用户登录
```http
POST /auth/login
Content-Type: application/json

{
  "username": "john_doe",
  "password": "123456"
}
```

#### 获取用户信息
```http
GET /auth/profile
Authorization: Bearer <token>
```

### 用户管理

#### 获取用户列表
```http
GET /users?page=1&page_size=10&sort_by=created_at&sort_order=desc
Authorization: Bearer <token>
```

#### 创建用户
```http
POST /users
Authorization: Bearer <token>
Content-Type: application/json

{
  "username": "new_user",
  "email": "new@example.com",
  "password": "123456",
  "nickname": "New User"
}
```

### 人员管理

#### 获取人员列表
```http
GET /persons?page=1&page_size=10&keyword=张三&search_by=name
```

#### 创建人员
```http
POST /persons
Content-Type: application/json

{
  "name": "张三",
  "age": 25,
  "gender": "男",
  "email": "zhangsan@example.com",
  "phone": "13800138000",
  "address": "北京市朝阳区",
  "company": "科技有限公司",
  "position": "软件工程师",
  "status": 1
}
```

### 时区处理

#### 获取当前时区
```http
GET /timezone/current
```

#### 获取指定时区时间
```http
GET /timezone/time?timezone=Asia/Shanghai
```

#### 时区转换
```http
GET /timezone/convert?time=2023-01-01 12:00:00&from_timezone=Asia/Shanghai&to_timezone=America/New_York
```

## 🔧 开发指南

### 代码生成

生成新的业务模块：

```bash
make wire
```

### 数据库迁移

```bash
# 执行迁移
make migrate-up

# 回滚迁移
make migrate-down

# 查看迁移状态
make migrate-status

# 重置数据库（危险操作）
make migrate-reset
```

### 生成 API 文档

```bash
make swagger
```

### 构建和部署

```bash
# 开发环境
make dev

# 测试环境
make test

# 生产环境
make production
```

## 🏗️ 架构设计

### Clean Architecture

```
┌─────────────────────────────────────────┐
│              Handler Layer              │  ← HTTP请求处理
├─────────────────────────────────────────┤
│              Service Layer              │  ← 业务逻辑
├─────────────────────────────────────────┤
│            Repository Layer             │  ← 数据访问
├─────────────────────────────────────────┤
│              Model Layer                │  ← 数据模型
└─────────────────────────────────────────┘
```

### 中间件链

```
请求 → Recovery → CORS → Trace → SQL Logger → Logger → Response → 业务逻辑
```

### 依赖注入

使用 Google Wire 进行依赖注入，支持：
- 自动依赖解析
- 接口绑定
- 生命周期管理

## 🔒 安全特性

- **密码加密**: 使用 bcrypt 进行密码加密
- **JWT认证**: 安全的 Token 认证机制
- **CORS配置**: 跨域安全配置
- **输入验证**: 完整的请求参数验证
- **错误处理**: 安全的错误信息返回

## 📊 监控和日志

### 健康检查

访问 `/health` 端点获取系统状态：

```json
{
  "status": "healthy",
  "timestamp": "2023-01-01T00:00:00Z",
  "uptime": "1h30m45s",
  "system": {
    "go_version": "go1.25.0",
    "go_os": "linux",
    "go_arch": "amd64",
    "num_cpu": 4,
    "num_goroutine": 10
  },
  "memory": {
    "alloc": 1024000,
    "total_alloc": 2048000,
    "sys": 4096000,
    "num_gc": 5
  }
}
```

### 日志系统

- **结构化日志**: JSON 格式输出
- **链路追踪**: 每个请求都有唯一的 TraceID
- **多级别日志**: Debug、Info、Warn、Error
- **文件轮转**: 自动日志文件轮转

## 🐳 Docker 部署

### 构建镜像

```bash
docker build -t toge-api .
```

### 运行容器

```bash
docker run -d \
  --name toge-api \
  -p 8080:8080 \
  -e ENV=production \
  toge-api
```

### 使用 Docker Compose

```yaml
version: '3.8'
services:
  api:
    build: .
    ports:
      - "8080:8080"
    environment:
      - ENV=production
    depends_on:
      - mysql
      - redis

  mysql:
    image: mysql:8.0
    environment:
      MYSQL_ROOT_PASSWORD: password
      MYSQL_DATABASE: toge
    ports:
      - "3306:3306"

  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"
```

## 📝 配置说明

### 环境配置

支持多环境配置：

- `config/dev.yaml` - 开发环境
- `config/test.yaml` - 测试环境  
- `config/production.yaml` - 生产环境

### 主要配置项

```yaml
app:
  name: "toge"
  version: "1.0.0"
  port: 8080
  mode: "debug"

database:
  driver: "mysql"
  host: "localhost"
  port: 3306
  username: "root"
  password: "password"
  database: "toge"

redis:
  host: "localhost"
  port: 6379
  password: ""
  database: 0

jwt:
  secret: "your-secret-key"
  expire_hours: 24
```

## 🤝 贡献指南

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 打开 Pull Request

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 🆘 支持

如果您遇到任何问题或有任何建议，请：

1. 查看 [Issues](https://github.com/your-username/toge-api/issues)
2. 创建新的 Issue
3. 联系维护者

## 📈 路线图

- [ ] 支持更多数据库类型
- [ ] 添加消息队列支持
- [ ] 实现分布式追踪
- [ ] 添加性能监控
- [ ] 支持 GraphQL API
- [ ] 添加单元测试覆盖
- [ ] 实现多租户支持

---

**toge API** - 让 API 开发更简单、更高效！
