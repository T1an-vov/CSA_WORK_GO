package main

import (
	"CSA_Work_Go/dao"
	"CSA_Work_Go/router"
)

func main() {
     err:=dao.SQLInit()
	if err != nil {
		panic(err)
	}
	defer dao.Close()
	 r:=router.SetRouter()
	 r.Run(":8080")
}
