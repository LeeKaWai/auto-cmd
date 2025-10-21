# Auto-CMD (智令)

一个智能命令行助手，用于保存、管理和快速执行您常用的、复杂的命令。

## 功能特性

- 🚀 **快速命令保存**：轻松保存常用命令，支持参数占位符
- 📝 **多行输入模式**：支持多步骤命令的便捷输入
- 🔍 **智能参数替换**：执行时动态输入参数值
- 📋 **命令列表管理**：查看和管理所有保存的命令
- 🎯 **跨平台支持**：Windows、Linux、macOS 全平台支持

## 安装

### 从源码编译

```bash
git clone git@github.com:LeeKaWai/auto-cmd.git
cd auto-cmd
go build -o qc.exe  # Windows
# 或
go build -o qc      # Linux/macOS
```

### 配置环境变量

将编译后的可执行文件添加到系统 PATH 中，或直接使用 `./qc` 运行。

## 使用方法

### 基本命令

```bash
qc add <alias> [command] [description]    # 添加命令
qc run <alias>                           # 执行命令
qc list                                 # 列出所有命令
qc --help                               # 查看帮助
```

### 添加命令

#### 1. 单行命令

```bash
# 基本用法
qc add myip "curl ifconfig.me" "查看本机公网IP"

# 带参数的命令
qc add dps "docker ps -a --format 'table {{.ID}}\t{{.Names}}\t{{.Status}}'" "查看所有 Docker 容器"

# 交互式添加
qc add test
请输入命令: echo "Hello World"
请输入命令描述（可选）: 测试命令
```

#### 2. 多行命令（新功能）

使用 `-m` 或 `--multiline` 标志进行多步骤命令输入：

```bash
qc add -m deploy
请输入多步骤命令（每行一个命令，输入 'END' 结束）:
> git pull origin main
> npm install
> npm run build
> pm2 restart myapp
> END
请输入命令描述（可选）: 部署应用程序
```

### 执行命令

```bash
# 执行保存的命令
qc run myip
qc run deploy

# 带参数的命令会提示输入
qc run dps
请输入 ID: 
请输入 Names: 
请输入 Status: 
```

### 管理命令

```bash
# 查看所有保存的命令
qc list

# 输出示例：
# 别名       命令                                                     描述
# ----     ----                                                      ----
# deploy   git pull origin main && npm install && npm run build    部署应用程序
# myip     curl ifconfig.me                                         查看本机公网IP
```

## 高级功能

### 参数占位符

使用 `{{参数名}}` 格式在命令中定义参数：

```bash
qc add build "go build -o {{output_name}} ." "构建Go项目"
qc add docker "docker build -t {{image_name}}:{{tag}} ." "构建Docker镜像"
```

执行时会提示输入参数值：

```bash
qc run build
请输入 output_name: myapp
正在执行命令: go build -o myapp .
```

### 多步骤命令

多行输入模式特别适合复杂的部署流程：

```bash
qc add -m deploy
请输入多步骤命令（每行一个命令，输入 'END' 结束）:
> git pull origin main
> npm install
> npm run build
> pm2 restart myapp
> END
```

## 配置文件

命令保存在用户主目录的 `.autocmd.yaml` 文件中：

```yaml
commands:
  myip:
    alias: myip
    command: curl ifconfig.me
    description: 查看本机公网IP
  deploy:
    alias: deploy
    command: git pull origin main && npm install && npm run build && pm2 restart myapp
    description: 部署应用程序
```

## 使用场景

### 开发环境

```bash
# 项目初始化
qc add -m init
> mkdir {{project_name}}
> cd {{project_name}}
> git init
> npm init -y
> END

# 代码格式化
qc add format "prettier --write . && eslint --fix ." "格式化代码"

# 测试运行
qc add test "npm test && npm run coverage" "运行测试并生成覆盖率报告"
```

### 部署流程

```bash
# 生产部署
qc add -m prod-deploy
> git checkout main
> git pull origin main
> npm ci
> npm run build
> pm2 restart production
> END

# 数据库迁移
qc add migrate "knex migrate:latest && knex seed:run" "运行数据库迁移"
```

### 系统维护

```bash
# 系统清理
qc add -m cleanup
> docker system prune -f
> npm cache clean --force
> go clean -cache
> END

# 日志查看
qc add logs "tail -f {{log_file}} | grep {{pattern}}" "实时查看日志"
```

## 命令参考

### add 命令

```bash
qc add <alias> [command] [description] [flags]

Flags:
  -h, --help        help for add
  -m, --multiline   多行输入模式，支持输入多个步骤
```

### run 命令

```bash
qc run <alias>
```

### list 命令

```bash
qc list
```

## 贡献

欢迎提交 Issue 和 Pull Request！

## 许可证

MIT License

## 更新日志

### v1.1.0
- ✨ 新增多行输入模式支持
- 🔧 改进用户交互体验
- 📝 更新帮助文档

### v1.0.0
- 🎉 初始版本发布
- ✨ 基本命令保存和执行功能
- ✨ 参数占位符支持
- ✨ 跨平台支持
