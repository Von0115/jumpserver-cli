/*
Copyright © 2024 Von0115 jackbewater0115@gmail.com

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const Version = "0.1"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "显示 jms-cli 版本信息",
	Long: `此命令可以显示 jms-cli 的版本信息`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("jms-cli version %s\n",Version)
	},
}

func init() {
	

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}