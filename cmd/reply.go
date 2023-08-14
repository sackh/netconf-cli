/*
Implementation of reply subcommand.
Command to get reply from the device for given input RPC request.
*/
package cmd

import (
	"github.com/gookit/color"
	"github.com/openshift-telco/go-netconf-client/netconf/message"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"regexp"
)

var reqFile string

// replyCmd represents the reply command
var replyCmd = &cobra.Command{
	Use:   "reply",
	Short: "Get the reply from device for inpur RPC request.",
	Long:  `Get the reply from device for inpur RPC request.`,
	Run: func(cmd *cobra.Command, args []string) {
		rpcRequest, err := ioutil.ReadFile(reqFile)
		if err != nil {
			color.Redln("failed reading data from file: %s", err)
			return
		}
		rpcReq := string(rpcRequest[:])
		// library adds RPC tag hence removing RPC tag if already present in input.
		rpcReq = removeRPCTag(rpcReq)
		temp := os.Stdout
		os.Stdout = nil
		rpcMessage := message.NewRPC(rpcReq)
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
	getCmd.AddCommand(replyCmd)
	replyCmd.Flags().StringVarP(&reqFile, "rpc-req-file", "r", "", "RPC request file path")
	replyCmd.MarkFlagRequired("rpc-req-file")
}

func removeRPCTag(rpcRequest string) string {
	re1 := regexp.MustCompile(`<rpc\s+message-id=.*">`)
	repStr1 := re1.ReplaceAllString(rpcRequest, "")
	re2 := regexp.MustCompile(`[\\S\\r\\n]+</rpc>`)
	repStr2 := re2.ReplaceAllString(repStr1, "")
	return repStr2
}
