package ctl

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Pwd Variable Structure
var Pwd = &cobra.Command{
	Use:   "pwd",
	Short: "Print current working directory",
	Long:  "Print Current Working Directory",
	Run: func(cmd *cobra.Command, args []string) {
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println(pwd)
	},
}
