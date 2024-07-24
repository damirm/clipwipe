package main

import (
	"context"
	"log"

	"golang.design/x/clipboard"
)

func main() {
	err := clipboard.Init()
	if err != nil {
		log.Fatalf("failed to initialize clipboard: %#v", err)
	}

	ctx := context.Background()
	ch := clipboard.Watch(ctx, clipboard.FmtText)
	for data := range ch {
		res := cleaned(string(data))
		clipboard.Write(clipboard.FmtText, []byte(res))
	}
}

func cleaned(text string) string {
	return text
}
