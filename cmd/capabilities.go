/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/openshift-telco/go-netconf-client/netconf"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
	"log"
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
			log.Fatal(err)
		}
		for _, cap := range s.Capabilities {
			fmt.Println(cap)
		}
	},
}

func init() {
	getCmd.AddCommand(capabilitiesCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// capabilitiesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// capabilitiesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
