package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

func init() {
	rootCmd = &cobra.Command{
		Use:   "aurora-password-manager",
		Short: "极光密码管理工具",
		Long:  "极光密码管理工具 - 一个安全、便捷的密码管理器",
		Run: func(cmd *cobra.Command, args []string) {
			color.Green("✓ 您好，欢迎使用极光密码管理工具！")
			color.Yellow("提示: 使用 --help 查看可用命令")
		},
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "执行命令时发生错误: %v\n", err)
		os.Exit(1)
	}
}
