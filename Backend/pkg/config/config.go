package config

import (
	"errors"
	"github.com/goccy/go-json"
	"os"
	"path"
)

import "Killspiel/pkg/helper"

var (
	NoConfigFound = errors.New("no config file found")
	localPaths    = [...]string{"./config", "."}
	configName    = "config.json"
)

type Config struct {
	UserCollector UserCollect `json:"userCollector"`
}

func FindConfigPath() (string, error) {
	file, exists := os.LookupEnv("KILLSPIEL_CONFIG")
	if exists && ExistsAndFile(file) {
		return file, nil
	}
	for _, base := range globalPaths {
		for _, end := range []string{helper.EnvOrDefault("KILLSPIEL_CONFIG_PATH", "Killspiel")} {
			p := path.Join(base, end, configName)
			if ExistsAndFile(p) {
				return path.Join(base, end), nil
			}
		}
	}
	for _, pa := range localPaths {
		p := path.Join(pa, configName)
		if ExistsAndFile(p) {
			return pa, nil
		}
	}

	return "", NoConfigFound
}

func ExistsAndFile(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !stat.IsDir()
}

func GetConfig(p string) (config Config, err error) {
	var f *os.File
	f, err = os.Open(path.Join(p, configName))
	if err != nil {
		return
	}

	err = json.NewDecoder(f).Decode(&config)
	if err != nil {
		return
	}

	return
}
