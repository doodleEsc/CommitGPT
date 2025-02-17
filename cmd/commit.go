/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/doodleEsc/CommitGPT/git"
	"github.com/doodleEsc/CommitGPT/provider"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	diffUnified int
	commitAmend bool
	excludeList []string
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Generate Commit Message By LLM",
	RunE: func(cmd *cobra.Command, args []string) error {
		g := git.New(
			git.WithDiffUnified(viper.GetInt("git.diff_unified")),
			git.WithExcludeList(viper.GetStringSlice("git.exclude_list")),
			git.WithEnableAmend(commitAmend),
		)

		diff, err := g.DiffFiles()
		if err != nil {
			return err
		}

		provider, err := provider.NewProvider("openai")
		if err != nil {
			return err
		}

		response, err := provider.Completion(cmd.Context(), diff)
		if err != nil {
			return err
		}

		fmt.Println(response.Content)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// commitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// commitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	commitCmd.Flags().IntVar(&diffUnified, "diff_unified", 3, "show <n> lines of diff context")
	commitCmd.Flags().BoolVar(&commitAmend, "amend", false, "amend previous commit")
	commitCmd.Flags().StringSliceVar(&excludeList, "exclude_list", []string{}, "exclude file from git diff command")
}
