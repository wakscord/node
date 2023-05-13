package discord

import (
	"github.com/valyala/fasthttp"
	"github.com/wakscord/new-wakscord-node/config"
)

func Initialize() {
	fasthttpClient = &fasthttp.Client{
		MaxConnsPerHost: config.Default.MaxConcurrent,
	}
}
