package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "acscraper",
	Short: "Get AC Infinity Login token",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("run with --help for more information\n")
	},
}
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l", "list-interfaces", "get-interfaces"},
	Short:   "List interfaces available on the machine",
	Run: func(cmd *cobra.Command, args []string) {
		ifaces := GetInterfaces()
		for _, iface := range ifaces {
			fmt.Printf("Interface Name: %s\n", iface)
		}
	},
}

var proxyCmd = &cobra.Command{
	Use:     "proxy {interface-name}",
	Aliases: []string{"p", "start-proxy", "get-interfaces"},
	Short:   "List interfaces available on the machine",
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] {
			ip := GetDefaultLocalIP()
		} else {
			ip := GetIPByInterfaceName(args[0])
		}

	},
}

func Execute() {
	rootCmd.AddCommand(listCmd, proxyCmd)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}
func initConfig() {
	return
}
