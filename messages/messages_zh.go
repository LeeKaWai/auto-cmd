package messages

const (
	// Root 命令
	RootShort    = "您的智能命令行助手"
	RootLong     = `智令 (Auto-CMD) 是一个智能命令行助手，用于保存、管理和快速执行您常用的、复杂的命令。`
	ExecuteError = "执行命令时发生错误: %v\n"

	// Add 命令
	AddShort       = "添加一个新命令到您的智令库"
	AddLong        = `使用 'add' 命令来保存一个您经常使用的长命令，并为其指定一个易于记忆的别名和描述。
支持参数占位符 {{参数名}}，执行时会提示用户输入。
支持多行输入模式，使用 -m 或 --multiline 标志。

例如：
  qc add dps "docker ps -a --format 'table {{.ID}}\t{{.Names}}\t{{.Status}}'" "查看所有 Docker 容器"
  qc add myip "curl ifconfig.me" "查看本机公网IP"
  qc add -m deploy  # 多行输入模式，支持输入多个步骤
`
	AddAliasExists = "错误：别名 '%s' 已存在，请使用其他别名或考虑编辑现有命令。"
	AddSuccess     = "命令 '%s' 已成功添加到智令库。"

	// Run 命令
	RunShort         = "执行一个已保存的命令"
	RunLong          = `使用 'run' 命令通过别名快速执行您之前保存的命令。
例如：
  qc run dps
  qc run myip
`
	RunAliasNotFound = "错误：未找到别名为 '%s' 的命令。"
	RunExecuting     = "正在执行命令: %s"
	RunShellNotFound = "错误：无法找到 shell 执行器 (sh/bash): %v"
	RunExecFailed    = "错误：执行命令失败: %v"

	// List 命令
	ListShort      = "列出所有已保存的命令"
	ListLong       = `使用 'list' 命令查看智令库中所有已保存的命令，包括它们的别名、实际命令和描述。`
	ListNoCommands = "智令库中还没有任何命令。使用 'qc add' 添加您的第一个命令吧！"
	ListHeader     = "别名\t命令\t描述"
	ListDivider    = "----\t----\t----"

	// Config 相关
	ConfigUserHomeError   = "无法获取用户主目录: %v"
	ConfigAccessError     = "无法访问配置文件 '%s': %v"
	ConfigReadFileError   = "无法读取配置文件 '%s': %v"
	ConfigParseError      = "解析配置文件 '%s' 失败，请检查格式: %v"
	ConfigMarshalError    = "序列化配置失败: %v"
	ConfigWriteFileError  = "写入配置文件 '%s' 失败: %v"
	ConfigLoadError       = "加载配置失败: %v"
	ConfigSaveError       = "保存配置失败: %v"
)

