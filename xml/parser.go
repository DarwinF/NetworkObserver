//--------------------------------------------
// xml/parser.go
//
// Reads xml into structs for displaying in
// webpages.
// Writes data into xml for storing in xml
// files to later be accessed.
//--------------------------------------------

package main

//package xml

import (
	"NetworkObserver/reporter"
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

func main() {
	Read(nil, "")
}

type structFill func(data map[string]string) interface{}

func Read(filler structFill, filePath string) interface{} {
	// Unmarshall the file contents
	//data := unpackXML(filePath)
	unpackXML("../sample_report.xml")

	// pass the map to the filler function
	// return the struct created
	//return filler(data)
	return nil
}

func unpackXML(path string) {
	xmlData, _ := ioutil.ReadFile(path)

	rd := reporter.ReportData{}

	err := xml.Unmarshal(xmlData, &rd)

	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(rd)
}
