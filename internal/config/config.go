package config

import (
	"gopkg.in/yaml.v3"
	"jing/pkg/util"
	"log"
	"os"
)

const (
	PornKind = "porn"
)

type Tv struct {
	Kind             string
	Name             string
	Url              string
	Domain           string
	HomePrefix       string `yaml:"homePrefix"`
	Heads            map[string]string
	DownloadNumLimit int `yaml:"downloadNumLimit"`
	DownloadSize     int `yaml:"downloadSize"`
	MaxPage          int `yaml:"maxPage"`
}

func ListTv(kind string, name string) Tv {
	var result Tv
	workDir := util.GetWorkDir()

	data, err := os.ReadFile(workDir + "/config/tv.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	// 解析YAML到结构体中
	var tvs []Tv
	err = yaml.Unmarshal(data, &tvs)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	for _, tv := range tvs {
		if kind == tv.Kind && name == tv.Name {
			result = tv
		}
	}
	return result
}
