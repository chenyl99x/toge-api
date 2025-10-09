# 项目名称
PROJECT_NAME := toge

# Docker 相关变量
DOCKER_IMAGE_NAME := $(PROJECT_NAME)
CONTAINER_NAME := $(PROJECT_NAME)-$(ENV)

# 默认目标
.DEFAULT_GOAL := help

# 帮助信息
.PHONY: help
help: ## 显示帮助信息
	@echo "可用的命令:"
	@echo "  make dev        # 构建并启动开发环境"
	@echo "  make test       # 构建并启动测试环境"
	@echo "  make production # 构建并启动生产环境"
	@echo "  make migrate-up # 执行数据库迁移"
	@echo "  make migrate-down # 回滚数据库迁移"
	@echo "  make migrate-status # 查看迁移状态"
	@echo "  make migrate-reset # 重置数据库（危险操作）"
	@echo "  make swagger   # 生成 Swagger 文档"
	@echo "  make wire      # 生成 Wire 依赖注入代码"
	@echo "  make help       # 显示此帮助信息"

# 主命令：构建并启动指定环境的容器
.PHONY: dev test production
dev: ## 构建并启动开发环境
	@$(MAKE) _deploy ENV=dev

test: ## 构建并启动测试环境
	@$(MAKE) _deploy ENV=test

production: ## 构建并启动生产环境
	@$(MAKE) _deploy ENV=production

# 内部部署命令
.PHONY: _deploy
_deploy:
	@echo "🚀 开始构建并启动 $(ENV) 环境..."
	@$(MAKE) cleanup-$(ENV)
	@$(MAKE) build-$(ENV)
	@$(MAKE) run-$(ENV)
	@echo "✅ $(ENV) 环境启动完成!"

# 清理现有容器和镜像
.PHONY: cleanup-dev cleanup-test cleanup-production
cleanup-dev cleanup-test cleanup-production:
	@echo "🧹 清理现有的容器和镜像..."
	@docker stop $(CONTAINER_NAME) 2>/dev/null || true
	@docker rm $(CONTAINER_NAME) 2>/dev/null || true
	@docker rmi $(DOCKER_IMAGE_NAME):$(ENV) 2>/dev/null || true
	@echo "✅ 清理完成"

# 构建指定环境的镜像
.PHONY: build-dev build-test build-production
build-dev build-test build-production:
	@echo "🔨 构建 $(ENV) 环境 Docker 镜像..."
	@if [ "$(ENV)" = "production" ]; then \
		docker build \
			-f Dockerfile \
			-t $(DOCKER_IMAGE_NAME):$(ENV) \
			--build-arg ENV=production \
			--build-arg GIN_MODE=release \
			.; \
	elif [ "$(ENV)" = "test" ]; then \
		docker build \
			-f Dockerfile \
			-t $(DOCKER_IMAGE_NAME):$(ENV) \
			--build-arg ENV=test \
			--build-arg GIN_MODE=release \
			.; \
	else \
		docker build \
			-f Dockerfile \
			-t $(DOCKER_IMAGE_NAME):$(ENV) \
			--build-arg ENV=dev \
			--build-arg GIN_MODE=debug \
			.; \
	fi
	@echo "✅ $(ENV) 环境镜像构建完成"

# 运行指定环境的容器
.PHONY: run-dev run-test run-production
run-dev run-test run-production:
	@echo "🚀 启动 $(ENV) 环境容器..."
	@if [ "$(ENV)" = "production" ]; then \
		docker run -d \
			--name $(CONTAINER_NAME) \
			-p 8080:8080 \
			-e ENV=production \
			$(DOCKER_IMAGE_NAME):$(ENV); \
	elif [ "$(ENV)" = "test" ]; then \
		docker run -d \
			--name $(CONTAINER_NAME) \
			-p 8081:8080 \
			-e ENV=test \
			$(DOCKER_IMAGE_NAME):$(ENV); \
	else \
		docker run -d \
			--name $(CONTAINER_NAME) \
			-p 8080:8080 \
			-e ENV=dev \
			$(DOCKER_IMAGE_NAME):$(ENV); \
	fi
	@echo "✅ $(ENV) 环境容器启动完成"
	@echo "🌐 访问地址: http://localhost:8080"
	@if [ "$(ENV)" = "test" ]; then \
		echo "🌐 测试环境访问地址: http://localhost:8081"; \
	fi

# 数据库迁移命令
.PHONY: migrate-up migrate-down migrate-status migrate-reset
migrate-up: ## 执行数据库迁移
	@echo "🔄 执行数据库迁移..."
	@go run cmd/migrate/main.go -action=up -env=dev
	@echo "✅ 数据库迁移完成"

migrate-down: ## 回滚数据库迁移（需要指定版本）
	@echo "⚠️  回滚数据库迁移..."
	@echo "请输入要回滚的迁移版本:"
	@read version; \
	go run cmd/migrate/main.go -action=down -version=$$version -env=dev
	@echo "✅ 数据库迁移回滚完成"

migrate-status: ## 查看迁移状态
	@echo "📊 查看迁移状态..."
	@go run cmd/migrate/main.go -action=status -env=dev

migrate-reset: ## 重置数据库（危险操作）
	@echo "⚠️  警告：这将删除所有数据！"
	@echo "确认重置数据库？(y/N):"
	@read confirm; \
	if [ "$$confirm" = "y" ] || [ "$$confirm" = "Y" ]; then \
		go run cmd/migrate/main.go -action=reset -env=dev; \
		echo "✅ 数据库重置完成"; \
	else \
		echo "❌ 操作已取消"; \
	fi

# Swagger 文档生成
.PHONY: swagger
swagger: ## 生成 Swagger 文档
	@echo "📝 生成 Swagger 文档..."
	@swag init -g cmd/main.go -o docs --parseDependency --parseInternal
	@echo "✅ Swagger 文档生成完成"
	@echo "📚 文档位置: docs/"
	@echo "🌐 访问地址: http://localhost:8080/swagger/index.html"

# Wire 依赖注入生成
.PHONY: wire
wire: ## 生成 Wire 依赖注入代码
	@echo "🔧 生成 Wire 依赖注入代码..."
	@cd internal/wire && go generate
	@echo "✅ Wire 代码生成完成"