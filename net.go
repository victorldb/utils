package utils

import (
	"net"
	"strings"
	"time"
)

// TimeoutDialer returns functions of connection dialer with timeout settings for http.Transport Dial field.
func TimeoutDialer(cTimeout time.Duration, rwTimeout time.Duration) func(net, addr string) (c net.Conn, err error) {
	return func(netw, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(netw, addr, cTimeout)
		if err != nil {
			return nil, err
		}
		err = conn.SetDeadline(time.Now().Add(rwTimeout))
		return conn, err
	}
}

// CheckTCPAddress 检查TCP地址
func CheckTCPAddress(s string) (res string, err error) {
	ind1 := strings.Index(s, ".")
	if ind1 == -1 {
		if ind2 := strings.Index(s, ":"); ind2 == -1 {
			s = ":" + s
		}
	}

	// 不使用net.SplitHostPort原因是没有对IP和端口合法性进行检测
	_, err = net.ResolveTCPAddr("tcp", s)
	if err != nil {
		return res, err
	}
	return s, nil
}
