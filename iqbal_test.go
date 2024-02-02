package geobackend

import (
	"fmt"
	"testing"
)

var dbname = "GIS"
var collname = "GIS"

func TestGeoIntersects(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", dbname)
	coordinates := Point{
		Coordinates: []float64{
			107.56875846592652, -6.892669045544224,
		},
	}
	datagedung := GeoIntersects(mconn, collname, coordinates)
	fmt.Println(datagedung)
}

func TestGeoWithin(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", dbname)
	coordinates := Polygon{
		Coordinates: [][][]float64{
			{
				{107.57824852359408, -6.888634062096415},
				{107.57820293585849, -6.889668334903732},
				{107.57868769409816, -6.8897273355695745},
				{107.57863484658856, -6.888611511305527},
				{107.578248523594085, -6.888634062096415},
			},
		},
	}
	datagedung := GeoWithin(mconn, collname, coordinates)
	fmt.Println(datagedung)
}

func TestNear(t *testing.T) {
	mconn := SetConnection2dsphere("MONGOSTRING", "GIS", "GIS")
	coordinates := Point{
		Coordinates: []float64{
			107.5683402532045, -6.887292845609608,
		},
	}
	datagedung := Near(mconn, "GIS", coordinates)
	fmt.Println(datagedung)
}

func TestNearSphere(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "GIS")
	coordinates := Point{
		Coordinates: []float64{
			107.57380761465691, -6.889092544178368,
		},
	}
	datagedung := NearSphere(mconn, "GIS", coordinates)
	fmt.Println(datagedung)
}

func TestBox(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "GIS")
	coordinates := Polyline{
		Coordinates: [][]float64{
			{107.55949279450914, -6.889856251464579},
			{107.55835040677192, -6.8918834318207445},
		},
	}
	datagedung := Box(mconn, collname, coordinates)
	fmt.Println(datagedung)
}
