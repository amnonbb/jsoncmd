package main

import "os"

func main() {
	a := App{}
	a.Initialize()
	a.Run(":" + os.Getenv("WEB_PORT"))
}