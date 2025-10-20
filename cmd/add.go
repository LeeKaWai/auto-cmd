package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"autocmd/config"
	"autocmd/messages"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <alias> [command] [description]",
	Short: messages.AddShort,
	Long:  messages.AddLong,
	Args:  cobra.MinimumNArgs(1), // 只需要别名
	Run: func(cmd *cobra.Command, args []string) {
		alias := args[0]
		multiline, _ := cmd.Flags().GetBool("multiline")
		
		cfg, err := config.LoadConfig()
		if err != nil {
			fmt.Println(fmt.Sprintf(messages.ConfigLoadError, err))
			return
		}

		if _, exists := cfg.Commands[alias]; exists {
			fmt.Println(fmt.Sprintf(messages.AddAliasExists, alias))
			return
		}

		// 输入命令
		var command string
		if len(args) > 1 {
			// 如果提供了命令参数，直接使用
			command = args[1]
		} else if multiline {
			// 多行输入模式
			fmt.Println("请输入多步骤命令（每行一个命令，输入 'END' 结束）:")
			var commands []string
			reader := bufio.NewReader(os.Stdin)
			for {
				fmt.Print("> ")
				line, _ := reader.ReadString('\n')
				line = strings.TrimSpace(line)
				if line == "END" {
					break
				}
				if line != "" {
					commands = append(commands, line)
				}
			}
			if len(commands) == 0 {
				fmt.Println("未输入任何命令")
				return
			}
			command = strings.Join(commands, " && ")
		} else {
			// 单行输入
			fmt.Print("请输入命令: ")
			reader := bufio.NewReader(os.Stdin)
			cmdInput, _ := reader.ReadString('\n')
			command = strings.TrimSpace(cmdInput)
		}

		// 输入描述
		var description string
		if len(args) > 2 {
			description = strings.Join(args[2:], " ")
		} else {
			fmt.Print("请输入命令描述（可选）: ")
			reader := bufio.NewReader(os.Stdin)
			desc, _ := reader.ReadString('\n')
			description = strings.TrimSpace(desc)
		}

		cfg.Commands[alias] = config.Command{
			Alias:       alias,
			Command:     command,
			Description: description,
		}

		err = config.SaveConfig(cfg)
		if err != nil {
			fmt.Println(fmt.Sprintf(messages.ConfigSaveError, err))
			return
		}

		fmt.Println(fmt.Sprintf(messages.AddSuccess, alias))
	},
}

func init() {
	addCmd.Flags().BoolP("multiline", "m", false, "多行输入模式，支持输入多个步骤")
	rootCmd.AddCommand(addCmd)
}

