package cmd

import (
	"fmt"
	"sort"
	"text/tabwriter"
	"os"

	"autocmd/config"
	"autocmd/messages"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: messages.ListShort,
	Long:  messages.ListLong,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			fmt.Println(fmt.Sprintf(messages.ConfigLoadError, err))
			return
		}

		if len(cfg.Commands) == 0 {
			fmt.Println(messages.ListNoCommands)
			return
		}

		// 按别名排序，方便查看
		aliases := make([]string, 0, len(cfg.Commands))
		for alias := range cfg.Commands {
			aliases = append(aliases, alias)
		}
		sort.Strings(aliases)

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
		fmt.Fprintln(w, messages.ListHeader) // 表头
		fmt.Fprintln(w, messages.ListDivider) // 分隔线

		for _, alias := range aliases {
			cmdEntry := cfg.Commands[alias]
			fmt.Fprintf(w, "%s\t%s\t%s\n", cmdEntry.Alias, cmdEntry.Command, cmdEntry.Description)
		}
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

