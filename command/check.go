package command

import (
	"github.com/spf13/cobra"
)

var checkCMD = &cobra.Command{
	Use: "",
}

func init() {
	rootCMD.AddCommand(checkCMD)
}
