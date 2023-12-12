package main

import (
	"online-course-assignment-go/router"
	"online-course-assignment-go/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	r := router.Router()

	r.Run("0.0.0.0:8081")
}
