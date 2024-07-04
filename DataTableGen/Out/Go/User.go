// Code generated by DataTableGen. DO NOT EDIT.

package GoData

import "strconv"

type DTUser struct {
    Id int
    Name string
    Height float64
    Sex bool
    AGE int
}

func (dtm *DataTableManager) GetDTUserTable() *DTUserTable {
	if dt, ok := dtm.DataTableMap["DTUser"]; ok {
		return dt.(*DTUserTable)
	}
	dt := &DTUserTable{}
	dtm.DataTableMap["DTUser"] = dt
	return dt
}

func (dt *DTUserTable) Rows() int {
	return len(dt.drs)
}

func (dt *DTUserTable) Get(index int) *DTUser {
	if index < 0 || index >= len(dt.drs) {
		return nil
	}
	return dt.drs[index]
}

func (dt *DTUserTable) GetById(id int) *DTUser {
	for i := 0; i < len(dt.drs); i++ {
		if dt.drs[i] != nil && dt.drs[i].Id == id {
			return dt.drs[i]
		}
	}
	return nil
}

func (dt *DTUserTable) GetAll() []*DTUser {
	return dt.drs
}

func newDTUser() *DTUser {
	return &DTUser{}
}

func (dt *DTUser) parseData(data []string) {
	dt.Id, _ = strconv.Atoi(data[0])
    dt.Name = data[1]
    dt.Height, _ = strconv.ParseFloat(data[2], 64)
    dt.Sex, _ = strconv.ParseBool(data[3])
	dt.AGE, _ = strconv.Atoi(data[4])
}

type DTUserTable struct {
	drs []*DTUser
}

func (dt *DTUserTable) parseData(data [][]string) error {
	l := len(data)
	var index = 0
	for i := 0; i < l; i++ {
		drs := newDTUser()
		drs.parseData(data[i])
		dt.drs = append(dt.drs, drs)
		index++
	}
	return nil
}
