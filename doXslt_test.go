package main

import (
	"fmt"
	"log"
	"testing"
)

const serverURL = "https://localhost:8090/"
const testXML = "https://localhost/xml/pubmed21n0001.xml"
const testXSL = "https://localhost/xslt/biostar.xsl"

//func TestMain(m *testing.M) {}

/*func XXXTestGetParams(t *testing.T) {
	base, err := url.Parse(serverURL)
	if err != nil {
		return
	}
	var r *http.Request
	r.Method = "POST"
	r.URL, _ = url.Parse(serverURL)
	//...

	//	var w http.ResponseWriter

	params := url.Values{}
	params.Add("doc", testXML)
	params.Add("style", testXSL)
	base.RawQuery = params.Encode()

	response, err := http.Get(base.String())
	if err != nil {
		log.Fatalln(err)
		fmt.Println("Error when sending request to the server")
		return
	}
	defer response.Body.Close()

	respBody, _ := ioutil.ReadAll(response.Body)

	fmt.Println(response.Status)
	fmt.Println(string(respBody))
}
*/
func TestDoXslt(t *testing.T) {
	result, err := doXslt(testXML, testXSL)
	if err != nil {
		log.Fatalln(err)
		fmt.Println("Error when sending request to the server")
		return
	}

	fmt.Println(string(result))
}
