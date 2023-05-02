package main

import (
	"libre-asi-api/internal/cmd"
	"libre-asi-api/internal/setup"
	_ "libre-asi-api/pkg/database"
)

func main() {
	setup.CheckPreconditions()
	cmd.Execute()
}
