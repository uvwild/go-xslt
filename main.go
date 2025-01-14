package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	xslt "github.com/wamuir/go-xslt"
)

const ServerPort = 8090

//var xmlLengthLarge = 195222279
//var xmlLength1MB = 991906
//var xslLength = 1027

var transCfg = &http.Transport{
	TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
}
var client = &http.Client{Transport: transCfg}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			_, err := fmt.Fprintf(w, "%v: %v\n", name, h)
			if err != nil {
				panic(err)
			}
		}
	}
}

func readDocument(docURL string) ([]byte, error) {
	var buf bytes.Buffer
	response, err := client.Get(docURL)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	_, err = io.Copy(&buf, response.Body) //use package "io" and "os"
	if err != nil {
		panic(err)
	}
	return buf.Bytes(), nil
}

func doXslt(docParm string, styleParm string) ([]byte, error) {
	// doc is the document to be transformed
	// style is the XSL stylesheet to be used for transformation
	var document []byte
	var stylesheet []byte

	if strings.HasPrefix(docParm, "<") {
		document = []byte(docParm)
	} else {
		doc, err := readDocument(docParm)
		if err != nil {
			panic(err)
		}
		document = doc
	}

	if strings.HasPrefix(docParm, "<") {
		stylesheet = []byte(docParm)
	} else {
		style, err := readDocument(styleParm)
		if err != nil {
			panic(err)
		}
		stylesheet = style
	}

	// create a new stylesheet from style
	xs, err := xslt.NewStylesheet(stylesheet)
	if err != nil {
		panic(err)
	}
	defer xs.Close()

	// transform the document using the style
	res, err := xs.Transform(document)
	if err != nil {
		panic(err)
	}

	// print the result of the transformation
	return res, nil
}

func GetUrlParams(w http.ResponseWriter, r *http.Request) {
	var doc string
	var style string
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case "GET":
		for k, v := range r.URL.Query() {
			log.Printf("%s: %s\n", k, v)
			switch k {
			case "doc":
				doc = v[0] //nolint:gofmt
			case "style":
				style = v[0] //nolint:gofmt
			default:
			}
		}
	case "POST":
		//reqBody, err := ioutil.ReadAll(r.Body)
		//fmt.Printf("%s", reqBody)
		r.ParseForm()
		doc = r.Form.Get("doc")     //nolint:gofmt    // will be "" if parameter is not set
		style = r.Form.Get("style") //nolint:gofmt
	default:
		w.WriteHeader(http.StatusNotImplemented)
		_, err := w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
		if err != nil {
			panic(err)
		}
	}
	if doc == "" {
		panic("missing parameter doc")
	}

	if style == "" {
		panic("missing parameter style")
	}

	_, err := w.Write([]byte(fmt.Sprintf("Received a %s request\n", r.Method)))
	if err != nil {
		panic(err)
	}

	res, err := doXslt(doc, style)
	if err != nil {
		panic(err)
	}

	// write the xsl transformation result back
	_, err = w.Write(res)
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", GetUrlParams)
	http.HandleFunc("/xslt", GetUrlParams)
	http.HandleFunc("/headers", headers)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", ServerPort), nil); err != nil {
		log.Fatal(err)
	}
}
