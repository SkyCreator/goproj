package main

import (
	"GenDataTable/GoData"
	"fmt"
	//"reflect"
)

func main() {
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
