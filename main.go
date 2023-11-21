package main

import (
	_ "github.com/bitker/bitadmin/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/bitker/bitadmin/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
