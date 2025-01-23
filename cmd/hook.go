/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/fatih/color"

	"github.com/doodleEsc/CommitGPT/git"
	"github.com/spf13/cobra"
)

// hookCmd represents the hook command
var hookCmd = &cobra.Command{
	Use:   "hook",
	Short: "install/uninstall git prepare-commit-msg hook",
}

// hookCmd represents the hook command
var hookInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install git hook",
	RunE: func(cmd *cobra.Command, args []string) error {
		g := git.New()

		if err := g.InstallHook(); err != nil {
			fmt.Println(err)
			return err
		}

		color.Green("git hook prepare-commit-msg installed")
		color.Green("see in ./.git/hooks/prepare-commit-msg")

		return nil
	},
}

// hookCmd represents the hook command
var hookUninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall git hook",
	RunE: func(cmd *cobra.Command, args []string) error {
		g := git.New()

		if err := g.UninstallHook(); err != nil {
			return err
		}
		color.Green("remove git hook: prepare-commit-msg successfully")
		return nil
	},
}

func init() {
	hookCmd.AddCommand(hookInstallCmd, hookUninstallCmd)
	rootCmd.AddCommand(hookCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hookCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hookCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
