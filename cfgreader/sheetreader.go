package cfgreader

import (
	"fmt"
	"strings"

	"github.com/tealeg/xlsx"
)

type SheetReader struct {
	SheetType string
	dict      KeyDict
	// ColumnTitleMaps map[string]string
}

func GetSheetReader(name string, dict KeyDict) *SheetReader {
	// sheetNames := map[string]string{
	// 	"用户":          "users",
	// 	"用户列表":        "users",
	// 	"users":       "users",
	// 	"主机":          "hosts",
	// 	"主机列表":        "hosts",
	// 	"hosts":       "hosts",
	// 	"文件系统":        "filesystems",
	// 	"filesystems": "filesystems",
	// }

	// regularTitles := map[string](map[string]string){
	// 	"users": {
	// 		"用户":     "user",
	// 		"用户名":    "user",
	// 		"uid":    "uid",
	// 		"userid": "uid",
	// 		"group":  "group",
	// 		"组":      "group",
	// 		"用户组":    "group",
	// 		"gid":    "gid",
	// 		"组id":    "gid",
	// 		"主目录":    "home",
	// 		"home":   "home",
	// 	},
	// 	"hosts": {
	// 		"主机名":  "hostname",
	// 		"主机":   "hostname",
	// 		"IP":   "ip",
	// 		"ip":   "ip",
	// 		"地址":   "ip",
	// 		"os":   "os",
	// 		"OS":   "os",
	// 		"操作系统": "os",
	// 	},
	// 	"filesystems": {
	// 		"路径":    "path",
	// 		"用户":    "user",
	// 		"逻辑卷":   "lv",
	// 		"大小":    "size",
	// 		"size":  "size",
	// 		"vg":    "vg",
	// 		"组":     "group",
	// 		"group": "group",
	// 	},
	// }

	regularName := dict.Key(name)
	if regularName == "" {
		fmt.Println("No dict mapping for sheetname " + name + ", use origin name.")
		regularName = name
	}

	return &SheetReader{regularName, dict}
}

// func (reader *SheetReader) GetTitles() map[string]int {
// 	titles := make(map[string]int)
// 	for _, key := range reader.ColumnTitleMaps {
// 		titles[key] = 1
// 	}
// 	return titles
// }

// func (reader *SheetReader) GetRegularTitle(title string) string {
// 	return reader.ColumnTitleMaps[strings.ToLower(title)]
// }

func (reader *SheetReader) ReadSheet(sheet *xlsx.Sheet) []map[string]string {
	dataList := make([]map[string]string, 0)

	titleCounts := make(map[string]int)
	i := 0
	for _, row := range sheet.Rows {
		if i == 0 {
			for index := 0; index < len(row.Cells); index++ {
				title := strings.TrimSpace(row.Cells[index].String())
				regularTitle := reader.dict.Key(title)
				if regularTitle == "" {
					regularTitle = title
				}

				titleCounts[regularTitle] = index + 1
				// fmt.Println(regularTitle)
				// if titleCounts[regularTitle] > 0 {

				// }

			}
		} else {
			datas := make(map[string]string)
			for j := 0; j < len(row.Cells); j++ {
				value := strings.TrimSpace(row.Cells[j].String())
				titleIndex := j + 1
				// title := ""
				for key, index := range titleCounts {
					if index == titleIndex {
						datas[key] = value
					}
				}

			}
			// for key, index := range titleCounts {
			// 	fmt.Println(len(row.Cells))
			// 	fmt.Println(index - 1)
			// 	if index-1 <  {

			// 	}
			// 	datas[key] = strings.TrimSpace(row.Cells[index-1].String())
			// }

			// If has no empty value, add the datas to list
			for _, value := range datas {
				if value != "" {
					dataList = append(dataList, datas)
					break
				}
			}

		}
		i++
	}
	fmt.Println(dataList)
	return dataList
}
