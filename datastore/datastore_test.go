package datastore

func getTestDataStore() (ds *DataStore, err error) {
	columns := []ColumnDef{
		ColumnDef{
			ColumnName: "id",
			ZeroValue:  int(0),
			Compare:    CompareIntFunc,
		},
		ColumnDef{
			ColumnName: "name",
			ZeroValue:  "",
			Compare:    CompareStringFunc,
		},
		ColumnDef{
			ColumnName: "phone",
			ZeroValue:  "",
			Compare:    CompareStringFunc,
		},
		ColumnDef{
			ColumnName:     "age",
			ZeroValue:      int(0),
			CalculateSlice: CalculateSumFloat64Slice,
			Compare:        CompareIntFunc,
		},
		ColumnDef{
			ColumnName:     "score",
			ZeroValue:      float64(0),
			CalculateSlice: CalculateSumFloat64Slice,
			Compare:        CompareFloat64Func,
		},
	}
	ds, err = NewDataStore(columns...)
	if err != nil {
		return nil, err
	}

	rows := []Row{
		NewDataRow([]interface{}{int(1), "name-1", "186-1", int(10), float64(60)}),
		NewDataRow([]interface{}{int(2), "name-3", "186-3", int(12), float64(68)}),
		NewDataRow([]interface{}{int(4), "name-2", "186-2", int(15), float64(61.5)}),
		NewDataRow([]interface{}{int(3), "name-2", "186-5", int(20), float64(99)}),
		NewDataRow([]interface{}{int(12), "name-2", "186-5", int(20), float64(99)}),
		NewDataRow([]interface{}{int(11), "name-2", "186-1", int(15), float64(65)}),
		NewDataRow([]interface{}{int(15), "name-6", "186-6", int(10), float64(63.5)}),
		NewDataRow([]interface{}{int(56), "name-5", "186-5", int(16), float64(65)}),
		NewDataRow([]interface{}{int(13), "name-1", "186-1", int(10), float64(60)}),
		NewDataRow([]interface{}{int(322), "name-3", "186-3", int(12), float64(68)}),
		NewDataRow([]interface{}{int(24), "name-2", "186-2", int(15), float64(61.5)}),
		NewDataRow([]interface{}{int(212), "name-2", "186-5", int(2), float64(99)}),
		NewDataRow([]interface{}{int(23), "name-2", "186-1", int(15), float64(65)}),
		NewDataRow([]interface{}{int(1), "name-1", "186-1", int(10), float64(60)}),
		NewDataRow([]interface{}{int(2), "name-3", "186-3", int(12), float64(68)}),
		NewDataRow([]interface{}{int(4), "name-2", "186-2", int(15), float64(61.5)}),
		NewDataRow([]interface{}{int(3), "name-2", "186-5", int(20), float64(99)}),
		NewDataRow([]interface{}{int(12), "name-2", "186-5", int(20), float64(99)}),
		NewDataRow([]interface{}{int(11), "name-2", "186-1", int(15), float64(65)}),
		NewDataRow([]interface{}{int(15), "name-6", "186-6", int(10), float64(63.5)}),
		NewDataRow([]interface{}{int(56), "name-5", "186-5", int(16), float64(65)}),
		NewDataRow([]interface{}{int(13), "name-1", "186-1", int(10), float64(60)}),
		NewDataRow([]interface{}{int(322), "name-3", "186-3", int(12), float64(68)}),
		NewDataRow([]interface{}{int(24), "name-2", "186-2", int(15), float64(61.5)}),
		NewDataRow([]interface{}{int(212), "name-2", "186-5", int(2), float64(99)}),
		NewDataRow([]interface{}{int(23), "name-2", "186-1", int(15), float64(65)}),
		NewDataRow([]interface{}{int(1), "name-1", "186-1", int(10), float64(60)}),
		NewDataRow([]interface{}{int(2), "name-3", "186-3", int(12), float64(68)}),
		NewDataRow([]interface{}{int(4), "name-2", "186-2", int(15), float64(61.5)}),
		NewDataRow([]interface{}{int(3), "name-2", "186-5", int(20), float64(99)}),
		NewDataRow([]interface{}{int(12), "name-2", "186-5", int(20), float64(99)}),
		NewDataRow([]interface{}{int(11), "name-2", "186-1", int(15), float64(65)}),
		NewDataRow([]interface{}{int(15), "name-6", "186-6", int(10), float64(63.5)}),
		NewDataRow([]interface{}{int(56), "name-5", "186-5", int(16), float64(65)}),
		NewDataRow([]interface{}{int(13), "name-1", "186-1", int(10), float64(60)}),
		NewDataRow([]interface{}{int(322), "name-3", "186-3", int(12), float64(68)}),
		NewDataRow([]interface{}{int(24), "name-2", "186-2", int(15), float64(61.5)}),
		NewDataRow([]interface{}{int(212), "name-2", "186-5", int(2), float64(99)}),
		NewDataRow([]interface{}{int(23), "name-2", "186-1", int(15), float64(65)}),
		NewDataRow([]interface{}{int(1), "name-1", "186-1", int(10), float64(60)}),
		NewDataRow([]interface{}{int(2), "name-3", "186-3", int(12), float64(68)}),
		NewDataRow([]interface{}{int(4), "name-2", "186-2", int(15), float64(61.5)}),
		NewDataRow([]interface{}{int(3), "name-2", "186-5", int(20), float64(99)}),
		NewDataRow([]interface{}{int(12), "name-2", "186-5", int(20), float64(99)}),
		NewDataRow([]interface{}{int(11), "name-2", "186-1", int(15), float64(65)}),
		NewDataRow([]interface{}{int(15), "name-6", "186-6", int(10), float64(63.5)}),
		NewDataRow([]interface{}{int(56), "name-5", "186-5", int(16), float64(65)}),
		NewDataRow([]interface{}{int(13), "name-1", "186-1", int(10), float64(60)}),
		NewDataRow([]interface{}{int(322), "name-3", "186-3", int(12), float64(68)}),
		NewDataRow([]interface{}{int(24), "name-2", "186-2", int(15), float64(61.5)}),
		NewDataRow([]interface{}{int(212), "name-2", "186-5", int(2), float64(99)}),
		NewDataRow([]interface{}{int(23), "name-2", "186-1", int(15), float64(65)}),
		NewDataRow([]interface{}{int(1), "name-1", "186-1", int(10), float64(60)}),
		NewDataRow([]interface{}{int(2), "name-3", "186-3", int(12), float64(68)}),
		NewDataRow([]interface{}{int(4), "name-2", "186-2", int(15), float64(61.5)}),
		NewDataRow([]interface{}{int(3), "name-2", "186-5", int(20), float64(99)}),
		NewDataRow([]interface{}{int(12), "name-2", "186-5", int(20), float64(99)}),
		NewDataRow([]interface{}{int(11), "name-2", "186-1", int(15), float64(65)}),
		NewDataRow([]interface{}{int(15), "name-6", "186-6", int(10), float64(63.5)}),
		NewDataRow([]interface{}{int(56), "name-5", "186-5", int(16), float64(65)}),
		NewDataRow([]interface{}{int(13), "name-1", "186-1", int(10), float64(60)}),
		NewDataRow([]interface{}{int(322), "name-3", "186-3", int(12), float64(68)}),
		NewDataRow([]interface{}{int(24), "name-2", "186-2", int(15), float64(61.5)}),
	}
	err = ds.InsertRows(rows)
	if err != nil {
		return nil, err
	}
	return ds, nil
}
