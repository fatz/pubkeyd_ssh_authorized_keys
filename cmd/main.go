package cmd

import (
	"fmt"
	"os"

	"github.com/fatz/pubkeyd_ssh_authorized_keys/onelogingh"
	"github.com/spf13/viper"
)

const AppName = "pubkeyd"

func Run() {
	viper.SetConfigName(AppName) // name of config file (without extension)
	viper.AddConfigPath("/etc")  // path to look for the config file in
	viper.AddConfigPath("$HOME") // call multiple times to add many search paths
	viper.AddConfigPath(".")     // optionally look for config in the working directory
	err := viper.ReadInConfig()  // Find and read the config file
	if err != nil {              // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if len(os.Args) < 2 || os.Args[1] == "" {
		fmt.Printf("Please provide a username.\n Uage: pubkeyd_ssh_authorized_keys_command <username>")
		os.Exit(1)
	}

	baseurl := viper.GetString("baseurl")

	if baseurl == "" {
		baseurl = "github.com"
	}

	client := onelogingh.NewOneloginGHClient(baseurl)

	client.Rewrite = viper.GetStringMapString("rewrite")

	if pathf := viper.GetString("pathf"); pathf != "" {
		client.Pathf = pathf
	}

	keys, err := client.RequestAuthorizedKeys(os.Args[1])

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error requesting authorized_keys - %v", err)
		os.Exit(1)
	}

	fmt.Print(keys)
	os.Exit(0)
}
