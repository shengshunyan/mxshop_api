package utils

import "github.com/spf13/viper"

const ENV_KEY = "MXSHOP_ENV"

func GetEnv() string {
	err := viper.BindEnv(ENV_KEY)
	if err != nil {
		panic(err)
	}
	env := viper.GetString(ENV_KEY)

	return env
}
