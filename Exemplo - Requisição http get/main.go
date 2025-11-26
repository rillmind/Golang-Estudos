package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("https://google.com")

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999)
	// res.Body.Read(bs)
	// fmt.Println(string(bs))

	// or

	// Isso implementa copy com um Writer (os.Stdout) e um Reader (res.Body (coprpo da resposta)).
	// Writer é uma interface e Reader também.
	io.Copy(os.Stdout, res.Body)
}
