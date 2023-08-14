/*
Implementation of reply subcommand.
Command to get either running or candidate config from the device.
*/
package cmd

import (
	"github.com/gookit/color"
	"github.com/openshift-telco/go-netconf-client/netconf/message"
	"github.com/spf13/cobra"
	"os"
)

var candidate, running bool

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Get the running or candidate config.",
	Long:  `Get the running or candidate config.`,
	Run: func(cmd *cobra.Command, args []string) {
		var msg string
		if candidate && running {
			color.Redln("Please provide either candidate or running and not both.")
			return
		} else if candidate {
			msg = "<get-config><source><candidate/></source></get-config>"
		} else if running {
			msg = "<get-config><source><running/></source></get-config>"
		} else {
			color.Redln("Please provide either candidate or running.")
			return
		}
		temp := os.Stdout
		os.Stdout = nil
		rpcMessage := message.NewRPC(msg)
		session := CreateSession(830)
		defer session.Close()
		reply, err := session.SyncRPC(rpcMessage, 1)
		os.Stdout = temp
		if err != nil {
			color.Redln(err)
			return
		}
		color.Greenln("====== Result ======")
		color.Greenln(reply.RawReply)
		color.Greenln("====== Result ======")
	},
}

func init() {
	getCmd.AddCommand(configCmd)
	configCmd.Flags().BoolVarP(&candidate, "candidate", "c", false, "Get the candidate config")
	configCmd.Flags().BoolVarP(&running, "running", "r", false, "Get the running config")
}
