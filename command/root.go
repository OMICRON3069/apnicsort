package command

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

var rootCMD = &cobra.Command{
	Use:   "apnicsort",
	Short: "Sorting APNIC address info",
	Long:  "This is a program that auto download & sort address info provide by APNIC",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func Execute() {
	if err := rootCMD.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
