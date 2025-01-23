/*
Copyright © 2025 doodleEsc cinuor@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "CommitGPT",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.CommitGPT.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		cfgDir := path.Join(home, ".config", "commitgpt")
		if err := os.MkdirAll(cfgDir, os.ModePerm); err != nil {
			cobra.CheckErr(err)
		}
		cfgFile = path.Join(cfgDir, "commitgpt.yaml")

		// Search config in home directory with name ".CommitGPT" (without extension).
		viper.AddConfigPath(cfgDir)
		viper.SetConfigType("yaml")
		viper.SetConfigName("commitgpt")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {

			// Config file not found, create a default one
			if _, err := os.Create(cfgFile); err != nil {
				cobra.CheckErr(err)
			}

			// Write default config
			defaultConfig := []byte("# CommitGPT default configuration\n")
			if err := os.WriteFile(cfgFile, defaultConfig, 0644); err != nil {
				cobra.CheckErr(err)
			}
			fmt.Fprintln(os.Stderr, "Created default config file:", cfgFile)
		} else {
			// Config file was found but another error was produced
			cobra.CheckErr(err)
		}
	}
}
