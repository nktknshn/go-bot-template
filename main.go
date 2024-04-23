package main

import "github.com/nktknshn/go-bot-template/cli"

func main() {
	if err := cli.CmdRoot.Execute(); err != nil {
		panic(err)
	}
}
