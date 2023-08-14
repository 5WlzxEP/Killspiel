package config

import (
	"errors"
	"os"
	"path"
)

import "Killspiel/pkg/helper"

var (
	NO_CONFIG_FOUND = errors.New("no config file found")
	localPaths      = [...]string{"./config", "."}
	configName      = "config.json"
)

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

	return "", NO_CONFIG_FOUND
}

func ExistsAndFile(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !stat.IsDir()
}

type Config struct {
}
