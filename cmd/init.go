/*
Copyright © 2024 Von0115 jackbewater0115@gmail.com

*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)


var projectDir string // 全局变量，用于指定项目目录

// InitProjectDir 用于创建项目目录和初始化配置文件
func InitProjectDir(projectDir string) error {
	
	// 尝试创建项目目录
	dirErr := os.Mkdir(projectDir, os.ModePerm)    //仅当目标目录不存在时创建它；如果目录已存在，会返回 os.ErrExist 错误
	if dirErr != nil {
		return fmt.Errorf("%v", dirErr)      //创建目录失败
	}

	// 配置文件路径
	configFilePath := filepath.Join(projectDir, ".jms-cli.yaml")

	// 使用 Viper 设置默认配置项
	viper.Set("access_key_id", "your-access-key-id")
	viper.Set("access_key_secret","your-access_key_secret")
	viper.Set("jumpserver_url", "http://your-jumpserver-url")

	// 写入配置文件
	configErr := viper.WriteConfigAs(configFilePath)
	if configErr != nil {
		return fmt.Errorf("%v", configErr)      //极小可能因为权限、磁盘空间不够导致写入配置文件失败
	}

	// 输出成功信息
	fmt.Printf("项目目录和配置文件已创建成功：%s\n", configFilePath)
	return nil
}


// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "初始化一个项目目录",
	Long: `此命令会创建一个项目目录，其中包含了 .jms-cli.yaml 配置文件`,
	Run: func(cmd *cobra.Command, args []string) {
		
		// 如果未提供 --dir 参数，则使用默认目录
		if projectDir == "" {
			projectDir = "./jms-cli"
		}

		// 调用初始化方法
		err := InitProjectDir(projectDir)
		if err != nil {
			fmt.Printf("初始化失败: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	
	// 添加 --dir 参数，用于指定项目目录
	initCmd.Flags().StringVar(&projectDir, "dir", "", "指定项目目录（默认是 ./jms-cli）")
}
