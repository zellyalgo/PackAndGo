package trip

import (
	"errors"

	"PackAndGo.com/trip/V1/pkg/api"
)

//Trip Manager Interface, to separate implementation from interface
type TripManager interface {
	List(start, size int32) 	([]api.Trip, int32)
	Add(trip api.Trip) 			(*api.Trip, error)
	Get(id int32)				*api.Trip
}

func New() TripManager {
	cities := NewCitiesManager()
	return &TripManagerImpl{citiesManager: cities, trips: fillTrips(cities)}
}

type TripManagerImpl struct {
	citiesManager CitiesManager
	trips map[int32]api.Trip
}

//List the elements, from start with size as max elements to return, will return the number of trips too
func (t * TripManagerImpl) List (start, size int32) ([]api.Trip, int32) {
	tripsLen := int32(len(t.trips))
	if start > tripsLen {
		return []api.Trip{}, tripsLen
	}
	size += start
	if size > tripsLen {
		size = tripsLen
	}
	trips := make([]api.Trip, size-start)
	for i := start+1; i <= size; i++ {
		trips[i-start-1] = t.trips[i]
	} 
	return trips, tripsLen
}

//Add the trip to the map, this only happends if origin and destination exist, if not will return error
func (t * TripManagerImpl) Add (trip api.Trip) (*api.Trip, error) {
	trip.Id = int32(len(t.trips)) + 1
	trip.Origin = t.citiesManager.Get(trip.Origin.Id)
	if trip.Origin == nil {
		return nil, errors.New("Origin does not exist")
	}
	trip.Destination = t.citiesManager.Get(trip.Destination.Id)
	if trip.Destination == nil {
		return nil, errors.New("Destination does not exist")
	}
	t.trips[trip.Id] = trip
	return &trip, nil
}

//Get the trip id, this will return nil if not exist
func (t * TripManagerImpl) Get (id int32) *api.Trip {
	trip := t.trips[id]
	if trip.Origin == nil {
		return nil
	}
	return &trip
}

//this funtion creates the map, in future this function should be replaced by BBDD
func fillTrips (citiesManager CitiesManager) map[int32]api.Trip {
	return map[int32]api.Trip {
		1: api.Trip{Id: 1, Origin: citiesManager.Get(1), Destination: citiesManager.Get(2), Dates: "Mon Tue Wed Fri", Price: 40.55},
		2: api.Trip{Id: 2, Origin: citiesManager.Get(2), Destination: citiesManager.Get(1), Dates: "Sat Sun", Price: 40.55},
		3: api.Trip{Id: 3, Origin: citiesManager.Get(3), Destination: citiesManager.Get(6), Dates: "Mon Tue Wed Thu Fri", Price: 32.10},
	}
}