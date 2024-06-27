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
    <title>Inline Links Example</title>
</head>
<body>
    <div class="content">
        <p>
            For more information about our services, visit our 
            <a href="https://www.example.com/services">Services page</a>.
        </p>
        <p>
            Learn more about our team on the 
            <a href="https://www.example.com/team">About Us page</a>.
        </p>
        <p>
            If you have any questions, feel free to 
            <a href="https://www.example.com/contact">Contact Us</a>.
        </p>
        <p>
            You can also read our 
            <a href="https://www.example.com/blog">latest blog posts</a> for updates and news.
        </p>
    </div>
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
