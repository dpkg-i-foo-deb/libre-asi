package main

import (
	"libre-asi-api/cmd"
	_ "libre-asi-api/database"
)

func main() {
	cmd.Execute()
}
