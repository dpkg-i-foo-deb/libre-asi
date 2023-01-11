package util

import "log"

func HandleErorLog(err error) {
	if err != nil {
		log.Println(err)
	}
}

func HandleErrorStop(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
