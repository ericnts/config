package config

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

var (
	Options     *Project
	CurrentFile string //编译时指定配置文件  go build -ldflags "-X gitlab.epshealth.com/library-go/config.CurrentFile=conf/config-test.yml"
	raw         []byte
)

func init() {
	configPath := findConfigPath()
	bytes, err := os.ReadFile(configPath)
	if err != nil {
		panic(errors.New(fmt.Sprintf("配置文件[%s]加载失败，%v", configPath, err)))
	}
	raw = bytes

	project, err := Load[Project]("project")
	if err != nil {
		panic(errors.New(fmt.Sprintf("配置文件[%s]解析失败，%v", configPath, err)))
	}
	Options = &project
	fmt.Printf("配置文件加载完成，%v\n", configPath)
}

func Load[T any](tag string) (T, error) {
	data := make(map[string]T, 1)
	err := yaml.Unmarshal(raw, &data)
	res := data[tag]
	return res, err
}

// 获取配置文件真实路径
func findConfigPath() string {
	finalPath := CurrentFile
	// 没有指定配置文件时，通过层级递归向上查找
	if len(finalPath) == 0 {
		confPath, err := os.Getwd()
		if err != nil {
			panic(errors.New(fmt.Sprintf("获取当前路径失败，%v", err)))
		}
		for len(confPath) > 0 {
			configFilePath := filepath.Join(confPath, "conf", "config.yml")
			if _, err := os.Stat(configFilePath); err == nil {
				finalPath = configFilePath
				break
			}

			configFilePath = filepath.Join(confPath, "config.yml")
			if _, err := os.Stat(configFilePath); err == nil {
				finalPath = configFilePath
				break
			}

			newPath := filepath.Dir(confPath)
			if newPath == confPath {
				break
			}
			confPath = newPath
		}
	}
	if len(finalPath) == 0 {
		panic(errors.New(fmt.Sprint("配置文件加载失败，没有找到配置文件")))
	}
	return finalPath
}
