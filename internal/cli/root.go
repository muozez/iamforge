package cli

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "iamforge",
	Short: "Zero-backend Go CLI to scan code and generate least-privilege IAM policies",
	Long: `iamforge is a zero-backend Go CLI that performs static analysis on your code
and Terraform configurations to produce least-privilege AWS IAM policies and minimal Terraform blocks.`,
	SilenceUsage:  true,
	SilenceErrors: true,
}

// Execute runs the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(newVersionCmd())
}

