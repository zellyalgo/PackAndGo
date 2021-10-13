package trip

import (
    "bufio"
    "fmt"
    "os"

	"PackAndGo.com/trip/V1/pkg/api"
	"github.com/spf13/viper"
)

//Cities Manager Interface, to separate implementation from interface
type CitiesManager interface {
	Get(id int32) *api.City
}

func NewCitiesManager() CitiesManager {
	return &CitiesManagerImpl{cities: getCities()}
}

type CitiesManagerImpl struct {
	cities map[int32]*api.City //cities map to reduce complextity when search by id
}

//Get, will return a complete city (id and name), given an id
func (c *CitiesManagerImpl) Get(id int32) *api.City {
	return c.cities[id]
}

//This private function will read the cities.txt file and convert to map, in future iteractions this should be in BBDD and this map can be deleted and changed to a cache
func getCities () map[int32]*api.City {
	cities := make(map[int32]*api.City, 0)
    file, err := os.Open(viper.GetString("CITIES_PATH"))
    if err != nil {
        fmt.Printf("%v\n", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    index := int32(1)
    for scanner.Scan() {
    	cities[index] = &api.City{
    		Id: index,
    		Name: scanner.Text(),
    	}
    	index++
    }

    if err := scanner.Err(); err != nil {
        fmt.Printf("%v\n", err)
    }

    return cities
}

