package email

import "gitlab.com/gocastsian/writino/config"

type emailService struct {
	config.EmailCfg
}

func New(cfg config.EmailCfg) emailService {
	return emailService{
		EmailCfg: cfg,
	}
}
