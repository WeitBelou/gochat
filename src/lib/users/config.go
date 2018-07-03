package users

import "lib/config"

type Config struct {
	Secret config.Secret
	DB     config.DB
}
