

package cmd


import (
    "os"
	"fmt"

	"github.com/spf13/cobra"
)


var profile string
var printJson bool
var printDebug bool
var RootCmd = &cobra.Command{
	Use:   "tags",
	Short: "Get AWS info via tags",
	Long: `Gets instance information out of AWS by performing
tag-based queries and returning properties from the results.

tags can be used either with an AWS CLI profile or inside
an instance with an IAM role.
`,
}


func Execute() {
	if err := RootCmd.Execute(); err != nil {
        fmt.Println("An error occured.")
		fmt.Println(err)
		os.Exit(-1)
	}
}


func init() {
	RootCmd.PersistentFlags().StringVar(&profile, "profile", "", "AWS profile to use")
    RootCmd.PersistentFlags().BoolVar(&printJson, "json", false, "Print json output")
	RootCmd.PersistentFlags().BoolVar(&printDebug, "debug", false, "Print debug output")
}
