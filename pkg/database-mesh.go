package main

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/SphereEx/database-mesh/pkg/server"
	"github.com/SphereEx/database-mesh/pkg/server/config"
	"github.com/mlycore/log"
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
