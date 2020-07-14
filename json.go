package utils

import (
	"bytes"
	"encoding/json"
)

// FormatJSON 格式化JSON数据
func FormatJSON(data []byte, seq string) (res []byte, err error) {
	if seq == "" {
		seq = "\t"
	}
	var out bytes.Buffer
	err = json.Indent(&out, data, "", seq)
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}
