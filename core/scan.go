package core

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

func isOpen(host string, port int, timeout time.Duration) bool {
	time.Sleep(time.Millisecond * 1)
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), timeout)
	if err == nil {
		_ = conn.Close()
		return true
	}
	return false
}

func Scan(host string, port int, project string, name string) {
	wg := &sync.WaitGroup{}
	timeout := time.Millisecond * 500

	wg.Add(1)
	go func(p int) {
		opened := isOpen(host, p, timeout)
		if !opened {
			logrus.Infof("Project: %s, Server: %s, Host: %s, Port: %d, the connection to the server failed!\n", project, name, host, port)
		} else {
			logrus.Infof("Project: %s, Server: %s, Host: %s, Port: %d, the connection to the server success!\n", project, name, host, port)
		}
		wg.Done()
	}(port)

	wg.Wait()
}
