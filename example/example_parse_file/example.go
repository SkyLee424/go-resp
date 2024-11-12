package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/Fusl/go-resp"
)

func main() {
	file, err := os.Open("./resp.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	srv := resp.NewServerWithFile(file)
	srv.SetRESP2Compat(true)
	defer srv.Close()

	for {
		args, err := srv.Next()

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}

		for _, arg := range args {
			fmt.Printf("%v ", string(arg))
		}
		fmt.Println()
	}
}
