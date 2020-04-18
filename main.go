package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

// Place : place schema
type Place struct {
	Features []Features `json: "features"`
}

// Features : mapping property to reach geometry
type Features struct {
	Geometry Geometry `json: "geometry"`
}

// Geometry : handle geometric coordinates
type Geometry struct {
	Coordinates []float64 `json: "coordinates"`
}

func main() {
	r := gin.Default()
	r.GET("/", search)
	r.Run()
}

func search(c *gin.Context) {
	searchURL := "https://api-adresse.data.gouv.fr/search/?q="
	Lpa := c.Request.URL.Query()["pA"]
	Lpb := c.Request.URL.Query()["pB"]

	if len(Lpa) != 1 || len(Lpb) != 1 {
		c.AbortWithError(http.StatusBadRequest, errors.New("Bad request, expecting query string pA and pB that are place search strings")).SetType(gin.ErrorTypePublic)
		return
	}
	Pa := strings.ReplaceAll(Lpa[0], " ", "+")
	Pb := strings.ReplaceAll(Lpb[0], " ", "+")
	var placeA Place
	var placeB Place

	var wg sync.WaitGroup
	wg.Add(2)
	go getPlace(searchURL+Pa, &placeA, &wg)
	go getPlace(searchURL+Pb, &placeB, &wg)
	wg.Wait()

	result := distance(placeA.Features[0].Geometry.Coordinates[0], placeA.Features[0].Geometry.Coordinates[1], placeB.Features[0].Geometry.Coordinates[0], placeB.Features[0].Geometry.Coordinates[1])
	fmt.Println(fmt.Sprintf("Search %f -> %s | %s", result, Pa, Pb))
	c.JSON(200, gin.H{"result": result})
}

func getPlace(url string, place *Place, wg *sync.WaitGroup) {
	response, err := http.Get(url)
	if err != nil {
		log.Println("Error searching text")
	}
	defer response.Body.Close()

	json.NewDecoder(response.Body).Decode(&place)

	wg.Done()
}

func distance(lat1 float64, lng1 float64, lat2 float64, lng2 float64) float64 {
	//This is a subset of the original code from https://www.geodatasource.com/developers/go that only fits this project needs
	const PI float64 = 3.141592653589793

	radlat1 := float64(PI * lat1 / 180)
	radlat2 := float64(PI * lat2 / 180)

	theta := float64(lng1 - lng2)
	radtheta := float64(PI * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist) * 180 / PI * 60 * 1.1515

	return dist * 1.609344
}
