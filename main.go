package main

import "task/Delivery/routers"

func main() {
	router := routers.SetupRouter()
	router.Run(":8080")
}
