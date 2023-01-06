package util

import "log"

func HandleErrorStop(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
