package utils

import (
	"os"
	"path/filepath"
	"github.com/gofor-little/env"

)

func GetAbsolutePath(path string) (string, error) {
	workingDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	absolutePath := filepath.Join(workingDir, path)
	return absolutePath, nil
}

func LoadEnvVariable(key string) (string, error) {
	appEnv := os.Getenv("APP_ENV")
	switch appEnv {
	case "dev":
		if err := env.Load(".env.local"); err != nil {
			return "", err
		}
	default:
		if err := env.Load(".env.prod"); err != nil {
			return "", err
		}
	}
	return env.Get(key, "not found"), nil
}