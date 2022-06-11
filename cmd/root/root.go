package root

import (
	"fmt"

	"github.com/spf13/cobra"

	"hello/version"
)

var Cmd = &cobra.Command{
	Use:   "hello",
	Short: "hello, " + version.Version,
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		fmt.Println("Hello World!")
	},
}
