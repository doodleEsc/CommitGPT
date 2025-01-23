/*
Copyright © 2025 doodleEsc cinuor@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "commitgpt",
	Short: "Manage Commit Message By LLM",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// 检查 git 命令
		if _, err := exec.LookPath("git"); err != nil {
			return fmt.Errorf("git command not found, please install git first")
		}
		return nil
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
