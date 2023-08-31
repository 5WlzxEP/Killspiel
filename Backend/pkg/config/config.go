package config

import (
	"errors"
	"github.com/goccy/go-json"
	"io/fs"
	"log"
	"os"
	"path"
)

import "Killspiel/pkg/helper"

var (
	NoConfigFound = errors.New("no config file found")
	localPaths    = [...]string{"./config", "."}
	configName    = helper.EnvOrDefault("ConfigName", "config.json")
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

// Default creates the folder config in the current location and creates the default config
func Default() (string, error) {

	// check if folder exists and create if not exists
	stat, err := os.Stat("./config")
	if errors.Is(err, fs.ErrNotExist) {
		err = os.Mkdir("./config", 0744)
		if err != nil {
			return "", err
		}
	} else if err != nil {
		return "", err
	} else {
		if !stat.IsDir() {
			return "", errors.New("config exists, but is file")
		}
	}

	f, err := os.OpenFile(path.Join("./config", configName), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return "", err
	}
	defer f.Close()

	err = json.NewEncoder(f).Encode(&Config{})
	if err != nil {
		return "", err
	}

	return "./config", nil
}

func FindConfOrDefault() (string, error) {
	configPath, err := FindConfigPath()
	if errors.Is(err, NoConfigFound) {
		log.Println("No config found, creating default.")
		configPath, err = Default()

		if err != nil {
			return "", err
		}
	} else if err != nil {
		return "", nil
	}

	return configPath, nil
}
