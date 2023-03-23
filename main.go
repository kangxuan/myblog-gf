package main

import (
	_ "myblog-gf/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"myblog-gf/internal/cmd"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func main() {
	cmd.Main.Run(gctx.New())
}
