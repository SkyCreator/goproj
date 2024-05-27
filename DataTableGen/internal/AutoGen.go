package internal

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"strings"
)

type GoGenColData struct {
	ColType  string   //表列类型名
	ColName  string   //表列名
	ColValue []string //表列的行数据值
}

func (cd *GoGenColData) SetColInfo(colType string, colName string) {
	cd.ColType = colType
	cd.ColName = colName
	cd.ColValue = make([]string, 0)
}
func (cd *GoGenColData) AddColValue(colValue string) {
	cd.ColValue = append(cd.ColValue, colValue)
}

type GoGenDataTable struct {
	TableName    string         //表名
	TableColData []GoGenColData //表列数据
}

func (dt *GoGenDataTable) SetColInfo(colTypes []string, colNames []string) {
	for i := 0; i < len(colTypes); i++ {
		dt.TableColData[i].SetColInfo(colTypes[i], colNames[i])
	}
}
func (dt *GoGenDataTable) AddColValue(colIdx int, colValue string) {
	dt.TableColData[colIdx].AddColValue(colValue)
}
func NewGoGenDataTable(tableName string, colsNum int) *GoGenDataTable {
	return &GoGenDataTable{TableName: tableName, TableColData: make([]GoGenColData, colsNum)}
}

func ReadAllDataTable(tableDir string) ([]*GoGenDataTable, error) {
	files, err := os.ReadDir(tableDir)
	if err != nil {
		fmt.Printf("ReadAllDataTable failed! err: %v\n", err)
		return nil, err
	}
	allDataTable := make([]*GoGenDataTable, 0)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		tableName, rows, err := ReadData(tableDir, file.Name())
		if err != nil {
			fmt.Printf("ReadData failed! err: %v\n", err)
			return nil, err
		}
		dt := ParseData(tableName, rows)
		allDataTable = append(allDataTable, dt)
		fmt.Printf("dt.TableName:%s\n", dt.TableName)
		for i := 0; i < len(dt.TableColData); i++ {
			data := dt.TableColData[i]
			fmt.Printf("i : %d, ColName : %s, ColType : %s, ColValue : = %v\n", i, data.ColName, data.ColType, data.ColValue)
		}
	}
	return allDataTable, nil
}

func ReadData(tableDir string, allName string) (tableName string, rows [][]string, err error) {
	f, err := excelize.OpenFile(tableDir + allName)
	if err != nil {
		fmt.Println(err)
		return "", nil, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// 获取 Sheet1 上所有单元格
	rows, err = f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return "", nil, err
	}
	// 确保 rows 至少有三行数据
	if len(rows) < 3 {
		fmt.Println("Sheet1 要包含至少三行数据")
		return "", nil, err
	}
	dtName, found := strings.CutSuffix(allName, ".xlsx")
	if !found {
		fmt.Println(allName + "不是以xlsx结尾的excel文件!")
		return "", nil, err
	}

	return dtName, rows, err
}

//过滤数据，去除第二列，去除第二行，去除掉#开头的行
func FilterData(rows [][]string) [][]string {
	colsNum := len(rows[0]) - 1
	filterRows := make([][]string, 0)
	for i := 0; i < len(rows); i++ {
		// 过滤掉第二行
		if i == 1 {
			continue
		}
		row := rows[i]
		// 过滤掉#开头的行
		rowID := row[0]
		if strings.IndexByte(rowID, '#') == 0 {
			continue
		}
		// 过滤掉第二列
		row = append(row[:1], row[2:]...)
		// 补空
		if colsNum > len(row) {
			row = append(row, "")
		}
		filterRows = append(filterRows, row)
	}

	return filterRows
}
func ParseData(tableName string, rows [][]string) *GoGenDataTable {
	filterRows := FilterData(rows)
	// 获取列类型
	colTypes := filterRows[0]
	// 获取列名
	colNames := filterRows[1]
	// 获取列数
	colsNum := len(colTypes)
	// 创建数据结构
	dt := NewGoGenDataTable(tableName, colsNum)
	dt.SetColInfo(colTypes, colNames)
	dataRows := filterRows[2:]
	// 获取行数
	rowsNum := len(dataRows)
	for i := 0; i < rowsNum; i++ {
		dataRow := dataRows[i]
		for j := 0; j < len(dataRow); j++ {
			dt.AddColValue(j, dataRow[j])
		}
	}
	return dt
}

func ForTemplate(dt *GoGenDataTable) *GoGenDataTable {
	for _, colData := range dt.TableColData {
		if colData.ColType != "STRING" {
			break
		}
	}
	return dt
}
