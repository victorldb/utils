package config

import (
	"encoding/json"

	"github.com/victorldb/utils"
)

// ParseJSONConfig --
func ParseJSONConfig(s string, v interface{}) (err error) {
	res, err := utils.RemoveTextRemark(s)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(res), v)
}
