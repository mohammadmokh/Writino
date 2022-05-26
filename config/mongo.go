package config

type MongoCfg struct {
	DBName string `yaml:"db_name"`
	Uri    string
}
