package pkg

const AliasKey = "alias-mgr"

type Alias struct {
	Name    string
	Command string
}

type AliasManager interface {
	List() []Alias
	Set(name, cmd string) error
	Delete(name string) error
	Init() error
}
