package GoData

func (d *DataTableManager) CreateAndParseDataTable(name string, rows [][]string) error {
	if name == "DTUserCommodity" {
		dt := d.GetDTUserCommodityTable()
		dt.ParseData(rows)
	}
	return nil
}
