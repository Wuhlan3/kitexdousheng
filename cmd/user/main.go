package main

import (
	user "kitexdousheng/kitex_gen/user/usersrv"
	"log"
)

func main() {
	svr := user.NewServer(new(UserSrvImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
