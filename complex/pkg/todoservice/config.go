package todoservice

type Config struct {
	HTTPPort int
}

func (s *Config) FileName() string {
	return "config"
}

func (s *Config) Path() []string {
	return []string{
		"./",
		"$GOPATH/src/github.com/thedevelopnik/go-kit-example",
	}
}

func (s *Config) EnvVarPrefix() string {
	return "TODO"
}
