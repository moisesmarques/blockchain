package main

import "os"

func main() {

	os.Setenv("NODE_ID", "123")

	cli := CLI{}
	cli.Run()
}
