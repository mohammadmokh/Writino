package email

import "github.com/mohammadmokh/writino/config"

type emailService struct {
	config.EmailCfg
}

func New(cfg config.EmailCfg) emailService {
	return emailService{
		EmailCfg: cfg,
	}
}
