package GoData

import (
	"strconv"
	"strings"
)

type DTUserCommodity struct {
	Id    int
	Code  string
	Name  string
	Type  int
	Price int
	Icon  string
}

func NewDTUserCommodity() *DTUserCommodity {
	return &DTUserCommodity{}
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

type DTUserCommodityTable struct {
	drs []*DTUserCommodity
}

func (dt *DTUserCommodityTable) Rows() int {
	return len(dt.drs)
}

func (dt *DTUserCommodityTable) Get(id int) *DTUserCommodity {
	for i := 0; i < len(dt.drs); i++ {
		if dt.drs[i] != nil && dt.drs[i].Id == id {
			return dt.drs[i]
		}
	}
	return nil
}
func (dt *DTUserCommodityTable) GetAll() []*DTUserCommodity {
	return dt.drs
}

func (dt *DTUserCommodityTable) parseData(data [][]string) error {
	l := len(data)
	var index = 0
	for i := 3; i < l; i++ {
		row := data[i][0]
		if strings.IndexByte(row, '#') == 0 {
			continue
		}
		drs := NewDTUserCommodity()
		drs.ParseData(data[i])
		dt.drs = append(dt.drs, drs)
		index++
	}
	return nil
}
func (dtm *DataTableManager) GetDTUserCommodityTable() *DTUserCommodityTable {
	if dt, ok := dtm.DataTableMap["DTUserCommodity"]; ok {
		return dt.(*DTUserCommodityTable)
	}
	dt := &DTUserCommodityTable{}
	dtm.DataTableMap["DTUserCommodity"] = dt
	return dt
}
