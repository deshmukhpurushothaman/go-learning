package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	// fmt.Println(resp)

	// Pass byte slice to Read
	// bs := make([]byte, 999999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))
	lw := logWriter{}
	io.Copy(lw, resp.Body)

	// Write to Interface
	// io.Copy(os.Stdout, resp.Body)

}

type logWriter struct{}

func (logWriter) Writer(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
