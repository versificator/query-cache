package main

import (

	"io/ioutil"
	"fmt"
	"github.com/Vertamedia/chproxy/log"
	"testing"
	"regexp"
	"github.com/Jeffail/gabs"

)
//var JSONAttr = []string {"$select", "$from", "$preWhere", "$where", "$groupBy", "$having", "$globalRangePreWhere", "$globalRangeWhere"}

const path string = "./resource/query_samples/"

const query1 string = `select 1,1,toUInt32(Aid),toUInt32(0),uniqIf(IP, Event = 'requestRow'),uniqIf(IP, Event = 'requestData'),uniqIf(IP, Event = 'campaignImpression')
from stat.ipc_raw_10_13
prewhere EventDate='2017-10-13' and event in (('requestRow', 'requestData', 'campaignImpression'))
group by  key`

func clean(in string) string{
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatalf("Error: ",err)
	}
	return reg.ReplaceAllString(in, "")
}

func TestParseRoot(t *testing.T) {

	fileNames, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatalf("Error: ", err)
	}
	for _, fileName := range fileNames {
		f, err := ioutil.ReadFile(path + fileName.Name()) // Error handling elided for brevity.
		if err != nil {
			fmt.Print(err)
		}
		jsonParsed, _ := gabs.ParseJSON(f)
		child := jsonParsed.Search("$query").Data()
		fmt.Printf(" \n\n\n\n\n\n\n\n\n\n\n %s \n", toString(parseRootRawJSON(child)))


	}
}

func BenchmarkParseGhPost(b *testing.B) {

}