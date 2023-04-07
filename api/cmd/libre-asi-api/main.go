package main

import (
	"libre-asi-api/internal/cmd"
	_ "libre-asi-api/pkg/database"
)

func main() {
	cmd.Execute()
}
