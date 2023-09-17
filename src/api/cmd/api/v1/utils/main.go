package utils

import (
	"os"
	"path/filepath"
)

func GetAbsolutePath(path string) (string, error) {
	workingDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	absolutePath := filepath.Join(workingDir, path)
	return absolutePath, nil
}