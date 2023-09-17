package services

import (
	"os"
	"github.com/EvandrooViegas/api-piece/cmd/api/v1/utils"
)

func GetImageBuf(file string) ([]byte, error) {
	path, err := utils.GetAbsolutePath("public/images/" + file)
	if err != nil {
		return nil, err
	}
	buf, err := os.ReadFile(path)
	return buf, nil
}