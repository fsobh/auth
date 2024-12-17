package util

import (
	"github.com/spf13/viper"
	"time"
)

// Config Stores all configuration env variables loaded from app.env using viper
// The annotations tell viper what the name of each value is in the .env file (uses map structure)
type Config struct {
	Enviroment           string        `mapstructure:"ENVIRONMENT"`
	DBDriver             string        `mapstructure:"DB_DRIVER"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	HttpServerAddress    string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	GrpcServerAddress    string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	MigrationUrl         string        `mapstructure:"MIGRATION_URL"`
	PasetoPublicKeyStr   string        `mapstructure:"PASETO_PUBLIC_KEY"`
	PasetoPrivateKeyStr  string        `mapstructure:"PASETO_PRIVATE_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	RedisAddress         string        `mapstructure:"REDIS_ADDRESS"`
	SESFromEmail         string        `mapstructure:"SES_FROM_EMAIL"`
	SESMailRegion        string        `mapstructure:"SES_MAIL_REGION"`
	AWSAccessKeyID     string `mapstructure:"AWS_ACCESS_KEY_ID"`
	AWSSecretAccessKey string `mapstructure:"AWS_SECRET_ACCESS_KEY"`
	AWSDefaultRegion   string `mapstructure:"AWS_DEFAULT_REGION"`
}

func LoadConfig(path string) (config Config, err error) {

	viper.AddConfigPath(path)  // tell it where the file is
	viper.SetConfigName("app") // tell it the file name
	viper.SetConfigType("env") // tell it the file type (could also be json, xml,...)

	//Override existing Env variable values with what's read from the file
	viper.AutomaticEnv()

	//Read in the config
	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
