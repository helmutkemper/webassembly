package main

import (
	"fmt"
	"math"
)

const (
	minLatitude  = -85.05112878
	maxLatitude  = 85.05112878
	minLongitude = -180
	maxLongitude = 180
	tileSize     = 256
)

// Clip clips a number to the specified minimum and maximum values.
func clip(n, minValue, maxValue float64) float64 {
	return math.Min(math.Max(n, minValue), maxValue)
}

// ClipByRange clips a number by a specified range.
func clipByRange(n, rangeVal float64) float64 {
	return math.Mod(n, rangeVal)
}

// LatLongToPixelXYOSM converts latitude and longitude to pixel coordinates for OpenStreetMap.
func latLongToPixelXYOSM(latitude, longitude float64, zoomLevel int) (pixelX, pixelY int) {
	mapSize := math.Pow(2, float64(zoomLevel)) * tileSize

	latitude = clip(latitude, minLatitude, maxLatitude)
	longitude = clip(longitude, minLongitude, maxLongitude)

	x := (longitude + 180.0) / 360.0 * float64(int(1<<zoomLevel))
	y := (1.0 - math.Log(math.Tan(latitude*math.Pi/180.0)+1.0/math.Cos(latitude*math.Pi/180.0))/math.Pi) / 2.0 * float64(int(1<<zoomLevel))

	tileX := int(math.Floor(x))
	tileY := int(math.Floor(y))
	pixelX = int(clipByRange((float64(tileX)*tileSize)+((x-float64(tileX))*tileSize), mapSize-1))
	pixelY = int(clipByRange((float64(tileY)*tileSize)+((y-float64(tileY))*tileSize), mapSize-1))

	return pixelX, pixelY
}

// PixelXYToLatLongOSM converts pixel coordinates to latitude and longitude for OpenStreetMap.
func pixelXYToLatLongOSM(pixelX, pixelY int, zoomLevel int) (latitude, longitude float64) {
	mapSize := math.Pow(2, float64(zoomLevel)) * tileSize

	//tileX := int(math.Floor(float64(pixelX) / tileSize))
	//tileY := int(math.Floor(float64(pixelY) / tileSize))

	n := math.Pi - (2.0*math.Pi*clipByRange(float64(pixelY), mapSize-1)/tileSize)/math.Pow(2.0, float64(zoomLevel))

	longitude = (clipByRange(float64(pixelX), mapSize-1)/tileSize)/math.Pow(2.0, float64(zoomLevel))*360.0 - 180.0
	latitude = 180.0 / math.Pi * math.Atan(math.Sinh(n))

	return latitude, longitude
}

func main() {
	// Test LatLongToPixelXYOSM
	lat, lon := 51.5074, -0.1278
	zoomLevel := 17
	pixelX, pixelY := latLongToPixelXYOSM(lat, lon, zoomLevel)
	fmt.Printf("Pixel coordinates: X = %d, Y = %d\n", pixelX, pixelY)

	// Test PixelXYToLatLongOSM
	lat2, lon2 := pixelXYToLatLongOSM(pixelX, pixelY, zoomLevel)
	fmt.Printf("Latitude: %.6f, Longitude: %.6f\n", lat2, lon2)
}
