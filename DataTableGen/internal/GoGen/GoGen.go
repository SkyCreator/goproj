package GoGen

import (
	"GenDataTable/internal"
	"fmt"
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
		internal.GenByTemplate(OUT_PATH+table.TableName+".go", TEMPLATE_NAME, table)
	}
	internal.GenByTemplate(OUT_PATH+FACTORY_NAME+".go", FACTORY_TEMPLATE_NAME, tables)
}

