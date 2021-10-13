package server

import (

	"PackAndGo.com/trip/V1/pkg/health"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//Server Interface, always is better to use interfaces for testing and separate implementation from definition
type Server interface {
	Start()
}

type ServerImpl struct {
	tripHandler TripHandler
}

//"Constructor" method, this let the users ignore how to create the implementation if they don't want to
func NewServer () Server {
	return &ServerImpl{NewTriphandler()}
}

//Start Function, will map all paths with handlers and stand up the server
func (s *ServerImpl) Start () {
    router := gin.Default()
    router.GET("/health", health.Check)
    router.GET("/trip", s.tripHandler.List)
    router.POST("/trip", s.tripHandler.Add)
    router.GET("/trip/:id", s.tripHandler.Get)
    router.Run(viper.GetString("HOST_EXPORT"))
}