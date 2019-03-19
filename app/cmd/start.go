package main

import "github.com/joshiomkarj/inMemoryServer/app/pkg/api"

func main() {
	app := &api.App{}
	app.Initialize()
	app.Run()
}
