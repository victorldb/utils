package utils

import (
	"fmt"
	"strings"
)

// RemoveTextRemark --
// 移除段落注释，[/*]和[*/]必须是成对出现，段注释中可以嵌套段注释
// 移除行注释，[//][#]前面可以是任意空格和TAB
// 返回的结果会删除空行
func RemoveTextRemark(s string) (res string, err error) {
	// 先移除段落注释
	res, err = rmSectionRemark(s)
	if err != nil {
		return "", err
	}

	if strings.Index(res, "/*") != -1 || strings.Index(res, "*/") != -1 {
		return "", fmt.Errorf("section remark format is error")
	}

	// 移除单行注释
	res = rmRowRemark(res)
	return res, nil
}

func rmSectionRemark(s string) (res string, err error) {
	ind1 := strings.Index(s, "*/")
	if ind1 != -1 {
		ind := strings.LastIndex(s[:ind1], "/*")
		if ind == -1 || ind > ind1 {
			return "", fmt.Errorf("section remark format is error")
		}
		s = s[:ind] + s[ind1+2:]
		if strings.Index(s, "/*") != -1 {
			s, err = rmSectionRemark(s)
			if err != nil {
				return "", err
			}
		}
	}
	return s, nil
}

func rmRowRemark(s string) (res string) {
	contents := make([]string, 0)
	rp1 := strings.NewReplacer(" ", "", "	", "")
	sp1 := strings.Split(s, "\n")
	for _, row := range sp1 {
		if strings.Index(row, "//") != -1 {
			if strings.HasPrefix(rp1.Replace(row), "//") {
				continue
			}
		}
		if strings.Index(row, "#") != -1 {
			if strings.HasPrefix(rp1.Replace(row), "#") {
				continue
			}
		}
		if len(row) > 0 {
			contents = append(contents, row)
		}
	}
	return strings.Join(contents, "\n")
}
