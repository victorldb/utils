package converttofile

import (
	"bytes"
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

// ConToXlsx ---
type ConToXlsx struct {
	// 表头
	Header []interface{}
	// 行
	Rows [][]interface{}
	// SheetName
	SheetName string
	// 表头样式
	HeadStyle string
}

// ConvertToXlsx --
func ConvertToXlsx(data ConToXlsx) (xlsxBuff *bytes.Buffer, err error) {
	xlsxFile := excelize.NewFile()
	firstSheetName := xlsxFile.GetSheetName(0)
	if firstSheetName == "" {
		xlsxFile.NewSheet(data.SheetName)
	} else if data.SheetName != firstSheetName {
		xlsxFile.SetSheetName(firstSheetName, data.SheetName)
	}
	row := 0
	colNum := 0
	if len(data.Header) > 0 {
		row++
		colNum = len(data.Header)
		xlsxFile.SetSheetRow(data.SheetName, fmt.Sprintf("A%d", row), &data.Header)
		if data.HeadStyle != "" {
			dataStyle, err := xlsxFile.NewStyle(data.HeadStyle)
			if err != nil {
				return nil, err
			}
			startCellName, _ := excelize.CoordinatesToCellName(1, 1)
			endCellName, _ := excelize.CoordinatesToCellName(len(data.Header), 1)
			xlsxFile.SetCellStyle(data.SheetName, startCellName, endCellName, dataStyle)
		}
	}

	for _, v := range data.Rows {
		row++
		if len(v) > colNum {
			colNum = len(v)
		}
		xlsxFile.SetSheetRow(data.SheetName, fmt.Sprintf("A%d", row), &v)
	}

	var maxWidth int
	colNumNames := getAllColNameWithLength(colNum)
	for _, v := range colNumNames {
		for i := 1; i <= row; i++ {
			if v, err := xlsxFile.GetCellValue(data.SheetName, getCellKey(v, i)); err == nil {
				if length := len([]rune(v)); length > maxWidth {
					maxWidth = length
				}
			}
		}
		xlsxFile.SetColWidth(data.SheetName, v, v, float64(maxWidth)*2.7)
		maxWidth = 0
	}
	dataCellStyle, err := xlsxFile.NewStyle(`{"alignment":{"horizontal":"left","vertical":"center","wrap_text":false}}`)
	if err != nil {
		return nil, err
	}
	scellName, _ := excelize.CoordinatesToCellName(1, 2)
	dcellName, _ := excelize.CoordinatesToCellName(len(data.Header), len(data.Rows)+1)
	xlsxFile.SetCellStyle(data.SheetName, scellName, dcellName, dataCellStyle)
	return xlsxFile.WriteToBuffer()
}

func getCellKey(prefix string, index int) string {
	return fmt.Sprintf("%s%d", prefix, index)
}

func getAllColNameWithLength(l int) (res []string) {
	if l == 0 {
		return []string{}
	}
	res = make([]string, l)
	for i := 1; i <= l; i++ {
		resByte := make([]byte, 0)
		idToColName(i, &resByte)
		res[i-1] = string(resByte)
	}
	return res
}

func getIDWithColName(s string) (res int, err error) {
	return colNameToID(s)
}

func getColNameWithID(id int) (res string) {
	if id == 0 {
		return ""
	}
	resByte := make([]byte, 0)
	idToColName(id, &resByte)
	return string(resByte)
}

func idToColName(id int, res *[]byte) {
	m := id / 26
	n := id % 26

	if m > 0 && n == 0 {
		m--
		n = 26
	}
	if m > 0 {
		idToColName(m, res)
	}
	if n > 0 {
		*res = append(*res, byte('A'+n-1))
	}
}

func colNameToID(s string) (n int, err error) {
	if len(s) == 0 {
		return 0, fmt.Errorf("syntaxError")
	}
	for _, c := range []byte(s) {
		var d byte
		switch {
		case 'a' <= c && c <= 'z':
			d = c - 'a' + 1
		case 'A' <= c && c <= 'Z':
			d = c - 'A' + 1
		default:
			return 0, fmt.Errorf("syntaxError:%s", s)
		}
		n *= 26
		n += int(d)
	}
	return n, nil
}

// func idToColNameStr(id int) (res string) {
// 	m := id / 26
// 	n := id % 26

// 	if m > 0 && n == 0 {
// 		m--
// 		n = 26
// 	}
// 	if m > 0 {
// 		res = idToColNameStr(m)
// 	}
// 	if n > 0 {
// 		res += string(byte('A' + n - 1))
// 	}
// 	return res
// }
