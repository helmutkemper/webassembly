package geolocation

type Coordinate struct {
	Latitude     float64
	Longitude    float64
	Accuracy     float64
	ErrorCode    int
	ErrorMessage string
}
