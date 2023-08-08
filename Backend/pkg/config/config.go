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
	configName      = "config.yaml"
)

func FindConfigPath() (string, error) {
	file, exists := os.LookupEnv("KILLSPIEL_CONFIG")
	if exists && existsAndFile(file) {
		return file, nil
	}
	for _, base := range globalPaths {
		for _, end := range []string{helper.EnvOrDefault("KILLSPIEL_CONFIG_PATH", "Killspiel")} {
			p := path.Join(base, end, configName)
			if existsAndFile(p) {
				return path.Join(base, end), nil
			}
		}
	}
	for _, pa := range localPaths {
		p := path.Join(pa, configName)
		if existsAndFile(p) {
			return pa, nil
		}
	}

	return "", NO_CONFIG_FOUND
}

func existsAndFile(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !stat.IsDir()
}

type Config struct {
}
