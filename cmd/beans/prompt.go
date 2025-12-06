package beans

import (
	"fmt"

	"github.com/spf13/cobra"
)

var promptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "Output instructions for AI coding agents",
	Long:  `Outputs a prompt that instructs AI coding agents on how to use the beans CLI to manage project issues.`,
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Print(agentPrompt)
		return nil
	},
}

const agentPrompt = `# Beans - Agentic Issue Tracker

This project uses **beans**, an agentic-first issue tracker. Issues are called "beans", and you can
use the "beans" CLI to manage them. Below are instructions on how to interact with beans.

All commands support --json for machine-readable output. Use this flag to parse responses easily.

### List available beans

To list all beans, use:

    beans list --json

`

func init() {
	rootCmd.AddCommand(promptCmd)
}
