package main

import (
	relation "kitexdousheng/kitex_gen/relation/relationsrv"
	"log"
)

func main() {
	svr := relation.NewServer(new(RelationSrvImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
