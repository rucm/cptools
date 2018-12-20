package config

import (
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// Data :
type Data struct {
	Common  Common  `yaml:"common"`
	AtCoder AtCoder `yaml:"atcoder"`
}

// Config :
var Config = defaultData

func init() {

	fileExists := isFileExists(configFile)

	if fileExists {
		buf, err := ioutil.ReadFile(configFile)

		if err != nil {
			log.Fatal(err)
		}

		err = yaml.Unmarshal(buf, &Config)

		if err != nil {
			log.Fatal(err)
		}

	} else {
		yaml.Marshal(&defaultData)
	}
}

const configFile = "cptools_config.yaml"

var defaultData = Data{
	Common{
		SessionFile: "cptools.session",
	},
	AtCoder{
		LoginURL:       "https://practice.contest.atcoder.jp/login",
		ContestBaseURL: "https://atcoder.jp/contests/<contestID>/tasks",
	},
}

func isFileExists(filename string) bool {

	_, err := os.Stat(filename)
	return err == nil
}
