package helpers

import (
	"github.com/Nivelian/codete-webscraping/model"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"os"
)

func InitLog() {
	logrus.SetFormatter(&prefixed.TextFormatter{
		TimestampFormat: "02.01.2006 15:04:05",
		FullTimestamp:   true,
		ForceFormatting: true,
	})
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(io.MultiWriter(os.Stdout))
}

func GetConfig() *model.Config {
	configBytes, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(LogErr(err, "Failed to read file"))
	}

	config := new(model.Config)
	if err := yaml.Unmarshal(configBytes, config); err != nil {
		panic(LogErr(err, "Failed to parse yaml"))
	}

	return config
}
