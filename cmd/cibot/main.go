package main

import (
	"github.com/spf13/pflag"

	"gitee.com/openeuler/ci-bot/pkg/cibot"
)

func main() {
	wh := cibot.NewWebHook()
	wh.AddFlags(pflag.CommandLine)
	wh.Run()
}
