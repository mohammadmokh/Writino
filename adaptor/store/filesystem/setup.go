package filesystem

import (
	"os"

	"github.com/mohammadmokh/writino/config"
)

type FsStore struct {
	config.FsCfg
}

func New(cfg config.FsCfg) (FsStore, error) {

	cfg.BasePath = cfg.BasePath + "/avatars/"
	err := os.MkdirAll(cfg.BasePath, os.ModePerm)
	if err != nil {
		return FsStore{}, err
	}

	return FsStore{
		cfg,
	}, nil
}
