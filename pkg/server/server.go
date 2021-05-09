package server

import (
	"fmt"
	"net"

	"github.com/SphereEx/database-mesh/pkg/connector"
	"github.com/SphereEx/database-mesh/pkg/server/config"
	"github.com/mlycore/log"
)

// ProxyProtocol defines currently supported common database type protocol
type ProxyProtocol string

const (
	ProxyProtocolMySQL = "MySQL"
	ProxyProtocolRedis = "Redis"
)

// Server setup a tcp server for backend runner
type Server struct {
	Proxies map[ProxyProtocol]Proxy
}

// Proxy route packets from port to different connectors
type Proxy struct {
	Port      string
	Connector connector.Connector
	ConnCh    chan *net.TCPConn
}

// NewServer returns a new server
func NewServer(conf config.Config) *Server {
	s := &Server{
		Proxies: make(map[ProxyProtocol]Proxy),
	}
	if conf.MySQL != nil {
		ch := make(chan *net.TCPConn, 128)
		s.Proxies[ProxyProtocolMySQL] = Proxy{
			Port: conf.MySQL.Port,
			Connector: &connector.MySQLConnector{
				ConnectorCore: connector.ConnectorCore{
					Host:   conf.MySQL.Proxy,
					ConnCh: ch,
				},
			},
			ConnCh: ch,
		}
	}
	if conf.Redis != nil {
		ch := make(chan *net.TCPConn, 128)
		s.Proxies[ProxyProtocolRedis] = Proxy{
			Port: conf.Redis.Port,
			Connector: &connector.RedisConnector{
				ConnectorCore: connector.ConnectorCore{
					Host:   conf.Redis.Proxy,
					ConnCh: ch,
				},
			},
			ConnCh: ch,
		}
	}
	return s
}

// Run spinning up the server
func (s *Server) Run() error {
	for p := range s.Proxies {
		go s.listen(p, s.Proxies[p].Port)
		go s.Proxies[p].Connector.Run()
	}
	select {}
}

func (s *Server) listen(protocol ProxyProtocol, port string) error {
	bind := fmt.Sprintf(":%s", port)
	log.Debugf("bind: %s", bind)
	addr, err := net.ResolveTCPAddr("tcp", bind)
	if err != nil {
		return err
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return err
	}

	defer listener.Close()
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Fatalf("accept: %s", err)
		}
		log.Infof("connection setup from %s to %s", conn.RemoteAddr().String(), conn.LocalAddr().String())
		s.Proxies[protocol].ConnCh <- conn
	}
}
