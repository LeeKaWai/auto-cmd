package cmd

import (
	"fmt"
	"os"

	"autocmd/messages"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "qc",
	Short: messages.RootShort,
	Long:  messages.RootLong,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, messages.ExecuteError, err)
		os.Exit(1)
	}
}

func init() {
	// 在这里可以绑定全局Flag
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.autocmd.yaml)")
}

