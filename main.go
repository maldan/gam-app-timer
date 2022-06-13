package main

import (
	"embed"
	"github.com/maldan/gam-app-timer/internal/app/timer/core"

	helloworld "github.com/maldan/gam-app-timer/internal/app/timer"
)

//go:embed frontend/dist/*
var frontFs embed.FS

//go:embed ring.mp3
var ring []byte

func main() {
	core.RingSound = ring
	helloworld.Start(frontFs) //
}
