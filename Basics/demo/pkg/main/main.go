package main

import (
	"basic/demo/pkg/display"
	"basic/demo/pkg/msg"
)

func main() {
	msg.Hi()
	display.Display("hello from display")
	msg.Exciting("an exciting message")
}