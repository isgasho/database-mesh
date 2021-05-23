package connector

import (
	"io"
	"net"

	"github.com/mlycore/log"
)

// Connector builds an connector for different database backend
type Connector interface {
	Run() error
}

// ConnectorCore defines a standard connector core
type ConnectorCore struct {
	Host   string
	ConnCh chan *net.TCPConn
}

// Run means running this connector core, listening to accepting tcp requests
// Once setup a connection, starts two goroutine to handle full duplex communications
func (c *ConnectorCore) Run() error {
	addr, err := net.ResolveTCPAddr("tcp", c.Host)
	if err != nil {
		return err
	}

	for {
		select {
		case conn := <-c.ConnCh:
			{
				log.Infof("connection setup from %s to %s", conn.RemoteAddr().String(), conn.LocalAddr().String())
				proxy, err := net.DialTCP("tcp", nil, addr)
				if err != nil {
					log.Fatalf("dial mysql error: %s", err)
				}
				log.Infof("proxy connection setup from %s to %s", proxy.LocalAddr().String(), proxy.RemoteAddr().String())
				go pipe(conn, proxy)
				go pipe(proxy, conn)
			}
		}
	}
}

func pipe(src io.Reader, dst io.Writer) int64 {
	n, err := io.Copy(dst, src)
	if err != nil {
		log.Errorf("copy error: %s", err)
	}
	return n
}
