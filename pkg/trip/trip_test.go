package trip

import (
	"testing"

	"PackAndGo.com/trip/V1/pkg/api"
	"github.com/google/go-cmp/cmp"
)

func TestListTrip (t *testing.T) {
	tripManager := New()
	actual, actualSize := tripManager.List(0, 3)
	expected := []api.Trip{
		api.Trip{Id: 1, Origin: &api.City{Id: 1, Name: "Barcelona"}, Destination: &api.City{Id: 2, Name: "Seville"}, Dates: "Mon Tue Wed Fri", Price: 40.55},
		api.Trip{Id: 2, Origin: &api.City{Id: 2, Name: "Seville"}, Destination: &api.City{Id: 1, Name: "Barcelona"}, Dates: "Sat Sun", Price: 40.55},
		api.Trip{Id: 3, Origin: &api.City{Id: 3, Name: "Madrid"}, Destination: &api.City{Id: 6, Name: "Malaga"}, Dates: "Mon Tue Wed Thu Fri", Price: 32.10},
	}

	if !cmp.Equal(actual, expected) {

		t.Fatalf(`want %v but got %v`, expected, actual)
	}
	if actualSize != int32(3) {
		t.Fatalf(`want 3 but got [%v]`, actualSize)	
	}
}
func TestListTripLessElements (t *testing.T) {
	tripManager := New()
	actual, actualSize := tripManager.List(0, 2)
	expected := []api.Trip{
		api.Trip{Id: 1, Origin: &api.City{Id: 1, Name: "Barcelona"}, Destination: &api.City{Id: 2, Name: "Seville"}, Dates: "Mon Tue Wed Fri", Price: 40.55},
		api.Trip{Id: 2, Origin: &api.City{Id: 2, Name: "Seville"}, Destination: &api.City{Id: 1, Name: "Barcelona"}, Dates: "Sat Sun", Price: 40.55},
	}
	if !cmp.Equal(actual, expected) {
		t.Fatalf(`want [%v] but got [%v]`, expected, actual)
	}
	if actualSize != int32(3) {
		t.Fatalf(`want 3 but got [%v]`, actualSize)	
	}
}
func TestListTripMoreElements (t *testing.T) {
	tripManager := New()
	actual, actualSize := tripManager.List(0, 6)
	expected := []api.Trip{
		api.Trip{Id: 1, Origin: &api.City{Id: 1, Name: "Barcelona"}, Destination: &api.City{Id: 2, Name: "Seville"}, Dates: "Mon Tue Wed Fri", Price: 40.55},
		api.Trip{Id: 2, Origin: &api.City{Id: 2, Name: "Seville"}, Destination: &api.City{Id: 1, Name: "Barcelona"}, Dates: "Sat Sun", Price: 40.55},
		api.Trip{Id: 3, Origin: &api.City{Id: 3, Name: "Madrid"}, Destination: &api.City{Id: 6, Name: "Malaga"}, Dates: "Mon Tue Wed Thu Fri", Price: 32.10},
	}
	if !cmp.Equal(actual, expected) {
		t.Fatalf(`want [%v] but got [%v]`, expected, actual)
	}
	if actualSize != int32(3) {
		t.Fatalf(`want 3 but got [%v]`, actualSize)	
	}
}
func TestListTripLimitElements (t *testing.T) {
	tripManager := New()
	actual, actualSize := tripManager.List(2, 1)
	expected := []api.Trip{
		api.Trip{Id: 3, Origin: &api.City{Id: 3, Name: "Madrid"}, Destination: &api.City{Id: 6, Name: "Malaga"}, Dates: "Mon Tue Wed Thu Fri", Price: 32.10},
	}
	if !cmp.Equal(actual, expected) {
		t.Fatalf(`want [%v] but got [%v]`, expected, actual)
	}
	if actualSize != int32(3) {
		t.Fatalf(`want 3 but got [%v]`, actualSize)	
	}
}
func TestListTripLimitOverheadElements (t *testing.T) {
	tripManager := New()
	actual, actualSize := tripManager.List(2, 3)
	expected := []api.Trip{
		api.Trip{Id: 3, Origin: &api.City{Id: 3, Name: "Madrid"}, Destination: &api.City{Id: 6, Name: "Malaga"}, Dates: "Mon Tue Wed Thu Fri", Price: 32.10},
	}
	if !cmp.Equal(actual, expected) {
		t.Fatalf(`want [%v] but got [%v]`, expected, actual)
	}
	if actualSize != int32(3) {
		t.Fatalf(`want 3 but got [%v]`, actualSize)	
	}
}
func TestListTripLimitOutOfBoundElements (t *testing.T) {
	tripManager := New()
	actual, actualSize := tripManager.List(4, 3)
	expected := []api.Trip{}
	if !cmp.Equal(actual, expected) {
		t.Fatalf(`want [%v] but got [%v]`, expected, actual)
	}
	if actualSize != int32(3) {
		t.Fatalf(`want 3 but got [%v]`, actualSize)	
	}
}

