package app

import (
	"fmt"
	"github.com/petersellars/loglet/internal/embed"
	"io/fs"
	"os"
	"path/filepath"
)

func Init(dir string, force bool) error {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Write loglet.yaml
	cfg := []byte("loglet:\n  name: example\ncontentDir: content\nOutputDir: public\n")
	if err := writeFile(filepath.Join(dir, "loglet.yaml"), cfg, force); err != nil {
		return err
	}

	// Copy embeded starter tree to the target directory
	if err := fs.WalkDir(embed.Starter, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if path == "." {
			return nil
		}
		target := filepath.Join(dir, path)
		if d.IsDir() {
			return os.MkdirAll(target, 0755)
		}
		data, err := fs.ReadFile(embed.Starter, path)
		if err != nil {
			return err
		}
		return writeFile(target, data, force)
	}); err != nil {
		return err
	}

	fmt.Printf("Loglet initialized successfully in %s\n", dir)
	return nil
}

func writeFile(path string, b []byte, force bool) error {
	if !force {
		if _, err := os.Stat(path); err == nil {
			return fmt.Errorf("file %s already exists", path)
		}
	}
	return os.WriteFile(path, b, 0644)
}
