package mod

import (
	"golang.org/x/mod/modfile"
	"golang.org/x/mod/module"
	"io/ioutil"
)

func Mod(path string) (map[string]*module.Version, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	file, err := modfile.Parse("Mod", bytes, nil)
	if err != nil {
		return nil, err
	}
	require := map[string]*module.Version{}

	if len(file.Require) > 0 {
		for _, r := range file.Require {
			require[r.Mod.Path] = &r.Mod
		}
	}
	if len(file.Replace) > 0 {
		for _, r := range file.Replace {
			require[r.Old.Path] = &r.New
		}
	}
	return require, nil
}
