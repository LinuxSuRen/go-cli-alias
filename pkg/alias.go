package pkg

import (
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"os"
	"path"
)

type DefaultAliasManager struct {
	Aliass map[string]string
	Path   string
}

func GetDefaultAliasMgr() (mgr *DefaultAliasManager, err error) {
	var userHome string
	if userHome, err = homedir.Dir(); err != nil {
		panic("cannot get the home directory")
	} else {
		mgr = &DefaultAliasManager{
			Aliass: map[string]string{},
			Path:   path.Join(userHome, ".config/ga/config.yaml"),
		}
		mgr.List()
	}
	return
}

func (a *DefaultAliasManager) List() map[string]string {
	var data []byte
	var err error
	if data, err = ioutil.ReadFile(a.Path); err == nil {
		err = yaml.Unmarshal(data, &a.Aliass)
	}
	return a.Aliass
}

func (a *DefaultAliasManager) Set(name, cmd string) (err error) {
	a.Aliass[name] = cmd
	err = a.save()
	return
}

func (a *DefaultAliasManager) Delete(name string) (err error) {
	delete(a.Aliass, name)
	err = a.save()
	return
}

func (a *DefaultAliasManager) save() (err error) {
	fmt.Println(path.Dir(a.Path))
	if err = os.MkdirAll(path.Dir(a.Path), 0771); err != nil {
		err = fmt.Errorf("cannot create directory: %s, error: %v", path.Dir(a.Path), err)
		return
	}

	var data []byte
	if data, err = yaml.Marshal(a.Aliass); err == nil {
		err = ioutil.WriteFile(a.Path, data, 0664)
	}
	return
}