func TestGetTrip (t *testing.T) {
	tripManager := New()
	actual := tripManager.Get(1)
	expected := api.Trip{
		Id: 1, 
		Origin: &api.City{Id: 1, Name: "Barcelona"}, 
		Destination: &api.City{Id: 2, Name: "Seville"}, 
		Dates: "Mon Tue Wed Fri", 
		Price: 40.55,
	}

	if !cmp.Equal(*actual, expected) {
		t.Fatalf(`want [%v] but got [%v]`, expected, actual)
	}
}
func TestGetNonExistTrip (t *testing.T) {
	tripManager := New()
	actual := tripManager.Get(9)
	if actual != nil {
		t.Fatalf(`want nil but got [%v]`, actual)
	}
}

func TestAddTrip (t *testing.T) {
	tripManager := New()
	trip := api.Trip {
		Origin: &api.City{Id: 1},
		Destination: &api.City{Id: 6},
		Dates: "Mon Tue Wed Fri",
		Price: 100.55,
	}
	actual, err := tripManager.Add(trip)
	if err != nil {
		t.Fatalf(`want nil but got [%v]`, err)
	}
	expected := api.Trip {
		Id: 4,
		Origin: &api.City{Id: 1, Name: "Barcelona"},
		Destination: &api.City{Id: 6, Name: "Malaga"},
		Dates: "Mon Tue Wed Fri",
		Price: 100.55,
	}
	if !cmp.Equal(*actual, expected) {
		t.Fatalf(`want [%v] but got [%v]`, expected, actual)
	}
}
func TestAddTripNoOrigin (t *testing.T) {
	tripManager := New()
	trip := api.Trip {
		Origin: &api.City{Id: 7},
		Destination: &api.City{Id: 6},
		Dates: "Mon Tue Wed Fri",
		Price: 100.55,
	}
	actual, err := tripManager.Add(trip)
	if actual != nil {
		t.Fatalf(`want nil but got [%v]`, actual)
	}
	expectedErr := "Origin does not exist"
	if err.Error() !=  expectedErr {
		t.Fatalf(`want [%v] but got [%v]`, expectedErr, err)
	}
}
func TestAddTripNoDestination (t *testing.T) {
	tripManager := New()
	trip := api.Trip {
		Origin: &api.City{Id: 6},
		Destination: &api.City{Id: 7},
		Dates: "Mon Tue Wed Fri",
		Price: 100.55,
	}
	actual, err := tripManager.Add(trip)
	if actual != nil {
		t.Fatalf(`want nil but got [%v]`, actual)
	}
	expectedErr := "Destination does not exist"
	if err.Error() !=  expectedErr {
		t.Fatalf(`want [%v] but got [%v]`, expectedErr, err)
	}
}