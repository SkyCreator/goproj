// Code generated by DataTableGen. DO NOT EDIT.

package GoData

func (d *DataTableManager) createAndParseDataTable(name string, rows [][]string) error {
    switch name {
    {{- range .}}   
	case "DT{{.TableName}}":
		d.GetDT{{.TableName}}Table().parseData(rows)
		return nil
    {{- end}}
     }
	return nil
}
