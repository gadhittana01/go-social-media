package helper

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/gadhittana01/socialmedia/config"
	"gopkg.in/yaml.v2"
)

func LoadConfig(c *config.GlobalConfig) {
	path := "config/social-media-http.yaml"
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	// change db host
	DBHost := os.Getenv("DB_HOST")
	if DBHost != "" {
		c.DB.Host = DBHost
	}
}
