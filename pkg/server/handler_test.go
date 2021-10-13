package server

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"net/http"
	"testing"
	
	"PackAndGo.com/trip/V1/pkg/api"
	"github.com/gin-gonic/gin"
)

func TestListHandler (t *testing.T) {
	tripManagerMock := &TripManagerMock{}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	tripHandler := &TripHandlerImpl{tripManager: tripManagerMock}
	tripHandler.List(c)

	actualCode := w.Code
	expectedCode := 200
	if actualCode != expectedCode {
		t.Fatalf(`want [%v] but got [%v]`, expectedCode, actualCode)
	}

	expected := `{
    "trips": [
        {
            "id": 1,
            "origin": {
                "id": 1,
                "name": "Barcelona"
            },
            "destination": {
                "id": 2,
                "name": "Seville"
            },
            "dates": "Mon Tue Wed Fri",
            "price": 40.55
        },
        {
            "id": 2,
            "origin": {
                "id": 2,
                "name": "Seville"
            },
            "destination": {
                "id": 1,
                "name": "Barcelona"
            },
            "dates": "Sat Sun",
            "price": 40.55
        },
        {
            "id": 3,
            "origin": {
                "id": 3,
                "name": "Madrid"
            },
            "destination": {
                "id": 6,
                "name": "Malaga"
            },
            "dates": "Mon Tue Wed Thu Fri",
            "price": 32.1
        }
    ],
    "total": 3,
    "limit": 100
}`
	actual := w.Body.String()
	if actual != expected {
		t.Fatalf(`want [%v] but got [%v]`, expected, actual)
	}
}
func TestListHandlerBadQueryLimit (t *testing.T) {
	tripManagerMock := &TripManagerMock{}
	tripHandler := &TripHandlerImpl{tripManager: tripManagerMock}

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
    c, r := gin.CreateTestContext(w)
    r.GET("/trip", tripHandler.List)
    c.Request, _ = http.NewRequest(http.MethodGet, "/trip?limit=hi&start=5", nil)
    r.ServeHTTP(w, c.Request)

	actualCode := w.Code
	expectedCode := 400
	if actualCode != expectedCode {
		t.Fatalf(`want [%v] but got [%v]`, actualCode, expectedCode)
	}

	expected := `{}`
	actual := w.Body.String()
	if actual != expected {
		t.Fatalf(`want [%v] but got [%v]`, expected, actual)
	}
}
func TestListHandlerBadQueryStart (t *testing.T) {
	tripManagerMock := &TripManagerMock{}
	tripHandler := &TripHandlerImpl{tripManager: tripManagerMock}

	gin.SetMode(gin.TestMode)
    w := httptest.NewRecorder()
    c, r := gin.CreateTestContext(w)
    r.GET("/trip", tripHandler.List)
    c.Request, _ = http.NewRequest(http.MethodGet, "/trip?limit=100&start=meh", nil)
    r.ServeHTTP(w, c.Request)

	actualCode := w.Code
	expectedCode := 400
	if actualCode != expectedCode {
		t.Fatalf(`want [%v] but got [%v]`, actualCode, expectedCode)
	}

	expected := `{}`
	actual := w.Body.String()
	if actual != expected {
		t.Fatalf(`want [%v] but got [%v]`, expected, actual)
	}
}

func TestAddTrip (t *testing.T) {
	tripManagerMock := &TripManagerMock{}
	tripHandler := &TripHandlerImpl{tripManager: tripManagerMock}

	gin.SetMode(gin.TestMode)
    w := httptest.NewRecorder()
    c, r := gin.CreateTestContext(w)
    r.POST("/trip", tripHandler.Add)
    tripJson := `{"id": 1, "origin": {"id": 1,"name": "Barcelona"},"destination": {"id": 2,"name": "Seville"},"dates": "Mon Tue Wed Fri","price": 40.55}`
    c.Request, _ = http.NewRequest(http.MethodPost, "/trip", bytes.NewBuffer([]byte(tripJson)))
    r.ServeHTTP(w, c.Request)

	actualCode := w.Code
	expectedCode := 201
	if actualCode != expectedCode {
		t.Fatalf(`want [%v] but got [%v]`, expectedCode, actualCode)
	}
	expected := `{
    "id": 1,
    "origin": {
        "id": 1,
        "name": "Barcelona"
    },
    "destination": {
        "id": 2,
        "name": "Seville"
    },
    "dates": "Mon Tue Wed Fri",
    "price": 40.55
}`
	actual := w.Body.String()
	if actual != expected {
		t.Fatalf(`want [%v] but got [%v]`, expected, actual)
	}
}
func TestAddBadTrip (t *testing.T) {
	tripManagerMock := &TripManagerMock{}
	tripHandler := &TripHandlerImpl{tripManager: tripManagerMock}

	gin.SetMode(gin.TestMode)
    w := httptest.NewRecorder()
    c, r := gin.CreateTestContext(w)
    r.POST("/trip", tripHandler.Add)
    tripJson := `{"id": 2, "origin": {"id": 1,"name": "Barcelona"},"destination": {"id": 2,"name": "Seville"},"dates": "Mon Tue Wed Fri","price": 40.55}`
    c.Request, _ = http.NewRequest(http.MethodPost, "/trip", bytes.NewBuffer([]byte(tripJson)))
    r.ServeHTTP(w, c.Request)

	actualCode := w.Code
	expectedCode := 400
	if actualCode != expectedCode {
		t.Fatalf(`want [%v] but got [%v]`, expectedCode, actualCode)
	}
	expected := `{}`
	actual := w.Body.String()
	if actual != expected {
		t.Fatalf(`want [%v] but got [%v]`, expected, actual)
	}
}

