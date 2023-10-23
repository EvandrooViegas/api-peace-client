package services

import (
	"encoding/json"
	"os"

	"github.com/EvandrooViegas/types"
	"github.com/EvandrooViegas/utils"
)


func GetAllArcs(addr string) ([]types.Arc, error) {
	path, err := utils.GetAbsolutePath( "public/data/arcs.json")
	if err != nil {
		return make([]types.Arc, 0), err
	}

	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return make([]types.Arc, 0), err
	}

	var arcs []types.Arc
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&arcs)
	if err != nil {
		return make([]types.Arc, 0), err
	}

	for idx := range arcs {
		arcs[idx].Image = addr + "/image" + arcs[idx].Image 
	}
	return arcs, nil
}