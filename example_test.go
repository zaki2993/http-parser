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

    fmt.Printf("Sheme:%s host:%s path:%s\n",
        u.Sheme, u.Host, u.Path)

    // Output:
    // Sheme:http host:zaki.com path:path
}
