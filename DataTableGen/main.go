package main

import (
	"GenDataTable/GoData"
	"fmt"
)

func main() {
	dtManager := GoData.MustInit()
	fmt.Printf("dtManager=%v\n", dtManager)
	dt := dtManager.GetDataTable("DTUserCommodity")
	uc := dt.GetDataRowById(91)
	ucData := (*uc).(*GoData.DTUserCommodity)
	fmt.Printf("Id:%d, Code:%s, Name:%s, Type:%d, Price:%d, Icon:%s\n", ucData.Id(), ucData.Code(), ucData.Name(), ucData.Type(), ucData.Price(), ucData.Icon())
}
