package cfgreader

import (
	"testing"
)

func TestSimpleFileKeyDict_Key(t *testing.T) {
	type args struct {
		value string
	}

	dict, _ := NewSimpleFileKeyDict("test/testdict.json")
	tests := []struct {
		name string
		dict *SimpleFileKeyDict
		args args
		want string
	}{
		{
			"simple test",
			dict,
			args{"主机名"},
			"hostname",
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dict.Key(tt.args.value); got != tt.want {
				t.Errorf("SimpleFileKeyDict.Key() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSimpleFileKeyDict(t *testing.T) {
	type args struct {
		filename string
	}

	nonDict, err := NewSimpleFileKeyDict("test/testdd.json")

	tests := []struct {
		name    string
		args    args
		want    *SimpleFileKeyDict
		wantErr bool
	}{
		{
			"Nonexist Dict File test",
			args{"key"},
			nonDict,
			true,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSimpleFileKeyDict() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("NewSimpleFileKeyDict() = %v, want %v", got, tt.want)
			// }
		})
	}
}
