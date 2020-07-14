package datastore

import (
	"fmt"
)

// GroupResult --
type GroupResult struct {
	Keys   []string
	Values [][]interface{}
}

// Group --
func (c *DataStore) Group(calColumns []string, groupColumns []string, filters []Filter) (res *GroupResult, err error) {
	if len(calColumns) == 0 {
		return nil, fmt.Errorf("calColumns is empty")
	}

	rows, ok := c.GetAllRow()
	if !ok || len(rows) == 0 {
		return nil, fmt.Errorf("row length is zero")
	}

	// 先过滤
	rows, err = c.filter(rows, filters)
	if err != nil {
		return nil, err
	}

	// 聚合
	groupColsDef, err := c.getColInfoWithNames(groupColumns)
	if err != nil {
		return nil, err
	}
	rows = c.sortRows(rows, ASC, groupColsDef)

	// for k, v := range rows {
	// 	fmt.Println(k, v.(*DataRow).String())
	// }

	calColsDef, err := c.getColInfoWithNames(calColumns)
	if err != nil {
		return nil, err
	}

	rowsLength := len(rows)
	calColLength := len(calColumns)
	groupColLength := len(groupColumns)
	resValuesLength := groupColLength + calColLength

	resValues := make([][]interface{}, 0)
	swapValue := make([]interface{}, resValuesLength)
	swapCalValues := make([][]interface{}, calColLength)
	for i := 0; i < rowsLength; i++ {
		isNew := false
		for m := 0; m < groupColLength; m++ {
			if i == 0 || isNew {
				swapValue[m] = rows[i].GetItem(groupColsDef[m].id)
				continue
			}
			v := rows[i].GetItem(groupColsDef[m].id)
			comRes, _ := groupColsDef[m].Compare(v, swapValue[m])
			if NotEqual&comRes == comRes {
				// 计算结果
				for k := 0; k < calColLength; k++ {
					cv, _ := calColsDef[k].CalculateSlice(swapCalValues[k])
					swapValue[groupColLength+k] = cv
				}
				// 保存到结果集
				resValues = append(resValues, swapValue)

				// 初始化
				swapValue = make([]interface{}, resValuesLength)
				for n := 0; n < calColLength; n++ {
					swapCalValues[n] = swapCalValues[n][:0]
				}

				// 重新遍历
				m = -1
				isNew = true
			}
		}

		for m := 0; m < calColLength; m++ {
			swapCalValues[m] = append(swapCalValues[m], rows[i].GetItem(calColsDef[m].id))
		}
	}

	if rowsLength > 0 {
		// 计算结果
		for k := 0; k < calColLength; k++ {
			cv, _ := calColsDef[k].CalculateSlice(swapCalValues[k])
			swapValue[groupColLength+k] = cv
		}
		// 保存到结果集
		resValues = append(resValues, swapValue)
	}
	resKeys := make([]string, groupColLength+calColLength)
	copy(resKeys, groupColumns)
	copy(resKeys[groupColLength:], calColumns)
	res = &GroupResult{
		Keys:   resKeys,
		Values: resValues,
	}
	return res, nil
}
