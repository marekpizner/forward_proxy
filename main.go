package main

import (
	"fmt"

	"go_proxy/cache"
	"go_proxy/proxy"
)

type ExampleConvertToByteArray struct {
	Name    string
	SurName string
}

func main() {
	cache := cache.NewInMemorry(cacheDir)
	blckList := proxy.NewBlacklist(blacklistPath)

	proxy.RunProxy(cache, blckList)

	fmt.Scanln()
	fmt.Println("done")
}
