package main 

import (
	fetch "github.com/FriendlyUser/tex-cite/pkg/fetch"
	"flag"
	"fmt"
)

func main() {
	var svar string
	flag.StringVar(&svar, "isbn", "9781451648546", "ISBN 13 code search")
	flag.Parse()
	fmt.Println(svar)
	fetch.GetBook("9781451648546")
}