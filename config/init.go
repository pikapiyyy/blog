package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	projectName := "blog"
	getConfig(projectName)
}

func getConfig(projectName string) {
	// name of config file (without extension)
	viper.SetConfigName("config")
	// path to look for the config file in
	viper.AddConfigPath(fmt.Sprintf("$GOPATH/src/%s", projectName))
	// Find and read the config file
	err := viper.ReadInConfig()
	// Handle errors reading the config file
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}

// GetMysqlConnectingString func
func GetMysqlConnectingString() string {
	usr := viper.GetString("mysql.user")
	pwd := viper.GetString("mysql.password")
	host := viper.GetString("mysql.host")
	db := viper.GetString("mysql.db")
	charset := viper.GetString("mysql.charset")

	return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=%s&parseTime=true", usr, pwd, host, db, charset)
}
