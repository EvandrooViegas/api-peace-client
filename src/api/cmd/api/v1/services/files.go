package services

import (
	"os"
	"github.com/EvandrooViegas/utils"
)

func GetImageBuf(file string) ([]byte, error) {
	path, err := utils.GetAbsolutePath("public/images/" + file)
	if err != nil {
		return nil, err
	}
	buf, err := os.ReadFile(path)
	return buf, nil
}