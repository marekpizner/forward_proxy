package main

import (
	"encoding/json"
	"fmt"

	"go_proxy/cache"
)

type ExampleConvertToByteArray struct {
	Name    string
	SurName string
}

func main() {
	// P.RunProxy()
	example := ExampleConvertToByteArray{
		Name:    "James",
		SurName: "Camara",
	}
	var exampleBytes []byte
	var err error

	exampleBytes, err = json.Marshal(example)
	if err != nil {
		print(err)
		return
	}

	cache := cache.NewInMemorry("./cache_dir")

	cache.Add("aa", exampleBytes)
	getData := cache.Get("aa")
	fmt.Println(getData)

}
