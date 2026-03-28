package cli

import (
	"github.com/petersellars/loglet/internal/app"
	"github.com/spf13/cobra"
	"path/filepath"
)

var (
	contentDir string
	outputDir  string
)

func init() {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Render content → static HTML",
		RunE: func(cmd *cobra.Command, args []string) error {
			absContent, _ := filepath.Abs(contentDir)
			absOut, _ := filepath.Abs(outputDir)
			return app.Generate(absContent, absOut)
		},
	}
	cmd.Flags().StringVar(&contentDir, "content", "content", "Content Directory")
	cmd.Flags().StringVar(&outputDir, "out", "public", "Output Directory")
	rootCmd.AddCommand(cmd)
}
