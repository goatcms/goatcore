package goatnet

import "net"

// GetFreePort return free port number
func GetFreePort() (port int, err error) {
	var (
		addr  *net.TCPAddr
		laddr *net.TCPListener
	)
	if addr, err = net.ResolveTCPAddr("tcp", "localhost:0"); err != nil {
		return port, err
	}
	if laddr, err = net.ListenTCP("tcp", addr); err != nil {
		return port, err
	}
	defer laddr.Close()
	return laddr.Addr().(*net.TCPAddr).Port, nil
}
