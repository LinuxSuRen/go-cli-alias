package pkg

const AliasKey = "alias-mgr"

type Alias struct {
	Name    string `yaml:"name"`
	Command string `yaml:"command"`
}

type AliasManager interface {
	List() []Alias
	Set(name, cmd string) error
	Delete(name string) error
	Init() error
}
