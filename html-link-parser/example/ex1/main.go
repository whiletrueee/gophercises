package main

import (
	"fmt"
	link "gophercises/html-link-parser"
	"strings"
)

var exampleHTML = `
<html>
<body>

<h1>My First Heading</h1>
<a href="/other-page">My 
first 



paragraph.</a>

</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHTML)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)

}
