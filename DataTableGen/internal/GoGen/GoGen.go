package GoGen

import (
	"GenDataTable/internal"
	"fmt"
	"os"
	"strings"
	"text/template"
)

type GoGenDataTable = internal.GoGenDataTable

const (
	TEMPLATE_NAME         = "./internal/GoGen/gogen.tpl"
	FACTORY_TEMPLATE_NAME = "./internal/GoGen/gofactorygen.tpl"
	OUT_PATH              = "./Out/Go/"
	FACTORY_NAME          = "DataTableFactory"
)

func AutoGen(tables []*GoGenDataTable) {
	fmt.Println("-----------GoGen AutoGen!-----------")
	for _, table := range tables {
		GenByTemplate(OUT_PATH+table.TableName+".go", TEMPLATE_NAME, table)
	}
	GenByTemplate(OUT_PATH+FACTORY_NAME+".go", FACTORY_TEMPLATE_NAME, tables)
}
func GenByTemplate(outPath string, tmplPath string, data any) {
	pos := strings.LastIndexByte(tmplPath, '/')
	tmplName := tmplPath[pos+1:]
	tmpl := template.New(tmplName)
	tmpl, err := tmpl.ParseFiles(tmplPath)
	CheckErr(err)
	os.Remove(outPath)
	f, err := os.OpenFile(outPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0755)
	CheckErr(err)
	defer f.Close()
	err = tmpl.Execute(f, data)
	CheckErr(err)
	fmt.Printf("gen file：%s\n", outPath)
}
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
