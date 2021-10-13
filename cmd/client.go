package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"os"

	"PackAndGo.com/trip/V1/pkg/api"
)

func main() {
	time.Sleep(time.Second)
	tripList := GetList()
	fmt.Printf("Listed Elements: %v\n", tripList)
	trip := Add()
	fmt.Printf("Added Elements: %v\n", trip)
	tripList = GetList()
	fmt.Printf("Listed Elements: %v\n", tripList)
	trip = Get()
	fmt.Printf("New Element: %v\n", trip)
}

func GetList() *api.TripList {
	req, err := http.NewRequest("GET", os.Getenv("HOST_API")+"/trip", nil)
	if err != nil {
	   fmt.Printf("%v\n", err)
	}

	body, err := do(req)
	if err != nil {
	    fmt.Printf("%v\n", err)
		return nil
	}
	var tripList api.TripList
	err = json.Unmarshal(body, &tripList)
	if err != nil {
	    fmt.Printf("%v\n", err)
		return nil
	}
	return &tripList
}

func Add () *api.Trip {
	j, err := json.Marshal(api.Trip {
		Origin: &api.City{Id: 1},
		Destination: &api.City{Id: 6},
		Dates: "Mon Tue Wed Fri",
		Price: 100.55,
	})
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil
	}

	req, err := http.NewRequest("POST", os.Getenv("HOST_API")+"/trip", bytes.NewBuffer(j))
	if err != nil {
	   fmt.Printf("%v\n", err)
	   return nil
	}

	trip, err := processTripResponse(req)
	if err != nil {
	   fmt.Printf("%v\n", err)
	   return nil
	}
	return trip
}

func Get () *api.Trip {
	req, err := http.NewRequest("GET", os.Getenv("HOST_API")+"/trip/4", nil)
	if err != nil {
	   fmt.Printf("%v\n", err)
	}

	trip, err := processTripResponse(req)
	if err != nil {
	   fmt.Printf("%v\n", err)
	}
	return trip
}

func processTripResponse (req *http.Request) (*api.Trip, error) {
	body, err := do(req)
	if err != nil {
		return nil, err
	}
	var trip api.Trip
	err = json.Unmarshal(body, &trip)
	if err != nil {
		return nil, err
	}
	return &trip, nil
}

func do(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if 300 < res.StatusCode {
		return nil, fmt.Errorf("%v, %s", res.StatusCode, body)
	}
	return body, nil
}