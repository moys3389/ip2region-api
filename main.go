package main

import (
	"log"

	"github.com/moys3389/ip2region-api/app"
	"github.com/samber/do/v2"
)

func main() {
	if err := do.MustInvoke[*app.App](nil).Start(); err != nil {
		log.Fatal(err)
	}
}
