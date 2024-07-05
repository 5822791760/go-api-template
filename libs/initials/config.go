package initials

import (
	"github.com/5822791760/go-api-template/config"
)

func InitConfig() error {
	if err := config.LoadConfig(); err != nil {
		return err
	}

	return nil
}
