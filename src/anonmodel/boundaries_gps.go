package anonmodel

import "strconv"

//GPSBoundary represents a boundary for a coordinate type field
type GPSBoundary struct {
	Latitude  NumericBoundary
	Longitude NumericBoundary
	isGlobal  bool
}

//GPSArea  holds the range for a gps range
type GPSArea struct {
	Latitude  NumericRange
	Longitude NumericRange
}

//GetRelativeArea calculates the relative area of coords
func (r *GPSArea) GetRelativeArea() float64 {
	LatitudeDiff := r.Latitude.Max - r.Latitude.Min
	LongitudeDiff := r.Longitude.Max - r.Longitude.Min
	actualArea := LatitudeDiff * LongitudeDiff
	var globalArea float64 = 360 * 180
	return actualArea / globalArea
}

//Clone ...
func (b *GPSBoundary) Clone() Boundary {
	return &GPSBoundary{Latitude: b.Latitude,
		Longitude: b.Longitude,
		isGlobal:  b.isGlobal}
}

//GetGeneralizedValue gv aaaa
func (b *GPSBoundary) GetGeneralizedValue() string {
	result := strconv.FormatFloat(*b.Latitude.UpperBound, 'f', -1, 64) + ":" + strconv.FormatFloat(*b.Latitude.LowerBound, 'f', -1, 64) + ", "
	result += strconv.FormatFloat(*b.Longitude.UpperBound, 'f', -1, 64) + ":" + strconv.FormatFloat(*b.Longitude.LowerBound, 'f', -1, 64)
	return result
}
