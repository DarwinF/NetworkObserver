//--------------------------------------------
// xml/parser.go
//
// Reads xml into structs for displaying in
// webpages.
// Writes data into xml for storing in xml
// files to later be accessed.
//--------------------------------------------

package xml

import (
	"encoding/xml"
)

type structFill func(data map[string]string) interface{}

func Read(filler structFill, filePath string) interface{} {
	// Unmarshall the file contents
	data := unpackXML(filePath)

	// pass the map to the filler function
	// return the struct created
	return filler(data)
}

func unpackXML(path string) map[string]string {
	return nil
}
