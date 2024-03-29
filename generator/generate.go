package main

import (
	"bytes"
	"os"
	"strings"
	"text/template"
)

type data struct {
	Type string
}

func main() {
	funcMap := template.FuncMap{
		"ToTitle": strings.Title,
	}
	types := []string{
		"bool",
		"int",
		"int8",
		"int16",
		"int32",
		"int64",
		"uint8",
		"uint16",
		"uint32",
		"uint64",
		"float32",
		"float64",
		"string",
	}
	header, err := os.ReadFile("header")
	if err != nil {
		panic(err)
	}
	{
		content, err := os.ReadFile("type.tmpl")
		if err != nil {
			panic(err)
		}
		tmpl, _ := template.New("").Funcs(funcMap).Parse(string(content))
		var buffer bytes.Buffer
		buffer.Write(header)
		buffer.WriteByte('\n')
		for _, t := range types {
			println("generate", t)
			templateData := data{
				Type: t,
			}
			var result bytes.Buffer

			tmpl.Execute(&result, templateData)
			buffer.Write(result.Bytes())
			buffer.WriteByte('\n')
		}
		os.WriteFile("../types.go", buffer.Bytes(), 0644)

	}
	{
		content, err := os.ReadFile("slice.tmpl")
		if err != nil {
			panic(err)
		}
		tmpl, _ := template.New("").Funcs(funcMap).Parse(string(content))
		var buffer bytes.Buffer
		buffer.Write(header)
		buffer.WriteByte('\n')
		for _, t := range types {
			println("generate", t)
			templateData := data{
				Type: t,
			}
			var result bytes.Buffer

			tmpl.Execute(&result, templateData)
			buffer.Write(result.Bytes())
			buffer.WriteByte('\n')
		}
		os.WriteFile("../slices.go", buffer.Bytes(), 0644)
	}

}
