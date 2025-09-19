package main

import (
	"math"
)

const earthRadius = 6371.0 

type Point struct {
	Lat float64
	Lon float64
}

func NewPoint(lat, lon float64) *Point {
	return &Point{
		Lat: lat,
		Lon: lon,
	}
}

func (p *Point) Distance(other *Point) float64 {
	lat1Rad := p.toRadians(p.Lat)
	lon1Rad := p.toRadians(p.Lon)
	lat2Rad := p.toRadians(other.Lat)
	lon2Rad := p.toRadians(other.Lon)
	dLat := lat2Rad - lat1Rad
	dLon := lon2Rad - lon1Rad
	chordLengthSquared := math.Pow(math.Sin(dLat/2), 2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Pow(math.Sin(dLon/2), 2)
	centralAngle := 2 * math.Atan2(math.Sqrt(chordLengthSquared), math.Sqrt(1-chordLengthSquared))

	return earthRadius * centralAngle
}

func (p *Point) toRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}