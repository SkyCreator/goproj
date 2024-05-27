package main

import (
	"GenDataTable/AutoGen/GoGen"
	"GenDataTable/AutoGen/LuaGen"
	//"GenDataTable/Out/GoData"
	"fmt"
	"os"
)

const (
	DATA_TABLE_PATH = "./DataTable/"
)

func main() {
	/*测试读表方法
	{
		dtManager := GoData.MustInit()
		dt := dtManager.GetDTUserCommodityTable()
		fmt.Printf("Rows:%d\n", dt.Rows())
		ucData := dt.Get(91)
		fmt.Printf("Id:%d, Code:%s, Name:%s, Type:%d, Price:%d, Icon:%s\n", ucData.Id, ucData.Code, ucData.Name, ucData.Type, ucData.Price, ucData.Icon)
		ucDataAll := dt.GetAll()
		for i, data := range ucDataAll {
			fmt.Printf("i = %d, Id:%d, Code:%s, Name:%s, Type:%d, Price:%d, Icon:%s\n", i, data.Id, data.Code, data.Name, data.Type, data.Price, data.Icon)
		}
	}
	*/
	fmt.Println("os.Args = ", len(os.Args), "os.Args[0] = ", os.Args[0])
	if len(os.Args) <= 1 {
		GoGen.AutoGen(DATA_TABLE_PATH)
		return
	}
	switch os.Args[1] {
	case "lua":
		LuaGen.AutoGen(DATA_TABLE_PATH)
	case "go":
		GoGen.AutoGen(DATA_TABLE_PATH)
	case "all":
		LuaGen.AutoGen(DATA_TABLE_PATH)
		GoGen.AutoGen(DATA_TABLE_PATH)
	}
}
