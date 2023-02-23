package main

import (
	"libre-asi-api/cmd"
	_ "libre-asi-api/database"
	_ "libre-asi-api/routes"
)

func main() {
	cmd.Execute()
}
