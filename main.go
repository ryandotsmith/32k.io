package main

import (
	"bytes"
	"io/ioutil"

	"golang.org/x/net/html"
	"github.com/russross/blackfriday"
)

var (
	layout []byte
	footer []byte
	header = []byte("*[r.32k.io](/)*\n")
)

func main() {
	var err error

	layout, err = ioutil.ReadFile("./layout.html")
	check(err)

	files, err := ioutil.ReadDir("./Documents")
	check(err)

	for i := range files {
		srcName := files[i].Name()
		if srcName[len(srcName)-3:len(srcName)] != ".md" {
			continue
		}
		srcBytes, err := ioutil.ReadFile("./Documents/" + srcName)
		check(err)

		destBytes := title(wrap(convert(append(append(header, srcBytes...), footer...))))
		destName := "./Site/" + srcName[0:len(srcName)-3]
		err = ioutil.WriteFile(destName, destBytes, 0644)
		check(err)

	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}


func title(body []byte) []byte {
	doc, err := html.Parse(bytes.NewReader(body))
	check(err)

	var title string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "h1" {
			title = n.FirstChild.Data
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return bytes.Replace(body, []byte("{{Title}}"), []byte(title), 1)
}

func wrap(body []byte) []byte {
	return bytes.Replace(layout, []byte("{{Body}}"), body, 1)
}

func convert(source []byte) []byte {
	htmlFlags := 0
	htmlFlags |= blackfriday.HTML_USE_XHTML
	htmlFlags |= blackfriday.HTML_USE_SMARTYPANTS
	htmlFlags |= blackfriday.HTML_SMARTYPANTS_FRACTIONS
	htmlFlags |= blackfriday.HTML_SMARTYPANTS_LATEX_DASHES
	htmlFlags |= blackfriday.HTML_SMARTYPANTS_DASHES
	renderer := blackfriday.HtmlRenderer(htmlFlags, "", "")

	extensions := 0
	extensions |= blackfriday.EXTENSION_NO_INTRA_EMPHASIS
	extensions |= blackfriday.EXTENSION_TABLES
	extensions |= blackfriday.EXTENSION_FENCED_CODE
	extensions |= blackfriday.EXTENSION_AUTOLINK
	extensions |= blackfriday.EXTENSION_STRIKETHROUGH
	extensions |= blackfriday.EXTENSION_SPACE_HEADERS
	extensions |= blackfriday.EXTENSION_HEADER_IDS
	extensions |= blackfriday.EXTENSION_LAX_HTML_BLOCKS
	extensions |= blackfriday.EXTENSION_AUTO_HEADER_IDS

	return blackfriday.Markdown(source, renderer, extensions)
}
