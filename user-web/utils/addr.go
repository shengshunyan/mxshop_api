package utils

import (
	"go.uber.org/zap"
	"net"
)

// 获取空闲的端口
func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	tcp, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer func(tcp *net.TCPListener) {
		err := tcp.Close()
		if err != nil {
			zap.S().Errorf("failed to close tcp listener: %v", err)
		}
	}(tcp)

	return tcp.Addr().(*net.TCPAddr).Port, nil
}
