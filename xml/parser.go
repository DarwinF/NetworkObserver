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
	"NetworkObserver/logger"
	"NetworkObserver/reporter"
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

func ReadReport(rd *reporter.ReportData, filePath string) {
	xmlData, _ := ioutil.ReadFile(filePath)

	err := xml.Unmarshal(xmlData, &rd)

	if err != nil {
		logger.WriteString("There was an error unmarshalling the XML")
		return
	}
}
