package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string

	PlayMobileUrl        string
	PlayMobileLogin      string
	PlayMobilePassword   string
	PlayMobileOriginator string

	SendGridApiKey string
	Mail           string
}

const (
	NotificationTypeSms   = "sms"
	NotificationTypeEmail = "email"
)

var NotificationTypes = []string{NotificationTypeSms, NotificationTypeEmail}

func Load() Config {
	config := Config{}

	config.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	config.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	config.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "postgres"))
	config.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "postgres"))
	config.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "buy_event"))

	config.PlayMobileUrl = cast.ToString(getOrReturnDefault("PLAY_MOBILE_URL", "http://91.204.239.44/broker-api/send"))
	config.PlayMobileOriginator = cast.ToString(getOrReturnDefault("PLAY_MOBILE_ORIGINATOR", "3700"))
	config.PlayMobileLogin = cast.ToString(getOrReturnDefault("PLAY_MOBILE_LOGIN", "delever"))
	config.PlayMobilePassword = cast.ToString(getOrReturnDefault("PLAY_MOBILE_PASSWORD", "aDev27hg$82@"))

	config.SendGridApiKey = cast.ToString(getOrReturnDefault("SENDGRID_API_KEY", ""))
	config.Mail = cast.ToString(getOrReturnDefault("MAIL", ""))

	return config
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)

	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
