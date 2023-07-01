package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func LoadConfig() (Configure, error) {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error in getting fs: ", err)
		return Configure{}, err
	}

	if _, err := os.Stat(filepath.Join(wd, "config.json")); os.IsNotExist(err) {
		return Configure{}, errors.New("config file not exist")
	}

	ctx, err := os.ReadFile(filepath.Join(wd, "config.json"))
	if err != nil {
		return Configure{}, err
	}

	var cfg = Configure{}
	err = json.Unmarshal(ctx, &cfg)
	if err != nil {
		return Configure{}, err
	}

	return cfg, nil
}
