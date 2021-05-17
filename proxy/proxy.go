package main

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/mlycore/log"

	"github.com/SphereEx/database-mesh/proxy/server"
	"github.com/SphereEx/database-mesh/proxy/server/config"
	//_ "github.com/mkevac/debugcharts"
)

func init() {
	config.NewServerConfigOrDie()
	log.SetLevel(config.GlobalServerConfig().Loglevel)
}

func main() {
	conf := config.GlobalServerConfig()
	log.Infof("database-mesh is spinning up...")

	s := server.NewServer(conf)
	go s.Run()
	if conf.DebugCharts {
		log.Errorln(http.ListenAndServe("localhost:6060", nil))
	}
}
