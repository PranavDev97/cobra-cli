package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// subcmdCmd represents the subcmd command
var subcmdCmd = &cobra.Command{
	Use:   "subcmd",
	Short: "Sub command of the CLI application",
	Long:  `Sub command which prints "Sub command" in the terminal`,
	Run: func(cmd *cobra.Command, args []string) {
		flagArg, err := cmd.Flags().GetString("flagexample")
		if err != nil {
			log.Println("Flag error: ", err)
		}
		if flagArg != "" {
			fmt.Println("Sub Command Along with falg value :", flagArg)
		} else {
			fmt.Println("Sub command")
		}
	},
}

func init() {
	rootCmd.AddCommand(subcmdCmd)
}
