package main

import (
	"embed"

	helloworld "github.com/maldan/gam-app-timer/internal/app/timer"
)

//go:embed frontend/build/*
var frontFs embed.FS

func main() {
	helloworld.Start(frontFs) //
}
