package conf

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

//init config infomation
func InitConfig() {
	workDir, err := os.Getwd()
	if err != nil {
		log.Fatal("get config directory err ", err)
	}
	log.Printf("app run directory %s\n", workDir)
	viper.SetConfigName("config")            //setting config name
	viper.SetConfigType("yaml")              //setting config type
	viper.AddConfigPath(workDir + "/config") //setting config path
	if err = viper.ReadInConfig(); err != nil {
		log.Fatal("ReadInConfig err ", err)
	}
}
