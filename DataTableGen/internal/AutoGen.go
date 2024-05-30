package internal

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"strings"
	"text/template"
)

type GoGenColData struct {
	ColType  string   //表列类型名
	ColName  string   //表列名
	ColValue []string //表列的行数据
}

func (cd *GoGenColData) SetColInfo(colType string, colName string) {
	cd.ColType = colType
	cd.ColName = colName
	cd.ColValue = make([]string, 0)
}
func (cd *GoGenColData) AddColValue(colValue string) {
	cd.ColValue = append(cd.ColValue, colValue)
}

type GoGenCellData struct {
	CellType  string
	CellName  string
	CellValue string
}
type GoGenRowData struct {
	CellId   string
	CellData []GoGenCellData
}

func (rd *GoGenRowData) AddRowInfo(cellType string, cellName string, CellValue string) {
	if cellType == "BOOL" {
		CellValue = strings.ToLower(CellValue)
	}
	rd.CellData = append(rd.CellData, GoGenCellData{CellType: cellType, CellName: cellName, CellValue: CellValue})
}

type GoGenDataTable struct {
	TableName      string         //表名
	TableColData   []GoGenColData //表列数据
	TableRowData   []GoGenRowData //表行数据							for lua
	SpecialRowData GoGenRowData   //特殊行数据,各个列数据出现最多的入选   for lua
}

func (dt *GoGenDataTable) SetColInfo(colTypes []string, colNames []string) {
	for i := 0; i < len(colTypes); i++ {
		dt.TableColData[i].SetColInfo(colTypes[i], colNames[i])
	}
}
func (dt *GoGenDataTable) AddColValue(colIdx int, colValue string) {
	dt.TableColData[colIdx].AddColValue(colValue)
}
func (dt *GoGenDataTable) SetRowInfo(rowsValue [][]string) {
	typeCols := rowsValue[0]
	nameCols := rowsValue[1]
	dataRows := rowsValue[2:]
	dt.TableRowData = make([]GoGenRowData, len(dataRows))
	for i := 0; i < len(dataRows); i++ {
		row := dataRows[i]
		dt.TableRowData[i].CellId = row[0]
		for j := 0; j < len(row); j++ {
			if j > 0 && row[j] == dt.SpecialRowData.CellData[j-1].CellValue {
				continue
			}
			dt.TableRowData[i].AddRowInfo(typeCols[j], nameCols[j], row[j])

		}
	}
}
func (dt *GoGenDataTable) SetSpecialRowData() {
	for i := 1; i < len(dt.TableColData); i++ {
		colData := dt.TableColData[i].ColValue
		specialValue := make(map[string]int, len(colData))
		for j := 0; j < len(colData); j++ {
			specialValue[colData[j]] = specialValue[colData[j]] + 1
		}
		maxCount := 0
		maxValue := ""
		for k, v := range specialValue {
			if v > maxCount {
				maxCount = v
				maxValue = k
			}
		}
		dt.SpecialRowData.AddRowInfo(dt.TableColData[i].ColType, dt.TableColData[i].ColName, maxValue)
	}
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
	//fmt.Println("AutoGen FilterData")
	colsNum := len(rows[0]) - 1
	filterRows := make([][]string, 0)
	//fmt.Println("AutoGen FilterData rows length is ", len(rows))
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
	//dt.SetRowInfo(filterRows)
	// 获取行数
	rowsNum := len(dataRows)
	for i := 0; i < rowsNum; i++ {
		dataRow := dataRows[i]
		for j := 0; j < len(dataRow); j++ {
			dt.AddColValue(j, dataRow[j])
		}
	}
	dt.SetSpecialRowData()
	dt.SetRowInfo(filterRows)
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

func GenByTemplate(outAllPath string, tmplPath string, data any) {
	outPathPos := strings.LastIndexByte(outAllPath, '/')
	outPath := outAllPath[:outPathPos]
	os.MkdirAll(outPath, os.ModeDir|os.ModePerm)
	pos := strings.LastIndexByte(tmplPath, '/')
	tmplName := tmplPath[pos+1:]
	tmpl := template.New(tmplName).Funcs(template.FuncMap{
		"add": func(a, b int) int { return a + b }})
	tmpl, err := tmpl.ParseFiles(tmplPath)
	CheckErr(err)
	os.Remove(outAllPath)
	f, err := os.OpenFile(outAllPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0755)
	CheckErr(err)
	defer f.Close()
	err = tmpl.Execute(f, data)
	CheckErr(err)
	fmt.Printf("gen file：%s\n", outAllPath)
}
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
