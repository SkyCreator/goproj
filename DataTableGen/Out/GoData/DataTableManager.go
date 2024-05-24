package GoData

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"strings"
)

const (
	DATA_TABLE_PATH = "./DataTable/"
)

type DataTableManager struct {
	DataTableMap map[string]any
}

func MustInit() *DataTableManager {
	manager := newDataTableManager()
	manager.ReadAllDataTable()
	return manager
}
func newDataTableManager() *DataTableManager {
	return &DataTableManager{DataTableMap: make(map[string]any)}
}
func (d *DataTableManager) ReadAllDataTable() error {
	files, err := os.ReadDir(DATA_TABLE_PATH)
	if err != nil {
		fmt.Printf("ReadAllDataTable failed! err: %v\n", err)
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
	err = d.CreateAndParseDataTable(dtName, rows)
	return err
}
