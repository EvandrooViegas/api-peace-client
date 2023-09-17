package services

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func GetAllArcs() ([]Arc, error) {
	path := "public/data/arcs.json"
	workingDir, err := os.Getwd()
	if err != nil {
		return make([]Arc, 0), err
	}
	absolutePath := filepath.Join(workingDir, path)
	file, err := os.Open(absolutePath)
	defer file.Close()
	if err != nil {
		return make([]Arc, 0), err
	}

	var arcs []Arc
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&arcs)
	if err != nil {
		return make([]Arc, 0), err
	}
	return arcs, nil
}