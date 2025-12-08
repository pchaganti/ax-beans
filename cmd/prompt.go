package cmd

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"hmans.dev/beans/internal/beancore"
	"hmans.dev/beans/internal/config"
)

//go:embed prompt.md
var agentPrompt string

var promptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "Output instructions for AI coding agents",
	Long:  `Outputs a prompt that instructs AI coding agents on how to use the beans CLI to manage project issues.`,
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Find beans directory; silently exit if none exists
		var root string
		var err error

		if beansPath != "" {
			root = beansPath
		} else {
			root, err = beancore.FindRoot()
			if err != nil {
				// No .beans directory found - silently exit
				return nil
			}
		}

		// Load config for dynamic sections
		cfg, err := config.Load(root)
		if err != nil {
			return fmt.Errorf("loading config: %w", err)
		}

		fmt.Print(agentPrompt)

		// Append dynamic sections from config
		if cfg != nil {
			var sb strings.Builder

			// Issue types section
			if len(cfg.Types) > 0 {
				sb.WriteString("\n## Issue Types\n\n")
				sb.WriteString("This project has the following issue types configured. Always specify a type with `-t` when creating beans:\n\n")
				for _, t := range cfg.Types {
					if t.Description != "" {
						sb.WriteString(fmt.Sprintf("- **%s**: %s\n", t.Name, t.Description))
					} else {
						sb.WriteString(fmt.Sprintf("- **%s**\n", t.Name))
					}
				}
			}

			// Statuses section
			if len(cfg.Statuses) > 0 {
				sb.WriteString("\n## Statuses\n\n")
				sb.WriteString("This project has the following statuses configured:\n\n")
				for _, s := range cfg.Statuses {
					if s.Description != "" {
						sb.WriteString(fmt.Sprintf("- **%s**: %s\n", s.Name, s.Description))
					} else {
						sb.WriteString(fmt.Sprintf("- **%s**\n", s.Name))
					}
				}
			}

			sb.WriteString("\n")
			fmt.Print(sb.String())
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(promptCmd)
}
