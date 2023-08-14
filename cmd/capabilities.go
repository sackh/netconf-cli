/*
Implementation of reply subcommand.
Command to get capabilities of the device.
*/
package cmd

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/openshift-telco/go-netconf-client/netconf"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
)

// capabilitiesCmd represents the capabilities command
var capabilitiesCmd = &cobra.Command{
	Use:   "capabilities",
	Short: "Get the capabilities of the device.",
	Long:  `capabilities command will return you a list of all the capabilties of the device.`,
	Run: func(cmd *cobra.Command, args []string) {
		sshConfig := &ssh.ClientConfig{
			User:            ConfigObj.String("username"),
			Auth:            []ssh.AuthMethod{ssh.Password(ConfigObj.String("password"))},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}
		s, err := netconf.DialSSH(fmt.Sprintf("%s:%d", ConfigObj.String("host"), 830), sshConfig)
		if err != nil {
			color.Redln(err)
		}
		for _, cap := range s.Capabilities {
			color.Greenln(cap)
		}
	},
}

func init() {
	getCmd.AddCommand(capabilitiesCmd)
}
