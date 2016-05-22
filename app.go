package main

import "projects/fileServer/server/routes"

func main() {
	router := router.New(5000)
	router.Start()
}
