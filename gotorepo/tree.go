package gotorepo

import (
	"os"

	"github.com/a8m/tree"
)

type treefs struct{}

func (f *treefs) Stat(path string) (os.FileInfo, error) {
	return os.Lstat(path)
}

func (f *treefs) ReadDir(path string) ([]string, error) {
	dir, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	names, err := dir.Readdirnames(-1)
	dir.Close()
	if err != nil {
		return nil, err
	}
	return names, nil
}

// Print the repo tree
func Print(path string) {
	//var outFile = os.Stdout
	opts := &tree.Options{
		Fs:      	new(treefs),
		OutFile: 	os.Stdout,
		FollowLink: true,
		Colorize: 	true,
	}

	inf := tree.New(path)
	inf.Visit(opts)
	inf.Print(opts)
}
