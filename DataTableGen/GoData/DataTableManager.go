package GoData

import "fmt"

func (d *DataTableManager) RegisterAllDataTable() {
	d.RegisterDataTable("DTUserCommodity", NewDataTable("DTUserCommodity"))
}
func CreateDataRow(name string) (IDataRow, error) {
	switch name {
	case "DTUserCommodity":
		return &DTUserCommodity{}, nil
	default:
		err := fmt.Errorf("CreateDataRow: %s not found", name)
		return nil, err
	}
}
