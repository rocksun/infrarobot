package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/rocksun/infra-robot/cfgreader"
	"github.com/rocksun/infra-robot/scriptswriter"
)

func printUsage(args []string) {
	fmt.Println("Usage:")
	fmt.Println(args[0] + " [DataDirectory] [TpmlDir](optional)")

}

func scanDirectory(dir string, tmplDir string, dictFile string) {
	dict, e := cfgreader.NewSimpleFileKeyDict(dictFile)
	if e != nil {
		fmt.Println(e)
		return
	}

	reader := cfgreader.NewCFGReader(dict)

	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}

		if strings.HasSuffix(filepath.Base(path), "xlsx") &&
			!strings.HasPrefix(filepath.Base(path), "~$") {
			fmt.Println("=== Scan file:	", path)
			reader.ReadFile(path)
		}

		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	data := reader.GetData()
	// fmt.Println(data)
	writer := scriptswriter.NewScriptsWriter()

	writer.WriteData(data, dir, tmplDir)

}

func main() {
	args := os.Args
	// fmt.Println(len(args))
	if len(args) != 3 && len(args) != 2 {
		printUsage(args)
		return
	}

	dir := args[1]

	rootPath := executablePath()
	dictConfigFile := filepath.Join(rootPath, "config/dict.json")
	fmt.Println("*** Use Dict File: 		" + dictConfigFile)

	tmpl := ""
	if len(args) == 2 {
		tmpl = filepath.Join(rootPath, "config/tmpl")
	} else {
		tmpl = args[2]
	}
	fmt.Println("*** Use Template At: 		" + tmpl)
	fmt.Println("*** Scan Target Directory: " + dir)

	scanDirectory(dir, tmpl, dictConfigFile)
}

func executablePath() string {
	ex, err0 := os.Executable()
	if err0 != nil {
		fmt.Println(err0)
		return ""
	}
	dir, err1 := filepath.Abs(filepath.Dir(ex))
	if err1 != nil {
		fmt.Println(err1)
		return ""
	} else {
		return dir
	}

}
