package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"../network"
)


var rootCmd = &cobra.Command{
	Use: "acscraper",
	Short "Get AC Infinity Login token",
	Run: func(cmd *cobra.Command, args []string){
		fmt.Print("run with --help for more information\n")
	}


}
var listCmd = &cobra.Command{
	Use: "list",
	Aliases: []string{"l","list-interfaces","get-interfaces"}
	Short: "List interfaces available on the machine",
	Run: func(cmd *cobra.Command, args []string){
		ifaces := getInterfaces()
		for _, iface := range ifaces{
			fmt.Printf("Interface Name: %s\n",iface)
		}
	}
}

var proxyCmd = &cobra.Command{
	Use: "proxy {interface-name}",
	Aliases: []string{"p","start-proxy","get-interfaces"}
	Short: "List interfaces available on the machine", 
	Run: func(cmd *cobra.Command, args []string){
		if args[0] {
			ip:= getDefaultLocalIP()
		}else {
			ip:= getIPByInterfaceName(args[0])
		}

	}

}


func Execute(){
	rootCmd.AddCommand(listCmd,proxyCmd)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize (initConfig)
}
func initConfig(){
	return
}