package command

import (
	"github.com/spf13/cobra"
	"time"
)

// The schema of apnic address file is
// registry|cc|type|start|value|date|status[|extensions...]
type ResultData struct {
	registry, cc, thetype string
	allDate               time.Time
}

var checkCMD = &cobra.Command{
	Use: "",
}

func init() {
	rootCMD.AddCommand(checkCMD)
}
