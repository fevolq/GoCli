package cmd

import (
	"github.com/spf13/cobra"
	"go-cli/model"
)

var CliCmd = &cobra.Command{
	Use:   "cli",
	Short: "生成cli",
	Long:  "读取文件，批量生成cli",
	Run:   run,
}

// 根据参数来判断执行的命令
func run(cmd *cobra.Command, args []string) {
	if filePath != "" {
		model.Execute(filePath)
	}
}

var filePath string

func init() {
	CliCmd.Flags().StringVarP(
		&filePath,   // 绑定的变量
		"file",      // 完整的命令标识
		"f",         // 短标识
		"",          // 默认值
		"请输入配置文件路径", // 提示信息
	)
	CliCmd.MarkFlagRequired("file") // 设置为必填
}
