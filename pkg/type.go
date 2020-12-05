package pkg

const AliasKey = "alias-mgr"

type Alias map[string]string

type AliasManager interface {
	List() map[string]string
	Set(name, cmd string) error
	Delete(name string) error
}
