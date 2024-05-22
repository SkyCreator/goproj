package GoData

import (
	"strconv"
)

type DTUserCommodity struct {
	DataRow
	Code  string
	Name  string
	Type  int
	Price int
	Icon  string
}

func (dt *DTUserCommodity) ParseData(data []string) {
	length := len(data)
	var idx int = 0
	if length > 0 {
		dt.Id, _ = strconv.Atoi(data[idx])
		idx++
	}
	idx++
	if length > 2 {
		dt.Code = data[idx]
		idx++
	}
	if length > 3 {
		dt.Name = data[idx]
		idx++
	}
	if length > 4 {
		dt.Type, _ = strconv.Atoi(data[idx])
		idx++
	}
	if length > 5 {
		dt.Price, _ = strconv.Atoi(data[idx])
		idx++
	}
	if length > 6 {
		dt.Icon = data[idx]
		idx++
	}
}
func (dt *DTUserCommodity) GetId() int {
	return dt.Id
}
