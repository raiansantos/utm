package main

import (
	"fmt"
	"github.com/raiansantos/utm"
)

func main() {
	easting, northing, zoneNumber, zoneLetter, err := utm.FromLatLon(-28.502990, -49.01296, false)
	if err != nil {
		panic(err.Error())
	}

	text := fmt.Sprintf("Easting: %.0f; Northing: %.0f; ZoneNumber: %d; ZoneLetter: %s;",
		easting, northing, zoneNumber, zoneLetter)
	fmt.Println(text)

	latitude, longitude, err := utm.ToLatLon(694478, 6845477, 22, "", false)
	fmt.Println(fmt.Sprintf("Latitude: %.5f; Longitude: %.5f;", latitude, longitude))
}
