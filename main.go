package main

import (
    "fmt"
    "log"
    "os"
    "net/http"

    "github.com/PuerkitoBio/goquery"
)

func isin(array []string, str string) bool {
	for _, value := range array {
		if value == str {
			return true
		}
	}

	return false
}

func main() {

    if len(os.Args) <= 1 {
        fmt.Println("No URL provided. Try: crtsh-cli kleiber.me")
        os.Exit(0)
    }

    url_to_test := os.Args[1]

    fmt.Println(fmt.Sprintf("Checking: %s\n", url_to_test))
     
    var domains []string

    resp, err := http.Get(fmt.Sprintf("https://crt.sh/?q=%%25.%s", url_to_test))
    if err != nil {
        print(err)
        return
    }

    defer resp.Body.Close()

    doc, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    doc.Find("tr").Each(func(i int, s *goquery.Selection) {
        common_name := s.Children().Eq(4).Text()

        if (common_name != "Common Name" && len(common_name) > 0) {
            if !isin(domains, common_name) {
                domains = append(domains, common_name)
            }
        }  

    })

    if len(domains) > 0 {
        for _, str := range domains {
            fmt.Println(str)
        }
    } else {
        fmt.Println("No further domains found!")
    }
}