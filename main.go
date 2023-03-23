package main

import (
	_ "myblog-gf/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"myblog-gf/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
