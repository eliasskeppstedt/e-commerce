package configs

import (
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DB MysqlConfig
}

type MysqlConfig struct {
	User string
	Pass string
	Url  string
	Name string
}

func LoadDbCfg() (Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	cfg := Config{
		DB: MysqlConfig{
			User: os.Getenv("DBUSER"),
			Pass: os.Getenv("DBPASS"),
			Url:  os.Getenv("DBURL"),
			Name: os.Getenv("DBNAME"),
		},
	}

	return cfg, nil
}

func LoadDsnCfg(mysqlConfig *MysqlConfig) *mysql.Config {
	// Capture connection properties.
	cfg := mysql.NewConfig()
	cfg.User = mysqlConfig.User
	cfg.Passwd = mysqlConfig.Pass
	cfg.Net = "tcp"
	cfg.Addr = mysqlConfig.Url
	cfg.DBName = mysqlConfig.Name
	return cfg
}
