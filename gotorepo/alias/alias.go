package alias

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Alias represents a single alias in the repo
type Alias struct {
	Alias string
	Data DataV1
}

// DataV1 is the alias content from .yaml V1
type DataV1 struct {
	Target string
}

// New creates a new alias
func New(alias string, target string) (*Alias) {
	data := DataV1{
		Target: target,
	}

	return &Alias{alias, data}
}

// Load reads an existing alias
func Load(repo string, alias string) (*Alias, error) {
	a := &Alias{Alias: alias}
	exists, err := a.Exists(repo)
	if err != nil { return a, err }
	if exists == false { return a, errors.New("Alias does not exist") }

	contents, err := ioutil.ReadFile(filepath.Join(repo, alias))
	if err != nil { return a, err }

	err = yaml.Unmarshal(contents, &a.Data)
	return a, err
}

// Exists returns true if the alias already exists
func (a *Alias) Exists(repo string) (bool, error) {
	path := filepath.Join(repo, a.Alias)
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return true, err
}

// Save writes the alias, returns error on failure
func (a *Alias) Save(repo string) error {
	content, err := yaml.Marshal(a.Data)
	if err != nil { return err }

	fullPath := filepath.Join(repo, a.Alias)
	path, _ := filepath.Split(fullPath)
	err = os.MkdirAll(path, 0700)
	if err != nil { return err }

	file, err := os.Create(fullPath)
	if err != nil { return err }
	defer file.Close()

	file.Write(content)
	file.Sync()

	return nil
}
