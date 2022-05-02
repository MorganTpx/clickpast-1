package main

import (
	_ "clickpast/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"clickpast/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
