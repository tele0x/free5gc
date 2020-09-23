package factory

import (
        "fmt"
        "io/ioutil"
	"free5gc/src/nrf/logger"
        "gopkg.in/yaml.v2"
)

type (
        Config struct {
		MongoDBName string `yaml:"MongoDBName"`
                MongoDBUrl string `yaml:"MongoDBUrl"`
                RanIP string `yaml:"RanIP"`
                AmfIP string `yaml:"AmfIP"`
		UpfIP string `yaml:"UpfIP"`
        }
)


var TestingConfig Config

func checkErr(err error) {
        if err != nil {
                err = fmt.Errorf("[Configuration] %s", err.Error())
                logger.AppLog.Fatal(err)
        }
}

func InitConfigFactory(f string) {
        content, err := ioutil.ReadFile(f)
        checkErr(err)

        TestingConfig = Config{}
        err = yaml.Unmarshal([]byte(content), &TestingConfig)
        checkErr(err)
}
