package cmd

import (
	"fmt"
	"os"

	"github.com/cars/ACScraper/network"

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
		ifaces := network.GetInterfaces()
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
		var ip string
		if len(args) == 0 {
			fmt.Println("Using default behavior")
			ip = network.GetDefaultLocalIP()
		} else {
			fmt.Printf("Looking for interface: %s", args[0])
			ip = network.GetIPByInterfaceName(args[0])
			fmt.Println("....found")
		}
		network.RunProxy(ip)
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
	//nop
	return
}
