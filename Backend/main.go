package main

import (
	"embed"
	"io/fs"
	"net/http"
)

var (
	//go:embed frontend_build/*
	frontendBuild embed.FS

	frontend, _ = fs.Sub(frontendBuild, "frontend_build")
)

func main() {

	http.Handle("/", http.FileServer(http.FS(frontend)))

	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		panic(err)
	}
}
