package html

import (
	"image/color"
	"log"
	"testing"
	"time"
)

func TestTypeToString(t *testing.T) {

	value := TypeToString(
		1.0,
		"",
		"",
	)

	if value.(string) != "1" {
		log.Printf("error: %v", value)
		t.FailNow()
	}

	value = TypeToString(
		10.5,
		"",
		"",
	)

	if value.(string) != "10.5" {
		log.Printf("error: %v", value)
		t.FailNow()
	}

	value = TypeToString(
		float32(1.0),
		"",
		"",
	)

	if value.(string) != "100%" {
		log.Printf("error: %v", value)
		t.FailNow()
	}

	value = TypeToString(
		float32(0.1),
		"",
		"",
	)

	if value.(string) != "10%" {
		log.Printf("error: %v", value)
		t.FailNow()
	}

	value = TypeToString(
		float32(0.05),
		"",
		"",
	)

	if value.(string) != "5%" {
		log.Printf("error: %v", value)
		t.FailNow()
	}

	value = TypeToString(
		0*time.Second,
		"",
		"",
	)

	if value.(string) != "0s" {
		log.Printf("error: %v", value)
		t.FailNow()
	}

	value = TypeToString(
		5*time.Second,
		"",
		"",
	)

	if value.(string) != "5s" {
		log.Printf("error: %v", value)
		t.FailNow()
	}

	value = TypeToString(
		5*time.Minute,
		"",
		"",
	)

	if value.(string) != "5m0s" {
		log.Printf("error: %v", value)
		t.FailNow()
	}

	value = TypeToString(
		5*time.Hour,
		"",
		"",
	)

	if value.(string) != "5h0m0s" {
		log.Printf("error: %v", value)
		t.FailNow()
	}

	value = TypeToString(
		time.Unix(0, 0),
		"",
		"",
	)

	if value.(string) != "1969-12-31 21:00:00 -0300 -03" {
		log.Printf("error: %v", value)
		t.FailNow()
	}

	value = TypeToString(
		color.RGBA{R: 50, G: 100, B: 150, A: 255},
		"",
		"",
	)

	if value.(string) != "rgba(50,100,150,1)" {
		log.Printf("error: %v", value)
		t.FailNow()
	}

	value = TypeToString(
		[]float64{0.0, 1.0, 2.5},
		",",
		";",
	)

	if value.(string) != "0,1,2.5" {
		log.Printf("error: %v", value)
		t.FailNow()
	}

	value = TypeToString(
		[]float32{0.0, 0.1, 1.0, 2.5},
		",",
		";",
	)

	if value.(string) != "0%,10%,100%,250%" {
		log.Printf("error: %v", value)
		t.FailNow()
	}

	value = TypeToString(
		[]time.Duration{0 * time.Second, 1 * time.Second, 2 * time.Second},
		",",
		";",
	)

	if value.(string) != "0s,1s,2s" {
		log.Printf("error: %v", value)
		t.FailNow()
	}

	value = TypeToString(
		[]time.Time{time.Unix(0, 0), time.Unix(50000, 0)},
		",",
		";",
	)

	if value.(string) != "1969-12-31 21:00:00 -0300 -03,1970-01-01 10:53:20 -0300 -03" {
		log.Printf("error: %v", value)
		t.FailNow()
	}

	value = TypeToString(
		[]color.RGBA{{R: 50, G: 100, B: 200, A: 255}, {R: 100, G: 200, B: 50, A: 128}},
		",",
		";",
	)

	if value.(string) != "rgba(50,100,200,1),rgba(100,200,50,0.50196)" {
		log.Printf("error: %v", value)
		t.FailNow()
	}

	value = TypeToString(
		[][]float64{{0.0, 1.0, 2.0}, {3.0, 4.0, 5.0}},
		",",
		";",
	)

	if value.(string) != "0,1,2;3,4,5" {
		log.Printf("error: %v", value)
		t.FailNow()
	}

	value = TypeToString(
		[][]float32{{0.0, 0.1, 0.2}, {0.05, 0.025, 2.0}},
		",",
		";",
	)

	if value.(string) != "0%,10%,20%;5%,2.5%,200%" {
		log.Printf("error: %v", value)
		t.FailNow()
	}

	value = TypeToString(
		[][]time.Duration{{0 * time.Second, 1 * time.Second}, {0 * time.Hour, 1 * time.Hour, 2 * time.Minute}},
		",",
		";",
	)

	if value.(string) != "0s,1s;0s,1h0m0s,2m0s" {
		log.Printf("error: %v", value)
		t.FailNow()
	}

	value = TypeToString(
		[][]time.Time{{time.Date(2022, time.October, 19, 0, 0, 0, 0, time.UTC)}, {time.Date(2020, time.October, 19, 0, 0, 0, 0, time.UTC)}},
		",",
		";",
	)

	if value.(string) != "2022-10-19 00:00:00 +0000 UTC;2020-10-19 00:00:00 +0000 UTC" {
		log.Printf("error: %v", value)
		t.FailNow()
	}

	value = TypeToString(
		[][]color.RGBA{{{R: 50, G: 100, B: 200, A: 255}, {R: 100, G: 200, B: 50, A: 128}}, {{R: 50, G: 100, B: 200, A: 255}, {R: 100, G: 200, B: 50, A: 128}}},
		",",
		";",
	)

	if value.(string) != "rgba(50,100,200,1),rgba(100,200,50,0.50196);rgba(50,100,200,1),rgba(100,200,50,0.50196)" {
		log.Printf("error: %v", value)
		t.FailNow()
	}

	value = TypeToString(
		"black",
		",",
		";",
	)

	if value.(string) != "black" {
		log.Printf("error: %v", value)
		t.FailNow()
	}

}
