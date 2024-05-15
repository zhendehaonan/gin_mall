package main

import (
	"gin_mall/conf"
	"gin_mall/routes"
)

func main() {
	conf.Init()
	router := routes.NewRouter()
	_ = router.Run(conf.HttpPort)

}
