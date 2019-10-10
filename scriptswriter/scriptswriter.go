package scriptswriter

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func NewScriptsWriter() ScriptsWriter {
	return ScriptsWriter{}
}

type ScriptsWriter struct {
}

func containModule(module string, mArr []string) bool {
	for _, m := range mArr {
		theM := strings.TrimSpace(m)
		if theM == module {
			return true
		}
	}
	return false

}

func (writer ScriptsWriter) WriteData(data []map[string](map[string]string), parentDir string, tmplDir string) {
	hosts := make(map[string](map[string]string), 0)
	for _, d := range data {
		hostname := d["host"]["hostname"]
		hosts[hostname] = d["host"]
	}

	pattern := filepath.Join(tmplDir, "*.tmpl")
	tmpl, err := template.ParseGlob(pattern)
	if err != nil {
		fmt.Println(err)
		return
	}

	// for _, host := range hosts {
	// 	host["content"] = "# Host: " + host["hostname"] + ", IP: " + host["ip"] + ", OS: " + host["os"]
	// }
	// var tpl bytes.Buffer
	// if err := t.Execute(&tpl, data); err != nil {
	// 	return err
	// }
	// result := tpl.String()

	for _, host := range hosts {

		os := strings.ToLower(host["os"])
		if os == "" {
			os = "linux"
		}

		fmt.Println("=== Generate host data with: 	" + host["hostname"])
		var hostTplBuffer bytes.Buffer
		hostTmplFile := "host." + os + ".tmpl"

		hostData := make(map[string](map[string]string))
		hostData["host"] = host

		fmt.Println("Prepare data with template: 	" + hostTmplFile)
		if err := tmpl.ExecuteTemplate(&hostTplBuffer, hostTmplFile, hostData); err != nil {
			fmt.Println("Error in process template: 	" + hostTmplFile)
			fmt.Println(err)
			return
		}
		hostResult := hostTplBuffer.String()
		host["content"] = hostResult

		for _, d := range data {
			if host["hostname"] == d["host"]["hostname"] {

				for key, _ := range d {
					if key != "host" {

						// module check, if set will filter

						if d[key]["module"] != "" {
							modulesStr := strings.Replace(host["module"], "，", ",", -1)
							modulesStr = strings.Replace(modulesStr, " ", " ", -1)
							moduleArr := strings.Split(modulesStr, ",")
							if !containModule(d[key]["module"], moduleArr) {
								continue
							}
						}

						var tpl bytes.Buffer

						// fmt.Println(os)
						tmplFile := key + "." + os + ".tmpl"
						fmt.Println("Prepare data with template: 	" + tmplFile)

						if err := tmpl.ExecuteTemplate(&tpl, tmplFile, d); err != nil {
							fmt.Println("Error in process template: 	" + tmplFile)
							fmt.Println(err)
							return
						}

						result := tpl.String()
						host["content"] = host["content"] + "\n" + result
						// fmt.Println(result)
					}
				}
			}
		}
	}

	scriptsDir := filepath.Join(parentDir, "scripts")

	if _, err := os.Stat(scriptsDir); err != nil {
		os.Mkdir(scriptsDir, os.ModePerm)
	}

	fmt.Println("=== Write hosts data:	")
	for _, host := range hosts {
		hostshpath := filepath.Join(scriptsDir, host["hostname"]+".sh")
		fmt.Println("Write host data to: 	" + hostshpath)
		file, err := os.Create(hostshpath)
		if err != nil {
			// handle the error here
			return
		}
		defer file.Close()
		file.WriteString(host["content"])

	}

}
