package options

type Key struct {
	Auth  string `yaml:"auth"`
	Super string `yaml:"super"`
	Acl   string `yaml:"acl"`
}

func (c Key) AclKey(value string) string {
	return c.Acl + ":" + value
}
