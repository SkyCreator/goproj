package GoData

import (
	"GenDataTable/libFile"
	"fmt"
	"github.com/xuri/excelize/v2"
	"strings"
)

const (
	DATA_TABLE_PATH = "./DataTable/"
)

//行数据接口
type IDataRow interface {
	GetId() int //所有表第一列都是Id
	ParseData([]string)
}

//一行数据
type DataRow struct {
	Id int
}

func (r *DataRow) GetId() int {
	return r.Id
}
func (r *DataRow) ParseData(row []string) {

}

//数据表接口
type IDataTable interface {
	ParseData([][]string) error
	GetDataRowById(int) *IDataRow
	GetDataRowByIndex(int) *IDataRow
	GetLength() int
}

//数据表
type DataTable struct {
	name string
	drs  []IDataRow
}

func NewDataTable(name string) *DataTable {
	return &DataTable{
		name: name,
		drs:  make([]IDataRow, 0),
	}
}

func (r *DataTable) ParseData(data [][]string) error {
	l := len(data)
	for i := 3; i < l; i++ {
		row := data[i][0]
		if strings.IndexByte(row, '#') == 0 {
			continue
		}
		drs, err := CreateDataRow(r.name)
		if err != nil {
			fmt.Println(err)
			return err
		}
		r.drs = append(r.drs, drs)
		drs.ParseData(data[i])
	}
	return nil
}

type DataTableManager struct {
	DataTableMap map[string]IDataTable
}

func newDataTableManager() *DataTableManager {
	return &DataTableManager{DataTableMap: make(map[string]IDataTable)}
}

func (d *DataTableManager) ReadAllDataTable() error {
	files, err := libFile.GetFilesFromDir(DATA_TABLE_PATH)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		d.ReadData(DATA_TABLE_PATH, file.Name())
	}
	return nil
}
func (d *DataTableManager) ReadData(path string, allName string) error {
	f, err := excelize.OpenFile(path + allName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return err
	}
	// 确保 rows 至少有三行数据
	if len(rows) < 3 {
		fmt.Println("Sheet1 要包含至少三行数据")
		return err
	}
	dtName, found := strings.CutSuffix(allName, ".xlsx")
	if !found {
		fmt.Println(allName + "不是以xlsx结尾的excel文件!")
		return err
	}
	dtName = "DT" + dtName
	d.DataTableMap[dtName] = NewDataTable(dtName)
	err = d.DataTableMap[dtName].ParseData(rows)
	return err
}
