package main

import (
	"GenDataTable/GoData"
	"fmt"
	"reflect"
)

func main() {
	dtManager := GoData.MustInit()
	fmt.Printf("dtManager=%v\n", dtManager)
	dt := dtManager.GetDataTable("DTUserCommodity")
	//fmt.Printf("dt=%v\n", dt)
	uc := dt.GetDataRowById(91)
	fmt.Printf("uc=%v\n", uc)
	ucData := (*uc).(*GoData.DTUserCommodity)
	fmt.Printf("ucData=%v\n", ucData)
	fmt.Printf("Id:%d, Code:%s, Name:%s, Type:%d, Price:%d, Icon:%s\n", ucData.Id, ucData.Code, ucData.Name, ucData.Type, ucData.Price, ucData.Icon)

	ucType := reflect.TypeOf(uc)
	ucValue := reflect.ValueOf(uc)
	fmt.Printf("ucType=%v, ucValue=%v\n", ucType, ucValue)

	ucTypeElem := reflect.TypeOf(uc).Elem()
	ucValueElem := reflect.ValueOf(uc).Elem()
	fmt.Printf("ucTypeElem=%v, ucValueElem=%v\n", ucTypeElem, ucValueElem)

	ucValueElemField := reflect.ValueOf(ucData).Elem().Field(0)
	fmt.Printf("ucValueElemField=%v\n", ucValueElemField)

}
