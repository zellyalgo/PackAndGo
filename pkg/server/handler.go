package server

import (
	"net/http"
    "strconv"

	"PackAndGo.com/trip/V1/pkg/api"
	"PackAndGo.com/trip/V1/pkg/trip"
	"github.com/gin-gonic/gin"
)

//Trip handler interface, always is better to use interfaces for testing and separate implementation from definition
type TripHandler interface {
	List 	(c *gin.Context)
	Add 	(c *gin.Context)
	Get 	(c *gin.Context)
}

//"Constructor" method, this let the users ignore how to create the implementation if they don't want to
func NewTriphandler() TripHandler {
	manager := trip.New()
	return &TripHandlerImpl{tripManager: manager}
}

type TripHandlerImpl struct {
	tripManager trip.TripManager
}

//Handler List, this only get the params, and call the manager, no bussiness logic here, if some problems with queryParams throws Bad Request
func (t *TripHandlerImpl) List (c *gin.Context) {
	limitBasic, err := strconv.Atoi(c.DefaultQuery("limit", "100"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, api.TripList{})
		return
	}
	startBasic, err := strconv.Atoi(c.DefaultQuery("start", "0"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, api.TripList{})
		return
	}
	limit := int32(limitBasic)
	start := int32(startBasic)
	list, total := t.tripManager.List(start, limit)
	tripList := api.TripList{
		TripList: list,
		Total: total,
		Limit: limit,
	}
	c.IndentedJSON(http.StatusOK, tripList)
}

//Handler Add, this only parse the body, and call the manager, no bussiness logic here, if some problem with parsing or add, will throw a Bad Request to the client
func (t *TripHandlerImpl) Add (c *gin.Context) {
	var trip api.Trip
	if err := c.BindJSON(&trip); err != nil {
		c.IndentedJSON(http.StatusBadRequest, api.Trip{})
        return
    }
	tripResponse, err := t.tripManager.Add(trip)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, api.Trip{})
		return
	}
	c.IndentedJSON(http.StatusCreated, tripResponse)
}

//Handler Get, this only get the param, and call the manager, no bussiness logic here, if some problem getting trip, will throw a Bad Request to the client
func (t *TripHandlerImpl) Get (c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, api.Trip{})
		return
	}
	c.IndentedJSON(http.StatusOK, t.tripManager.Get(int32(id)))
}