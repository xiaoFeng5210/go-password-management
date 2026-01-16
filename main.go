package main

import (
	"fmt"
	"os"

	redisDB "password-management/database/redis-db"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

func init() {
	redisDB.Connect()
	rootCmd = &cobra.Command{
		Use:   "aurora-password-manager",
		Short: "极光密码管理工具",
		Long:  "极光密码管理工具 - 一个安全、便捷的密码管理器",
		Run: func(cmd *cobra.Command, args []string) {
			color.Green("✓ 您好，欢迎使用极光密码管理工具！")
			if len(args) > 0 {
				switch args[0] {
				case "add":
					// 1, 2, 3分别是账号，密码，备注
					account := args[1]
					password := args[2]
					remark := args[3]
					fmt.Printf("account: %s, password: %s, remark: %s\n", account, password, remark)
					err := redisDB.AddPassword(account, password, remark)
					if err != nil {
						color.Red("添加密码失败: %v", err)
					} else {
						color.Green("添加密码成功")
					}
				case "update":
					account := args[1]
					password := args[2]
					redisDB.UpdatePassword(account, password)
				case "delete":
					account := args[1]
					redisDB.DeletePassword(account)
				case "get":
					account := args[1]
					redisDB.GetPassword(account)
				}
			}
		},
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "执行命令时发生错误: %v\n", err)
		os.Exit(1)
	}
}
