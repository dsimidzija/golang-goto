package gotorepo

import (
	//"fmt"
	"os"
	"path/filepath"

	//"github.com/spf13/viper"

	homedir "github.com/mitchellh/go-homedir"
)

// GetDefaultRepoPath returns the default $HOME/.goto repo path
func GetDefaultRepoPath() (string, error) {
	home, err := homedir.Dir()
	if err != nil { return "", err }

	path, err := GetRepoPath(filepath.Join(home, ".goto"))
	if err != nil { return "", err }

	return path, nil
}

// GetRepoPath builds the currently active absolute path to the repo, from the environment GOTO_ROOT_PATH,
// for now doing only Abs()
func GetRepoPath(root string) (string, error) {
	abspath, err := filepath.Abs(root)
	if err != nil { return "", nil }

	return abspath, nil
}

// IsValidRepo returns whether the given directory is a valid goto repo.
func IsValidRepo(dir string) (bool, error) {
    ident := filepath.Join(dir, ".gotoversion")
    _, err := os.Stat(ident)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return true, err
}

// IsValidTarget checks if the target folder exists
func IsValidTarget(dir string) (bool, error) {
    stat, err := os.Stat(dir)
    if err != nil {
	if os.IsNotExist(err) { return false, nil }
	return false, err
    }

    if stat.IsDir() != true { return false, nil }
    return true, err
}
