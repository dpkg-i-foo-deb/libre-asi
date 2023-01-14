package database

import (
	"libre-asi-api/decoder"
	"libre-asi-api/models"
	"libre-asi-api/util"
	"log"

	"gorm.io/gorm"
)

// Check if the world data needed exists
//This data comes from https://github.com/dr5hn/countries-states-cities-database

func checkWorld() {

	var c models.Country

	log.Println("Checking if world data exists")

	res := DB.Take(&c)

	if res.Error == gorm.ErrRecordNotFound {
		log.Println("Creating world in the background...")
		go DB.Transaction(createWorld)

	} else {
		util.HandleErrorStop(res.Error)
	}

}

func createWorld(db *gorm.DB) error {

	world := *decoder.DecodeWorld()

	for i := range world {
		res := DB.Omit("States", "Cities").Create(&world[i])

		if res.Error != nil {
			util.HandleErrorStop(err)
		}

	}

	for i := range world {
		for j := range world[i].States {
			world[i].States[j].CountryID = world[i].ID
			res := DB.Omit("Cities").Create(&world[i].States[j])

			if res.Error != nil {
				util.HandleErrorStop(err)
			}
		}
	}

	for i := range world {
		for j := range world[i].States {
			for k := range world[i].States[j].Cities {
				world[i].States[j].Cities[k].StateID = world[i].States[j].ID
				res := DB.Create(&world[i].States[j].Cities[k])

				if res.Error != nil {
					util.HandleErrorStop(err)
				}
			}
		}
	}

	log.Println("Finished creating world")

	return DB.Error
}
