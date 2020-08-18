package jsonmsg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

const (
	// StatusOK --
	StatusOK = 200
	// StatusFailed --
	StatusFailed = 400
)

// JSONData --
type JSONData struct {
	Status interface{} `json:"status"`
	Time   string      `json:"time"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data,omitempty"`
}

// Marshal --
func Marshal(status interface{}, msg string, d interface{}) (data []byte) {
	_, intCheck := status.(int)
	_, stringCheck := status.(string)
	if !intCheck && !stringCheck {
		return []byte(fmt.Sprintf("{\"status\":204,\"time\":\"%s\",\"msg\":\"failed\"}", time.Now().Format("2006-01-02 15:04:05")))
	}
	temp := JSONData{
		Status: status,
		Time:   time.Now().Format("2006-01-02 15:04:05"),
		Msg:    msg,
		Data:   d,
	}
	data, err := json.Marshal(temp)
	if err != nil {
		log.Println(err)
		if _, ok := status.(int); ok {
			return []byte(fmt.Sprintf("{\"status\":204,\"time\":\"%s\",\"msg\":\"failed\"}", time.Now().Format("2006-01-02 15:04:05")))
		}
		return []byte(fmt.Sprintf("{\"status\":\"204\",\"time\":\"%s\",\"msg\":\"failed\"}", time.Now().Format("2006-01-02 15:04:05")))
	}
	return data
}

// Unmarshal --
func Unmarshal(data []byte, d interface{}) (res JSONData, err error) {
	res = JSONData{
		Data: d,
	}
	err = json.Unmarshal(data, &res)
	return res, err
}

// UnmarshalWithInt --
func UnmarshalWithInt(data []byte, d interface{}) (res JSONData, err error) {
	var status int
	res = JSONData{
		Data: d,
	}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return res, err
	}
	statusFloat64, ok := res.Status.(float64)
	if !ok {
		return res, fmt.Errorf("status not int")
	}
	status = int(statusFloat64)
	res.Status = status
	return res, err
}

// UnmarshalWithString --
func UnmarshalWithString(data []byte, d interface{}) (res JSONData, err error) {
	res = JSONData{
		Data: d,
	}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return res, err
	}
	statusString, ok := res.Status.(string)
	if !ok {
		return res, fmt.Errorf("status not string")
	}
	res.Status = statusString
	return res, err
}

// FormatJSON 格式化输出JSON
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

const (
	// CodeStatusOK --
	CodeStatusOK = 1
	// CodeFail --
	CodeFail = 1001
	// CodeError --
	CodeError = 1002
	// CodeMiss --
	CodeMiss = 1003
	// CodeNotFound --
	CodeNotFound = 1004
)

// JSONCodeData 业务固定响应头
type JSONCodeData struct {
	// 状态；1 成功
	Code int `json:"code"`
	// 描述
	Message string `json:"message"`
	// 毫秒时间戳
	Time int64 `json:"time"`
	// 内容
	Data interface{} `json:"data,omitempty"`
}

// MarshalCodeData --
func MarshalCodeData(code int, message string, data interface{}) (res []byte) {
	tms := time.Now().UnixNano() / 1e6
	temp := JSONCodeData{
		Code:    code,
		Time:    tms,
		Message: message,
		Data:    data,
	}
	var err error
	res, err = json.Marshal(temp)
	if err != nil {
		log.Println(err)
		return []byte(fmt.Sprintf("{\"code\":1002,\"time\":%d,\"msg\":\"json marshal failed: %s\"}", tms, err.Error()))
	}
	return res
}

// UnmarshalCodeData --
func UnmarshalCodeData(data []byte, d interface{}) (res JSONCodeData, err error) {
	res = JSONCodeData{
		Data: d,
	}
	err = json.Unmarshal(data, &res)
	return res, err
}
