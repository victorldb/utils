package datastore

import (
	"sort"
)

/*
当前只支持列单一条件，>,>=,<,<=,==
*/

// Filter 现在只支持单一条件
type Filter struct {
	Key string
	Dms []Determine
}

// Determine --
type Determine struct {
	Value   interface{}
	ComType CompareType
}

func (c *DataStore) filter(rows []Row, filters []Filter) (res []Row, err error) {
	length := len(filters)
	if length == 0 {
		return rows, nil
	}

	// 获取当前要处理的列
	filterCols := make([]string, length)
	for k, v := range filters {
		filterCols[k] = v.Key
	}

	// 获取对应列定义
	currentCols := make([]*ColumnDef, length)
	if len(filterCols) > 0 {
		currentCols, err = c.getColInfoWithNames(filterCols)
		if err != nil {
			return nil, err
		}
	}

	// 计算交集
	setContainer := make(map[uint64]int)
	res = make([]Row, 0)

	rowsLength := len(rows)
	for i := 0; i < length; i++ {
		swapRows := make([]Row, rowsLength)
		copy(swapRows, rows)

		// 排序
		swapRows = c.sortRows(swapRows, ASC, currentCols[i:i+1])

		// startIndex, endIndex
		startIndex, endIndex := c.getFilterIndex(swapRows, filters[i], currentCols[i])
		if startIndex == endIndex {
			return
		}

		// 交集
		swapRows = swapRows[startIndex:endIndex]
		for _, v := range swapRows {
			count := 0
			id := v.GetID()
			if n, ok := setContainer[id]; ok {
				count = n + 1
				setContainer[id] = count
			} else {
				count = 1
				setContainer[id] = count
			}
			if count == length {
				res = append(res, v)
			}
		}
	}

	return res, nil
}

// rows必须是按增序排序 [i:n)
func (c *DataStore) getFilterIndex(rows []Row, filter Filter, currentCol *ColumnDef) (startIndex, endIndex int) {
	rowsLength := len(rows)
	for _, v := range filter.Dms {
		isNegate := false
		comType := v.ComType
		if comType == Less || comType == LessEqual {
			comType = negateCompareType(comType)
			isNegate = true
		} else if v.ComType == Equal {
			comType = GreaterEqual
		}
		findIndex := sort.Search(len(rows), func(i int) bool {
			comRes, _ := currentCol.Compare(rows[i].GetItem(currentCol.id), v.Value)
			return comType&comRes == comRes
		})
		if isNegate {
			startIndex = 0
			endIndex = findIndex
		} else {
			startIndex = findIndex
			endIndex = rowsLength
		}

		// 处理相等情况
		if v.ComType == Equal {
			startIndex = 0
			endIndex = 0
			if findIndex < rowsLength {
				for i := findIndex; i < rowsLength; i++ {
					comRes, _ := currentCol.Compare(rows[i].GetItem(currentCol.id), v.Value)
					if Equal != comRes {
						break
					}
					if i == findIndex {
						startIndex = i
					}
					endIndex = i + 1
				}
			}
		}
	}
	return startIndex, endIndex
}
