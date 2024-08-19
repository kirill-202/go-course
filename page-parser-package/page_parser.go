package page_parser

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"os"
	"strings"
)

const PATH string = "./page-parser-package/test.html"


func saveFileString(path string) ([]byte, error) {


	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()


	bytes, err := io.ReadAll(file) 
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, fmt.Errorf("error opening file: %w", err)
	}

	return bytes, nil
}


func visit(node *html.Node, words, pics *int) {

	if node.Type == html.TextNode {
		*words += len(strings.Fields(node.Data))
	} else if node.Type == html.ElementNode && node.Data == "img" {
		*pics++

	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		visit(c, words, pics)
	}
}

func countWordsAndImages(doc *html.Node) (int, int) {
	var words, pics int

	visit(doc, &words, &pics)

	return words, pics
}


func ReadHtml() {
    m_bytes, _ := saveFileString(PATH)
	doc, err := html.Parse(bytes.NewReader(m_bytes))
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse failed: %s\n", err)
	}

	words, pics := countWordsAndImages(doc)

	fmt.Printf("%d words and %d images\n", words, pics)
}