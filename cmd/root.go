package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/charmbracelet/fang"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cmdName string = "bubbles"
var configFile string = "config.yaml"

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   cmdName,
	Short: "charmbracelet/bubbles lib examples",
	Long:  "charmbracelet/bubbles lib examples",
}

func Execute() {
	if err := fang.Execute(context.Background(), rootCmd); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file locations are: ./"+configFile+", $HOME/.config/"+cmdName+"/"+configFile+", /etc/"+cmdName+"/"+configFile)
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "enable verbose output")

	// Bind flag to viper key
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/.config/" + cmdName)
		viper.AddConfigPath("/etc/" + cmdName)
		viper.SetConfigName("config")
	}

	_ = godotenv.Load()
	viper.AutomaticEnv()        // read env vars that match
	viper.SetEnvPrefix(cmdName) // BUBBLES_ARGNAME=...

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
