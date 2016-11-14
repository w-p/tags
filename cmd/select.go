

package cmd


import (
	"fmt"
    "errors"
    "encoding/json"

	"github.com/spf13/cobra"
    "github.com/w-p/tags/awsapi"
)


var selectCmd = &cobra.Command{
	Use:   "select [property] where [filter pairs]",
	Short: "Select a property of an EC2 instance",
    Long: `Gets a property from a filtered list of EC2 instances.

A property is one of the following:
  PrivateDnsName
  PrivateIpAddress
  PublicDnsName
  PublicIpAddress
  InstanceType
  InstanceId
  SubnetId
  ImageId
  VpcId
  State

A filter is a space separated key value pair, eg:
  tag:Name MyInstance
  instance-type t2.micro

A key is one of the filter keys listed at:
  http://docs.aws.amazon.com/cli/latest/reference/ec2/describe-instances.html
`,
	RunE: func(cmd *cobra.Command, args []string) error {
        if printDebug {
            fmt.Println("AWS profile:   ", profile)
    		fmt.Println("print json:    ", printJson)
            fmt.Println("print debug:   ", printDebug)
        }

        if len(args) < 3 {
            msg := "Invalid clause, use: <property> where <filter pairs>"
            return errors.New(msg)
        }
        property := args[0]
        filters := args[2:]
        session, err := awsapi.GetSession(profile)
        instances, err := awsapi.FilterInstances(session, filters)
        results := awsapi.SelectProperty(&instances, property)

        if err == nil {
            if printJson {
                out, err := json.MarshalIndent(&results, "", "    ")
                if err == nil {
                    fmt.Println(string(out))
                }
            } else {
                for _, result := range results {
                    fmt.Println(result)
                }
            }
        }
        return err
	},
}


func init() {
	RootCmd.AddCommand(selectCmd)
}
