package config

type EmailCfg struct {
	Address  string
	Password string
	SmtpHost string `yaml:"smtp_host"`
	SmtpPort int    `yaml:"smtp_port"`
}
