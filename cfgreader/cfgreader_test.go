package cfgreader

import (
	"fmt"
	"testing"
)

func TestCFGReader_ReadFile(t *testing.T) {
	type args struct {
		excelFile string
		dict      KeyDict
	}

	reader := NewCFGReader()

	dict, _ := NewSimpleFileKeyDict("test/testdict.json")
	fmt.Println(dict)

	tests := []struct {
		name      string
		cfgreader *CFGReader
		args      args
	}{
		{
			"Test Read Excel File",
			&reader,
			args{
				"test/core-tuxedo.xlsx",
				dict,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cfgreader.ReadFile(tt.args.excelFile, tt.args.dict)
		})
	}
}
