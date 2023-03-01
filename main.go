package main

import (
	"im/common"

	"github.com/gin-gonic/gin"

	//"github.com/go-gin-gonic"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()
	r = CollectRoute(r)

	panic(r.Run())
}
