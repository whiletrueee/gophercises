package main

import (
	"fmt"
	link "gophercises/html-link-parser"
	"strings"
)

var exampleHTML = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Blog Page Example</title>
</head>
<body>
    <div class="navbar">
        <a href="#home"><span>Home</span></a>
        <a href="#about"><span>About</span></a>
        <a href="#blog"><span>Blog</span></a>
        <a href="#contact"><span>Contact</span></a>
    </div>
    <div class="content">
        <h1>Welcome to My Blog</h1>
        <p>Discover the latest updates and insights on our <a href="#blog"><span>Blog page</span></a>.</p>
        <h2>Latest Posts</h2>
        <ul>
            <li><a href="#post1">How to Learn HTML in a Week</a></li>
            <li><a href="#post2">10 Tips for CSS Beginners</a></li>
            <li><a href="#post3">Understanding JavaScript Closures</a></li>
        </ul>
        <p>Don't miss our <a href="#newsletter"><span>monthly newsletter</span></a> for exclusive content.</p>
        <h2>Popular Tags</h2>
        <a href="#html"><span>#HTML</span></a>
        <a href="#css"><span>#CSS</span></a>
        <a href="#javascript"><span>#JavaScript</span></a>
        <h2>Follow Us</h2>
        <p>Stay connected through our social media channels:</p>
        <p>
            <a href="https://facebook.com"><span>Facebook</span></a> |
            <a href="https://twitter.com"><span>Twitter</span></a> |
            <a href="https://instagram.com"><span>Instagram</span></a>
        </p>
    </div>
    <div class="footer">
        <a href="#privacy"><li>Privacy Policy</li></a>
        <a href="#terms"><li>Terms of Service</li></a>
        <a href="#support"><li>Support</li></a>
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
