package main

import (
	_ "libre-asi-api/database"
	"libre-asi-api/internal/cmd"
)

func main() {
	cmd.Execute()
}
