package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"

	"autocmd/config"
	"autocmd/messages"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run <alias>",
	Short: messages.RunShort,
	Long:  messages.RunLong,
	Args:  cobra.ExactArgs(1), // 只需要一个别名参数
	Run: func(cmd *cobra.Command, args []string) {
		alias := args[0]

		cfg, err := config.LoadConfig()
		if err != nil {
			fmt.Println(fmt.Sprintf(messages.ConfigLoadError, err))
			return
		}

		cmdEntry, exists := cfg.Commands[alias]
		if !exists {
			fmt.Println(fmt.Sprintf(messages.RunAliasNotFound, alias))
			return
		}

		// 处理参数替换
		finalCommand := processParameters(cmdEntry.Command)
		fmt.Println(fmt.Sprintf(messages.RunExecuting, finalCommand))

		// 跨平台执行命令
		var execCmd *exec.Cmd
		if isWindows() {
			// Windows 使用 cmd.exe
			execCmd = exec.Command("cmd", "/C", finalCommand)
		} else {
			// Unix/Linux 使用 sh
			execCmd = exec.Command("sh", "-c", finalCommand)
		}
		
		execCmd.Stdin = os.Stdin
		execCmd.Stdout = os.Stdout
		execCmd.Stderr = os.Stderr
		
		err = execCmd.Run()
		if err != nil {
			fmt.Println(fmt.Sprintf(messages.RunExecFailed, err))
			os.Exit(1)
		}
	},
}

// isWindows 检测是否为 Windows 系统
func isWindows() bool {
	return runtime.GOOS == "windows"
}

// processParameters 处理命令中的参数占位符
func processParameters(command string) string {
	// 查找所有 {{参数名}} 格式的占位符
	re := regexp.MustCompile(`\{\{([^}]+)\}\}`)
	
	return re.ReplaceAllStringFunc(command, func(match string) string {
		// 提取参数名
		paramName := strings.Trim(match, "{}")
		
		// 提示用户输入
		fmt.Printf("请输入 %s: ", paramName)
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		
		// 去除换行符
		return strings.TrimSpace(input)
	})
}

func init() {
	rootCmd.AddCommand(runCmd)
}

