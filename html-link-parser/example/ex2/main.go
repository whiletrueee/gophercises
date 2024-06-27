package main

import (
	"fmt"
	link "gophercises/html-link-parser"
	"strings"
)

var exampleHTML = `
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Anchor Tag Example 2</title>
</head>
<body>
    <ul>
        <li>First item</li>
        <li><a href="https://www.example.com">Visit Example <!-- this is a comment --> </a></li>
        <li>Second item</li>
    </ul>
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
