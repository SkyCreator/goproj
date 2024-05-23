package main

import (
	"GenDataTable/GoData"
	"fmt"
	//"reflect"
)

func main() {
	dtManager := GoData.MustInit()
	dt := dtManager.GetDTUserCommodityTable()
	ucData := dt.Get(91)
	fmt.Printf("Id:%d, Code:%s, Name:%s, Type:%d, Price:%d, Icon:%s\n", ucData.Id, ucData.Code, ucData.Name, ucData.Type, ucData.Price, ucData.Icon)
}
