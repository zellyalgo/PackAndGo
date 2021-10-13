package trip

import (
	"testing"

	"PackAndGo.com/trip/V1/pkg/api"
	"github.com/spf13/viper"
)

func TestGetCity (t *testing.T) {
	viper.SetDefault("CITIES_PATH", "../../resources/cities.txt")
	cityManager := NewCitiesManager()
	actual := cityManager.Get(1)
	expected := api.City {
		Id: 1,
		Name: "Barcelona",
	}
	if *actual != expected {
		t.Fatalf(`want [%v] but got [%v]`, expected, actual)
	}
}