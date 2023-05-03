package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// working with the Read function here
	// this is saying to give me a byte slice that can contain 99999 elements
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// -------
	// condensed version of the above
	// --------
	io.Copy(os.Stdout, resp.Body)

}
