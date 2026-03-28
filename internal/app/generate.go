package app

import (
	"io/fs"
	"os"
	"path/filepath"
)

// Generate renders a single page (thin slice) from content/index.md using a trivial template.

func Generate(contentDir, outDir string) error {
	// contentDir is a CLI-provided local directory for this trusted developer tool.
	// #nosec G304
	fsys := os.DirFS(filepath.Clean(contentDir))
	b, err := fs.ReadFile(fsys, "index.md")
	if err != nil {
		return err
	}

	html := string(b) // Render(string(b))
	// html, err := Render(string(b))
	// if err != nil {
	// 	return err
	// }

	if err := os.MkdirAll(outDir, 0750); err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(outDir, "index.html"), []byte(html), 0600)
}