func TestGetTrip (t *testing.T) {
	tripManagerMock := &TripManagerMock{}
	tripHandler := &TripHandlerImpl{tripManager: tripManagerMock}

	gin.SetMode(gin.TestMode)
    w := httptest.NewRecorder()
    c, r := gin.CreateTestContext(w)
    r.GET("/trip/:id", tripHandler.Get)
    c.Request, _ = http.NewRequest(http.MethodGet, "/trip/1", nil)
    r.ServeHTTP(w, c.Request)

	actualCode := w.Code
	expectedCode := 200
	if actualCode != expectedCode {
		t.Fatalf(`want [%v] but got [%v]`, actualCode, expectedCode)
	}
	expected := `{
    "id": 1,
    "origin": {
        "id": 1,
        "name": "Barcelona"
    },
    "destination": {
        "id": 2,
        "name": "Seville"
    },
    "dates": "Mon Tue Wed Fri",
    "price": 40.55
}`
	actual := w.Body.String()
	if actual != expected {
		t.Fatalf(`want [%v] but got [%v]`, expected, actual)
	}
}
func TestGetBadRequestTrip (t *testing.T) {
	tripManagerMock := &TripManagerMock{}
	tripHandler := &TripHandlerImpl{tripManager: tripManagerMock}

	gin.SetMode(gin.TestMode)
    w := httptest.NewRecorder()
    c, r := gin.CreateTestContext(w)
    r.GET("/trip/:id", tripHandler.Get)
    c.Request, _ = http.NewRequest(http.MethodGet, "/trip/hi", nil)
    r.ServeHTTP(w, c.Request)

	actualCode := w.Code
	expectedCode := 400
	if actualCode != expectedCode {
		t.Fatalf(`want [%v] but got [%v]`, actualCode, expectedCode)
	}
	expected := `{}`
	actual := w.Body.String()
	if actual != expected {
		t.Fatalf(`want [%v] but got [%v]`, expected, actual)
	}
}


type TripManagerMock struct {
}

func (t *TripManagerMock) List (start, size int32) ([]api.Trip, int32) {
	return []api.Trip{
		api.Trip{Id: 1, Origin: &api.City{Id: 1, Name: "Barcelona"}, Destination: &api.City{Id: 2, Name: "Seville"}, Dates: "Mon Tue Wed Fri", Price: 40.55},
		api.Trip{Id: 2, Origin: &api.City{Id: 2, Name: "Seville"}, Destination: &api.City{Id: 1, Name: "Barcelona"}, Dates: "Sat Sun", Price: 40.55},
		api.Trip{Id: 3, Origin: &api.City{Id: 3, Name: "Madrid"}, Destination: &api.City{Id: 6, Name: "Malaga"}, Dates: "Mon Tue Wed Thu Fri", Price: 32.10},
	}, 3
}

func (t *TripManagerMock) Add (trip api.Trip) (*api.Trip, error) {
	if trip.Id == 1 {
		return &api.Trip{
			Id: 1, 
			Origin: &api.City{Id: 1, Name: "Barcelona"}, 
			Destination: &api.City{Id: 2, Name: "Seville"}, 
			Dates: "Mon Tue Wed Fri", 
			Price: 40.55,
		}, nil
	} else {
		return &api.Trip{}, errors.New("Destination does not exist")
	}
	
}

func (t *TripManagerMock) Get(id int32)	*api.Trip {
	return &api.Trip{
		Id: 1, 
		Origin: &api.City{Id: 1, Name: "Barcelona"}, 
		Destination: &api.City{Id: 2, Name: "Seville"}, 
		Dates: "Mon Tue Wed Fri", 
		Price: 40.55,
	}
}