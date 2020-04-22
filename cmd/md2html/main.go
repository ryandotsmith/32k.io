package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	gmhtml "github.com/yuin/goldmark/renderer/html"
	"golang.org/x/net/html"
)

var (
	dir    string
	layout []byte
	header = []byte("*[← home page](/)*\n")
)

func main() {
	s := flag.Bool("s", false, "serve html")
	d := flag.String("d", os.Getenv("K23")+"/r", "markdown directory")
	flag.Parse()
	dir = *d

	do()
	if *s {
		var handler http.Handler
		handler = http.FileServer(http.Dir(dir + "/site"))
		handler = doHandler(handler)
		err := http.ListenAndServe(":8080", handler)
		check(err)
	}
}

func doHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		do()
		next.ServeHTTP(w, req)
	})
}

func do() {
	var err error
	layout, err = ioutil.ReadFile(dir + "/layout.html")
	check(err)

	files, err := ioutil.ReadDir(dir + "/docs")
	check(err)

	for i := range files {
		srcName := files[i].Name()
		if srcName[len(srcName)-3:len(srcName)] != ".md" {
			continue
		}
		srcBytes, err := ioutil.ReadFile(dir + "/docs/" + srcName)
		check(err)

		if srcName != "index.md" {
			srcBytes = append(header, srcBytes...)
		}

		destBytes := title(wrap(convert(srcBytes)))
		destName := dir + "/site/" + srcName[0:len(srcName)-3]
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
	md := goldmark.New(
		goldmark.WithRendererOptions(
			gmhtml.WithHardWraps(),
			gmhtml.WithXHTML(),
			gmhtml.WithUnsafe(),
		),
		goldmark.WithExtensions(extension.GFM, extension.Typographer),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
	)
	var buf bytes.Buffer
	err := md.Convert(source, &buf)
	check(err)
	return buf.Bytes()
}
