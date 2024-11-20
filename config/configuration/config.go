package configuration

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
	"regexp"
	"strings"
)

var Config = configurationInit()

type configuration struct {
	Env string `yaml:""`
	Bot bot    `yaml:"bot"`
}
type bot struct {
	AppId       uint64 `yaml:"app-id"`
	AccessToken string `yaml:"access-token"`
}

func configurationInit() *configuration {
	file, err := os.ReadFile("./config/dev.yml")
	if err != nil {
		logrus.Error("读取配置文件异常", err)
		panic(err)
	}

	config := &configuration{}
	er := yaml.NewDecoder(strings.NewReader(expand(string(file), os.Getenv))).Decode(config)
	if er != nil {
		logrus.Error("解析配置文件异常", err)
		panic(err)
	}

	return config
}

// expand 使用正则进行环境变量展开
// os.ExpandEnv 字符 $ 无法逃逸
// https://github.com/golang/go/issues/43482
func expand(s string, mapping func(string) string) string {
	r := regexp.MustCompile(`\${([a-zA-Z_]+[a-zA-Z0-9_:/.]*)}`)
	return r.ReplaceAllStringFunc(s, func(s string) string {
		s = strings.Trim(s, "${}")
		before, after, ok := strings.Cut(s, ":")
		m := mapping(before)
		if ok && m == "" {
			return after
		}
		return m
	})
}
