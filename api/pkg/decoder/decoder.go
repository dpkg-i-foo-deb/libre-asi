package decoder

import (
	"encoding/json"
	"libre-asi-api/pkg/models"
	"libre-asi-api/pkg/util"
	"os"
)

func DecodeWorld() *[]models.Country {
	var world []models.Country

	reader, err := os.Open("world.json")

	util.HandleErrorStop(err)

	decoder := json.NewDecoder(reader)

	err = decoder.Decode(&world)

	util.HandleErrorStop(err)

	return &world
}
