package gotorepo

import (
	"os"
	"path/filepath"
)

// Init creates the repo root and the basic repo structure
func Init(root string) (error) {

	err := os.MkdirAll(root, 0700)
	if err != nil { return err }

	versionPath := filepath.Join(root, ".gotoversion")

	f, err := os.Create(versionPath)
	if err != nil { return err }
	defer f.Close()

	_, err = f.WriteString("1")
	if err != nil { return err }
	f.Sync()

	return nil
}
