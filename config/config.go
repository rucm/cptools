package config

// Data :
type Data struct {
	Common  Common  `yaml:"common"`
	AtCoder AtCoder `yaml:"atcoder"`
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

func init() {

}
