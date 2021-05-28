package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/wamuir/go-xslt"
	"io"
	"log"
	"net/http"
)

const ServerPort = 8090

//var xmlLengthLarge = 195222279
var xmlLength1MB = 991906
var xslLength = 1027

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

func doXslt(docURL string, styleURL string) ([]byte, error) {
	// doc is the document to be transformed
	// style is the XSL stylesheet to be used for transformation
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
	doc := buf.Bytes()
	if len(doc) != xmlLength1MB {
		panic("wrong length XML")
	}

	var buf2 bytes.Buffer
	response2, err := client.Get(styleURL)
	if err != nil {
		panic(err)
	}
	defer response2.Body.Close()
	_, err = io.Copy(&buf2, response2.Body) //use package "io" and "os"
	if err != nil {
		panic(err)
	}
	style := buf2.Bytes()
	if len(style) != xslLength {
		panic("wrong length XSL")
	}

	// create a new stylesheet from style
	xs, err := xslt.NewStylesheet(style)
	if err != nil {
		panic(err)
	}
	defer xs.Close()

	// transform the document using the style
	res, err := xs.Transform(doc)
	if err != nil {
		panic(err)
	}

	// print the result of the transformation
	return res, nil
}

func GetParams(w http.ResponseWriter, r *http.Request) {
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
	http.HandleFunc("/", GetParams)
	http.HandleFunc("/xslt", GetParams)
	http.HandleFunc("/headers", headers)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", ServerPort), nil); err != nil {
		log.Fatal(err)
	}
}
