// Package cmd provides the command-line interface for CommitGPT.
/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"html"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/doodleEsc/CommitGPT/git"
	"github.com/doodleEsc/CommitGPT/prompt"
	"github.com/doodleEsc/CommitGPT/provider"
	"github.com/doodleEsc/CommitGPT/provider/message"
	"github.com/doodleEsc/CommitGPT/utils"
	"github.com/erikgeiser/promptkit/confirmation"
	"gopkg.in/yaml.v2"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	diffUnified int
	commitAmend bool
	preview     bool
	noConfirm   bool = false
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Generate Commit Message By LLM",
	RunE: func(cmd *cobra.Command, args []string) error {
		g := git.New(
			git.WithDiffUnified(viper.GetInt("git.diff_unified")),
			git.WithEnableAmend(viper.GetBool("git.amend")),
			git.WithExcludeList(viper.GetStringSlice("git.exclude_list")),
		)

		preview = viper.GetBool("commit.preview")
		noConfirm = viper.GetBool("commit.noConfirm")

		diff, err := g.DiffFiles()
		if err != nil {
			return err
		}

		systemContent, err := utils.GetTemplateByString(prompt.SystemTemplate, nil)
		if err != nil {
			return err
		}

		diffContent, err := utils.GetTemplateByString(prompt.CommitTemplate, utils.Data{
			"diff": diff,
		})
		if err != nil {
			return err
		}

		messages := []message.Message{
			{
				Role:    "system",
				Content: systemContent,
			},
			{
				Role:    "user",
				Content: diffContent,
			},
		}

		provider, err := provider.NewProvider("openai")
		if err != nil {
			return err
		}

		var commitMessage string
		response, err := provider.Completion(cmd.Context(), messages)
		if err != nil {
			return err
		}

		// Format commitMessage
		commitMessage = response.Content
		commitMessage = html.EscapeString(commitMessage)
		commitMessage = strings.TrimSpace(commitMessage)

		// Output commit summary data from AI
		color.Yellow("================Commit Summary====================")
		color.Yellow("\n" + commitMessage + "\n\n")
		color.Yellow("==================================================")
		color.Cyan("Prompt Tokens: %d, Completion Tokens: %d, Total Tokens: %d", response.Usage.PromptTokens, response.Usage.CompletionTokens, response.Usage.TotalTokens)

		// Handle commit message change prompt when confirmation is enabled
		if viper.GetBool("commit.preview") {
			if noConfirm {
				return nil
			}

			if ready, err := confirmation.New("Commit this preview summary?", confirmation.Yes).RunPrompt(); err != nil || !ready {
				if err != nil {
					return err
				}
				return nil
			}

			return nil
		}

		if !viper.GetBool("commit.no_confirm") {
			if change, err := confirmation.New("Do you want to modify the commit message?", confirmation.No).RunPrompt(); err != nil {
				return err
			} else if change {
				m := initialPrompt(commitMessage)
				p := tea.NewProgram(m, tea.WithContext(cmd.Context()))
				if _, err := p.Run(); err != nil {
					return err
				}
				p.Wait()
				commitMessage = m.textarea.Value()
			}
		}

		color.Cyan("Recording changes to the repository")
		finalCommitMessage, err := g.Commit(commitMessage)
		if err != nil {
			return err
		}
		color.Yellow(finalCommitMessage)

		return nil
	},
}

func GetDefaultCommitConfigAsYAML() (string, error) {
	config := map[string]any{
		"commit": map[string]any{
			"no_confirm": noConfirm,
			"preview":    preview,
		},
	}

	yamlBytes, err := yaml.Marshal(config)
	if err != nil {
		return "", fmt.Errorf("failed to marshal config to YAML: %w", err)
	}

	return string(yamlBytes), nil
}

func init() {
	rootCmd.AddCommand(commitCmd)
	commitCmd.Flags().IntVar(&diffUnified, "diff_unified", 3, "show <n> lines of diff context")
	commitCmd.Flags().BoolVar(&commitAmend, "amend", false, "amend previous commit")
	commitCmd.Flags().BoolVar(&preview, "preview", true, "preview commit message instead of committing")
	commitCmd.Flags().BoolVar(&noConfirm, "no_confirm", false, "git commit without confirm")

	viper.BindPFlag("git.diff_unified", commitCmd.Flags().Lookup("diff_unified"))
	viper.BindPFlag("git.amend", commitCmd.Flags().Lookup("amend"))
	viper.BindPFlag("git.preview", commitCmd.Flags().Lookup("preview"))
	viper.BindPFlag("commit.preview", commitCmd.Flags().Lookup("preview"))
	viper.BindPFlag("git.amend", commitCmd.Flags().Lookup("amend"))
}
