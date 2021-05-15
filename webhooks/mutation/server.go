package main

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/mlycore/log"
	admission "k8s.io/api/admission/v1"
)

// DefaultCert path
const (
	DefaultCertFile = "/certs/cert.pem"
	DefaultKeyFile  = "/certs/key.pem"
)

func main() {
	NewController(&restful.WebService{}).API().Register()
	log.Infof("server start up")
	http.ListenAndServeTLS(":443", DefaultCertFile, DefaultKeyFile, nil)
}

// NewController returns a new controller
func NewController(ws *restful.WebService) *Controller {
	return &Controller{WebService: ws, Injector: &Injector{}}
}

// Controller defines a Sidecar injector controller
type Controller struct {
	*restful.WebService
	Injector *Injector
}

// API register urls and controllers
func (c *Controller) API() *Controller {
	c.Path("/")

	c.Route(c.POST("/inject").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON).
		Reads(&admission.AdmissionReview{}).
		Writes(nil).
		Returns(200, "OK", nil).
		To(c.Injector.Do))
	return c
}

// Register register apis
func (c *Controller) Register() *Controller {
	restful.Add(c.WebService)
	return c
}
