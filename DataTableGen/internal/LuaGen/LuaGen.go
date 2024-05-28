package LuaGen

import (
	"GenDataTable/internal"
	"fmt"
)

type GoGenDataTable = internal.GoGenDataTable

const (
	TEMPLATE_NAME = "./internal/LuaGen/lua.tpl"
	OUT_PATH      = "./Out/Lua/"
)

func AutoGen(tables []*GoGenDataTable) {
	fmt.Println("")
	fmt.Println("-----------Lua AutoGen start!-----------")
	for _, table := range tables {
		internal.GenByTemplate(OUT_PATH+table.TableName+".lua", TEMPLATE_NAME, table)
	}
	fmt.Println("-----------Lua AutoGen end!-----------")
	fmt.Println("")
}
