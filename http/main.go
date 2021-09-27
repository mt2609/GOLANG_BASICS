package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	lw := logWriter{}
	io.Copy(lw, resp.Body)
	//lw is a value that implements WI, coz it has a fn Write-
	//that take a byte slice and returns int, err

}
func (logWriter) Write(bs []byte) (int, error) {
	//return 1, nil
	fmt.Println(string(bs))
	fmt.Println("just wroyte many bytes:", len(bs))

	return len(bs), nil
}
