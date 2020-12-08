package pkg

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
)

type DefaultAliasManager struct {
	Name         string
	Aliass       []Alias
	Path         string
	InitalAliass []Alias
}

func GetDefaultAliasMgrWithNameAndInitialData(name string, aliass []Alias) (mgr *DefaultAliasManager, err error) {
	if name == "" {
		name = "ga"
	}

	var userHome string
	if userHome, err = homedir.Dir(); err != nil {
		panic("cannot get the home directory")
	} else {
		mgr = &DefaultAliasManager{
			Aliass:       []Alias{},
			InitalAliass: aliass,
			Path:         path.Join(userHome, fmt.Sprintf(".config/%s/alias.yaml", name)),
		}
		mgr.List()
	}
	return
}

func GetDefaultAliasMgrWithName(name string) (mgr *DefaultAliasManager, err error) {
	return GetDefaultAliasMgrWithNameAndInitialData(name, nil)
}

func GetDefaultAliasMgr() (mgr *DefaultAliasManager, err error) {
	return GetDefaultAliasMgrWithName("")
}

func (a *DefaultAliasManager) List() []Alias {
	var data []byte
	var err error
	if data, err = ioutil.ReadFile(a.Path); err == nil {
		err = yaml.Unmarshal(data, &a.Aliass)
	}
	return a.Aliass
}

func (a *DefaultAliasManager) Set(name, cmd string) (err error) {
	exist := false
	for i, _ := range a.Aliass {
		if a.Aliass[i].Name == name {
			a.Aliass[i].Command = cmd
			exist = true
			break
		}
	}

	if !exist {
		a.Aliass = append(a.Aliass, Alias{
			Name:    name,
			Command: cmd,
		})
	}
	err = a.save()
	return
}

func (a *DefaultAliasManager) Delete(name string) (err error) {
	var index int
	for index, _ := range a.Aliass {
		if a.Aliass[index].Name == name {
			break
		}
	}

	if index != -1 {
		a.Aliass = append(a.Aliass[:index], a.Aliass[index+1:]...)
	}
	err = a.save()
	return
}

func (a *DefaultAliasManager) save() (err error) {
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

func (a *DefaultAliasManager) Init() (err error) {
	for _, v := range a.InitalAliass {
		if err = a.Set(v.Name, v.Command); err != nil {
			return
		}
	}

	if len(a.InitalAliass) > 0 {
		err = a.save()
	}
	return
}
