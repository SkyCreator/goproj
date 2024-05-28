package main

import (
	"GenDataTable/internal"
	"GenDataTable/internal/GoGen"
	"GenDataTable/internal/LuaGen"
	//"GenDataTable/GoData"
	"fmt"
	"os"
)

const (
	DATA_TABLE_PATH = "./DataTable/"
)

func main() {
	/*
		{
			//测试读表方法
			dtManager := GoData.MustInit()
			dt := dtManager.GetDTUserCommodityTable()
			fmt.Printf("Rows:%d\n", dt.Rows())
			ucData := dt.Get(91)
			fmt.Printf("Id:%d, Code:%s, Name:%s, Type:%d, Price:%d, Icon:%s\n", ucData.Id, ucData.Code, ucData.Name, ucData.Type, ucData.Price, ucData.Icon)
			ucDataAll := dt.GetAll()
			for i, data := range ucDataAll {
				fmt.Printf("i = %d, Id:%d, Code:%s, Name:%s, Type:%d, Price:%d, Icon:%s\n", i, data.Id, data.Code, data.Name, data.Type, data.Price, data.Icon)
			}
		}*/

	fmt.Println("os.Args = ", len(os.Args), "os.Args[0] = ", os.Args[0])
	allDataTable, err := internal.ReadAllDataTable(DATA_TABLE_PATH)
	if err != nil {
		fmt.Println("AutoGen err = ", err)
		return
	}
	if len(os.Args) <= 1 {
		LuaGen.AutoGen(allDataTable)
		return
	}
	switch os.Args[1] {
	case "lua":
		LuaGen.AutoGen(allDataTable)
	case "go":
		GoGen.AutoGen(allDataTable)
	case "all":
		LuaGen.AutoGen(allDataTable)
		GoGen.AutoGen(allDataTable)
	}
}
