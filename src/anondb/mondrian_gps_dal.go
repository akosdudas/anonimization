package anondb

import	"anonbll"


//GPSCoord represend a coordinate pair
type GPSCoord struct {
	Latitude  float64 `json:"latitude" bson:"latitude"`
	Longitude float64 `json:"longitude" bson:"longitude"`
}

func PreprocessCoord(coordStr string) (coord GPSCoord, err error){
	format:=anonbll.FindFormat(coordStr)
	coord.Latitude, coord.Longitude,  err:= anonbll.ReadCordsValue(coord, format)
	return
}