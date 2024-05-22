package GoData

func (r *DataTable) GetDataRowByIndex(index int) *IDataRow {
	return &r.drs[index]
}

func (r *DataTable) GetDataRowById(id int) *IDataRow {
	for i := 0; i < len(r.drs); i++ {
		if r.drs[i].GetId() == id {
			return &r.drs[i]
		}
	}
	return nil
}
func (r *DataTable) GetLength() int {
	return len(r.drs)
}

func MustInit() *DataTableManager {
	manager := newDataTableManager()
	manager.ReadAllDataTable()
	return manager
}

func (d *DataTableManager) GetDataTable(tableName string) IDataTable {
	return d.DataTableMap[tableName]
}
