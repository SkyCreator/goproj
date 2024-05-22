package GoData

import (
	"strconv"
)

type DTUserCommodity struct {
	DataRow
	m0 int
	m1 string
	m2 string
	m3 int
	m4 int
	m5 string
}

func (dt *DTUserCommodity) Id() int {
	return dt.m0
}
func (dt *DTUserCommodity) Code() string {
	return dt.m1
}
func (dt *DTUserCommodity) Name() string {
	return dt.m2
}
func (dt *DTUserCommodity) Type() int {
	return dt.m3
}
func (dt *DTUserCommodity) Price() int {
	return dt.m4
}
func (dt *DTUserCommodity) Icon() string {
	return dt.m5
}

func (dt *DTUserCommodity) ParseData(data []string) {
	length := len(data)
	var idx int = 0
	if length > 0 {
		dt.m0, _ = strconv.Atoi(data[idx])
		idx++
	}
	idx++
	if length > 2 {
		dt.m1 = data[idx]
		idx++
	}
	if length > 3 {
		dt.m2 = data[idx]
		idx++
	}
	if length > 4 {
		dt.m3, _ = strconv.Atoi(data[idx])
		idx++
	}
	if length > 5 {
		dt.m4, _ = strconv.Atoi(data[idx])
		idx++
	}
	if length > 6 {
		dt.m5 = data[idx]
		idx++
	}
}
func (dt *DTUserCommodity) GetId() int {
	return dt.m0
}
func NewDTUserCommodity(drId int, m0 int, m1 string, m2 string, m3 int, m4 int, m5 string) *DTUserCommodity {
	return &DTUserCommodity{
		DataRow{drId}, m0, m1, m2, m3, m4, m5,
	}
}
