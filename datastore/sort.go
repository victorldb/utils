package datastore

import (
	"sort"
)

// SortType --
type SortType uint8

const (
	// ASC --
	ASC = 1

	// DESC --
	DESC = 2
)

func (c *DataStore) sortRows(rows []Row, sortType SortType, currentCols []*ColumnDef) (res []Row) {
	var checkCompare, reverseCheckCompare CompareType
	if sortType == ASC {
		checkCompare = Less
		reverseCheckCompare = Greater
	} else {
		checkCompare = Greater
		reverseCheckCompare = Less
	}

	length := len(currentCols)
	sort.Slice(rows, func(i, j int) bool {
		var v1, v2 interface{}
		var coldef *ColumnDef
		var compareRes CompareType

		for n := 0; n < length; n++ {
			v1 = rows[i].GetItem(currentCols[n].id)
			v2 = rows[j].GetItem(currentCols[n].id)

			coldef = currentCols[n]
			compareRes, _ = coldef.Compare(v1, v2)
			if n == length-1 {
				break
			}

			if compareRes == checkCompare {
				return true
			}
			if compareRes == reverseCheckCompare {
				return false
			}
		}

		return compareRes == checkCompare
	})
	return rows
}
