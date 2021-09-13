package main

import (
	"freado/core"
	"freado/initialize"
	"freado/initialize/global"
)

type User struct {
	name string
}

var user User

func main() {
	global.ORMER = initialize.GormRouter()
	global.VIPER = core.Viper()
	core.RunWindowsServer()
}
