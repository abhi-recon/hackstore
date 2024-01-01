package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
    "os"

	"gopkg.in/yaml.v2"

	_ "github.com/go-sql-driver/mysql"
)

// DBConfig holds the database configuration
type DBConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
}

var db *sql.DB

func ConnectDB() {
	config, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Username, config.Password, config.Host, config.Port, config.Database)

	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func loadConfig() (*DBConfig, error) {
	var config DBConfig

	yamlFile, err := ioutil.ReadFile(os.Getenv("HOME")+"/.config/hackstore/hackstore-config.yml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
