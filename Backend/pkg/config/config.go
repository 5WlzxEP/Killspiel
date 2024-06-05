package config

import (
	"cmp"
	"errors"
	"github.com/goccy/go-json"
	"io/fs"
	"os"
	"path"
)

var (
	NoConfigFound = errors.New("no config file found")
	localPaths    = [...]string{"./config", "."}
	configName    = cmp.Or(os.Getenv("ConfigName"), "config.json")
)

type Config struct {
	UserCollector UserCollect `json:"userCollector"`
	// location holds the path the config.json
	// <br>
	// is unexported to prevent json serialization,
	// and I'm too lazy to implementMarshalJSON and maintain it
	location  string
	Precision float64 `json:"precision"`
}

func (c *Config) GetLocation() string {
	return c.location
}

func (c *Config) Save() error {
	file, err := os.OpenFile(c.location, os.O_TRUNC|os.O_WRONLY, 0)
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	return err
}

func FindConfigPath() (string, error) {
	file, exists := os.LookupEnv("KILLSPIEL_CONFIG")
	if exists && ExistsAndFile(file) {
		return file, nil
	}
	for _, base := range globalPaths {
		for _, end := range []string{cmp.Or(os.Getenv("KILLSPIEL_CONFIG_PATH"), "Killspiel")} {
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

func GetConfig(p string) (*Config, error) {
	var config Config
	loc := path.Join(p, configName)
	f, err := os.Open(loc)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(f).Decode(&config)
	if err != nil {
		return nil, err
	}

	config.Precision = max(config.Precision, 0.01)
	config.location = loc
	return &config, nil
}

// Default creates the folder config in the current location and creates the default config
func Default() (string, error) {

	// check if the folder exists and create if not exists
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

	err = json.NewEncoder(f).Encode(&Config{Precision: 0.01})
	if err != nil {
		return "", err
	}

	return "./config", nil
}

func FindConfOrDefault() (string, error) {
	configPath, err := FindConfigPath()
	if errors.Is(err, NoConfigFound) {
		logger.Println("No config found, creating default.")
		configPath, err = Default()

		if err != nil {
			return "", err
		}
	} else if err != nil {
		return "", err
	}

	return configPath, nil
}
