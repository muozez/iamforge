package cli

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

var (
	// Version is populated via ldflags at build time.
	Version = "dev"
	// Commit is populated via ldflags at build time.
	Commit = "none"
	// Date is populated via ldflags at build time.
	Date = "unknown"
	// DataVersion is the IAM action map data set version.
	DataVersion = "2026-11-07"
)

// VersionInfo models structured version output.
type VersionInfo struct {
	Version     string `json:"version"`
	Commit      string `json:"commit"`
	Date        string `json:"date"`
	DataVersion string `json:"data_version"`
}

func newVersionCmd() *cobra.Command {
	var jsonOutput bool

	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print iamforge version information",
		RunE: func(cmd *cobra.Command, args []string) error {
			return writeVersion(cmd.OutOrStdout(), jsonOutput)
		},
	}

	cmd.Flags().BoolVar(&jsonOutput, "json", false, "Output version information in JSON format")
	return cmd
}

func writeVersion(w io.Writer, asJSON bool) error {
	info := VersionInfo{
		Version:     Version,
		Commit:      Commit,
		Date:        Date,
		DataVersion: DataVersion,
	}

	if asJSON {
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "  ")
		return encoder.Encode(info)
	}

	_, err := fmt.Fprintf(w, "iamforge %s (commit: %s, date: %s, data: %s)\n",
		info.Version, info.Commit, info.Date, info.DataVersion)
	return err
}

