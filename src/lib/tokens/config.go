package tokens

import (
	"time"

	"lib/config"
)

type Config struct {
	Secret           config.Secret
	OneTimeTokensTTL time.Duration
}
