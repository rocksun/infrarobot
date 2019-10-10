package cfgreader

import (
	"fmt"
	"strings"

	"github.com/tealeg/xlsx"
)

func NewCFGReader(dict KeyDict) CFGReader {
	outputData := make([]map[string](map[string]string), 0)

	return CFGReader{outputData, dict}
}

type CFGReader struct {
	outputData []map[string](map[string]string)
	dict       KeyDict
}

func (cfgreader *CFGReader) AddData(data map[string](map[string]string)) {
	cfgreader.outputData = append(cfgreader.outputData, data)
}

func (cfgreader *CFGReader) GetData() []map[string](map[string]string) {
	return cfgreader.outputData
}

func (cfgreader *CFGReader) ReadFile(excelFile string) {
	xlFile, err := xlsx.OpenFile(excelFile)
	if err != nil {
		fmt.Println("Error Open File: " + excelFile)
		return
	}

	sheetDatas := make(map[string]([]map[string]string))

	for _, sheet := range xlFile.Sheets {
		// fmt.Println(sheet.Name)
		sheetName := strings.ToLower(strings.TrimSpace(sheet.Name))
		sheetReader := GetSheetReader(sheetName, cfgreader.dict)
		if sheetReader != nil {
			fmt.Println("--- Read Sheet '"+sheetReader.SheetType+"' :", sheetName)
			sheetDatas[sheetReader.SheetType] = sheetReader.ReadSheet(sheet)
		}
	}

	for _, hostMap := range sheetDatas["host"] {

		for key, dataMaps := range sheetDatas {
			if key != "host" {
				for _, data := range dataMaps {
					singleData := make(map[string](map[string]string))
					// fmt.Println(key[:len(key)-1])
					// fmt.Println(data)
					// fmt.Println(hostMap)
					singleData["host"] = hostMap
					singleData[key] = data
					cfgreader.AddData(singleData)
					// fmt.Println(singleData)
				}
			}
		}
		// fmt.Println(hostMaps["hostname"])
		// singleData["host"] = hostMap
	}
}
