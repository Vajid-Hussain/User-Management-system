package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBUser     string `mapstructure:"DBUSER"`
	DBName     string `mapstructure:"DBNAME"`
	DBPassword string `mapstructure:"PASSWORD"`
	DBHost     string `mapstructure:"HOST"`
	DBPort     string `mapstructure:"PORT"`
	DBTimeZone string `mapstructure:"TIMEZONE"`
}

// var envs = []string{"USER", "DBNAME", "PASSWORD", "HOST", "PORT", "TIMEZONE"}

func LoadConfig() (*Config, error) {
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	// envData:=viper.AllSettings()

	// for key,val :=range envData{
	// 	fmt.Println(key ,  val)
	// }

	// for _, env := range envs {
	// 	// fmt.Println(data,env)
	// 	err := viper.BindEnv(env)
	// 	if err != nil {
	// 		fmt.Println("error at config bindenv", err)
	// 		return config, err
	// 	}
	// }

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	// if err := validator.New().Struct(&config); err != nil {
	// 	return config, err
	// }
	return &config, nil
}
