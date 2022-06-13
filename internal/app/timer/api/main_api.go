package api

import (
	"github.com/maldan/gam-app-timer/internal/app/timer/core"
	"github.com/maldan/go-rapi/rapi_core"
)

type MainApi struct {
}

func (r MainApi) GetIndex(args ArgsEmpty) string {
	return "Test"
}

func (r MainApi) GetRing(ctx *rapi_core.Context) {
	ctx.IsSkipProcessing = true
	ctx.RW.Header().Add("Content-Type", "audio/mp3")
	ctx.RW.Write(core.RingSound)
}
