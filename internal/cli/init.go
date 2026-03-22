package cli

import (
	"fmt"
	"github.com/petersellars/loglet/internal/app"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var (
	initDir string
	force   bool
)

func init() {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Scaffold a site from the embedded template",
		Long:  "Scaffold a blog site from the embedded template to the specified directory.",
		RunE: func(cmd *cobra.Command, args []string) error {
			abs, err := filepath.Abs(initDir)
			if err != nil {
				return err
			}
			if !force {
				if _, err := os.Stat(filepath.Join(abs, "site.yaml")); err == nil {
					return fmt.Errorf("site already initialized at %s (site.yaml exists); use --force to overwrite", abs)
				}
			}
			return app.Init(abs, force)
		},
	}
	cmd.Flags().StringVarP(&initDir, "dir", "d", ".", "Directory to initialize the site in")
	cmd.Flags().BoolVarP(&force, "force", "f", false, "Force initialization by overwriting existing files")
	rootCmd.AddCommand(cmd)
}
