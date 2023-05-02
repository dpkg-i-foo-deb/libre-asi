package main

import (
	"libre-asi-api/internal/cmd"
	"libre-asi-api/internal/global"
	"libre-asi-api/internal/setup"
	_ "libre-asi-api/pkg/database"
)

func main() {
	setup.CheckPreconditions()
	global.InitGlobalAttributes()
	cmd.Execute()
}
