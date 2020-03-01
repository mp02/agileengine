package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/net/html"
)

func main() {

	if len(os.Args) < 5 {
		fmt.Println("Not enough arguments.")
		return
	}

	originalFile := os.Args[1]
	otherFile := os.Args[2]
	key := os.Args[3]
	value := os.Args[4]
	fmt.Println("BaseHTML", originalFile)
	fmt.Println("HTMLToBeCompared", otherFile)
	fmt.Println("FilerKey: ", key, "FilterValue: ", value)

	original, err := ioutil.ReadFile(originalFile)
	if err != nil {
		panic(err)
	}

	other, err := ioutil.ReadFile(otherFile)
	if err != nil {
		panic(err)
	}

	lement := getOriginalElement(original, key, value)
	if lement == nil {
		fmt.Println("No filters found.")
	} else {
		fmt.Println("path: ", path(findPathWereChanged(other, lement)))
	}

}

func path(elemento *html.Node) string {
	var path string
	if elemento == nil {
		fmt.Println("elemento vacio")
	} else {
		var f func(*html.Node)
		f = func(n *html.Node) {
			if n.Parent != nil {
				path = n.Data + "/" + path
				f(n.Parent)
			}
		}
		f(elemento)
	}
	return path
}

func getOriginalElement(data []byte, key, value string) *html.Node {
	var returned *html.Node

	original, err := html.Parse(bytes.NewReader(data))
	if err != nil {
		// ...
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, element := range n.Attr {
				if element.Key == key && element.Val == value {
					returned = n
				}

			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(original)
	return returned
}

func findPathWereChanged(data []byte, filter *html.Node) *html.Node {
	var returned *html.Node
	original, err := html.Parse(bytes.NewReader(data))
	if err != nil {
		// ...
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode { //&& n.Data == "a" {
			for _, element := range n.Attr {
				for _, filter := range filter.Attr {
					if element.Key == filter.Key && element.Val == filter.Val {
						returned = n
					}

				}

			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(original)
	return returned
}
