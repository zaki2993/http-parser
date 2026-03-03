package url_test

import (
	"fmt"
	"log"

	"github.com/zaki/url"
)

func ExampleParse() {
    u, err := url.Parse("https://zaki.com/path")
		if err != nil{
			log.Fatal(err)
		}
		u.Sheme = "http"
		fmt.Print(u)
    // Output:
    // Sheme:http Host:zaki.com Path:path
}
