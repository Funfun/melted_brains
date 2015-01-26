package memcache

import (
	"bufio"
	"net"
	"strings"
	"time"
)

type Connection struct {
	conn     net.Conn
	buffered bufio.ReadWriter
	timeout  time.Duration
}

type Result struct {
	Key   string
	Value []byte
	Flags uint16
	Cas   uint64
}

func Connect(address string, timeout time.Duration) (conn *Connection, err error) {
	var network string
	if strings.Contains(address, "/") {
		network = "unix"
	} else {
		network = "tcp"
	}
	var nc net.Conn
	nc, err = net.DialTimeout(network, address, timeout)
	if err != nil {
		return nil, err
	}
	return &Connection{
		conn: nc,
		buffered: bufio.ReadWriter{
			Reader: bufio.NewReader(nc),
			Writer: bufio.NewWriter(nc),
		},
		timeout: timeout,
	}, nil
}

func (mc *Connection) Close() {
	mc.conn.Close()
}

func (mc *Connection) Get(keys ...string) (results []Result, err error) {
	defer handleError(&err)
	results = mc.get("get", keys)
	return
}

func (mc *Connection) Gets(keys ...string) (results []Result, err error) {
	defer handleError(&err)
	results = mc.get("gets", keys)
	return
}
