/*
This file implements root command of the netconf-cli.
*/
package cmd

import (
	"fmt"
	"os"

	"log"

	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/openshift-telco/go-netconf-client/netconf"
	"github.com/openshift-telco/go-netconf-client/netconf/message"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
)

var cfgFile string

var ConfigObj = koanf.New(".")

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "netconf-cli",
	Short: "A simple netconf CLI",
	Long:  `A CLI tool to connect to devices using netconf protocol`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("ROOT COMMAND")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "netconf-cli.env", "config file path")
	rootCmd.MarkPersistentFlagRequired("config")
}

func initConfig() {
	f := file.Provider(cfgFile)
	if err := ConfigObj.Load(f, dotenv.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}
}

func CreateSession(port int) *netconf.Session {
	sshConfig := &ssh.ClientConfig{
		User:            ConfigObj.String("username"),
		Auth:            []ssh.AuthMethod{ssh.Password(ConfigObj.String("password"))},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	s, err := netconf.DialSSH(fmt.Sprintf("%s:%d", ConfigObj.String("host"), port), sshConfig)
	if err != nil {
		log.Fatal(err)
	}
	capabilities := netconf.DefaultCapabilities
	err = s.SendHello(&message.Hello{Capabilities: capabilities})
	if err != nil {
		log.Fatal(err)
	}

	return s
}
