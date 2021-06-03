package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/wapc/wapc-go"
)

func main() {
	ctx := context.Background()
	code, err := ioutil.ReadFile("testdata/gateway.wasm")
	if err != nil {
		panic(err)
	}

	module, err := wapc.New(code, hostCall)
	if err != nil {
		panic(err)
	}
	module.SetLogger(wapc.Println)
	module.SetWriter(wapc.Print)
	defer module.Close()

	instance, err := module.Instantiate()
	if err != nil {
		panic(err)
	}
	defer instance.Close()

	result, err := instance.Invoke(ctx, "get_user_info_from_token", []byte("joaaaaaaaaaaaaaakiangjoaaaaaaaaaaaaaakiangjoaaaaaaaaaaaaaakiangjoaaaaaaaaaaaaaakiangjoaaaaaaaaaaaaaakiangjoaaaaaaaaaaaaaakiangjoaaaaaaaaaaaaaakiangjoaaaaaaaaaaaaaakiangjoaaaaaaaaaaaaaakiangjoaaaaaaaaaaaaaakiang"))
	if err != nil {
		panic(err)
	}

	fmt.Println(string(result))
}

func hostCall(ctx context.Context, binding, namespace, operation string, payload []byte) ([]byte, error) {
	// Route the payload to any custom functionality accordingly.
	// You can even route to other waPC modules!!!
	switch namespace {
	case "foo":
		switch operation {
		case "echo":
			return payload, nil // echo
		}
	}
	return []byte("default"), nil
}
