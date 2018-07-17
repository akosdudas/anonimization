package anonbll


import (
	"fmt"
	"anondb"
	"anonmodel"
	"math"
)

type GPSDimension struct {
	anonCollectionName string
	fieldName          string
	originalArea      anonmodel.GPSArea
	currentArea       anonmodel.GPSArea
}

func (d *GPSDimension) initialize(anonCollectionName string, fieldName string) {
	d.anonCollectionName = anonCollectionName
	d.fieldName = fieldName

}

func (d *GPSDimension) getInitialBoundaries() anonmodel.Boundary {
	return &anonmodel.GPSBoundary{
		Latitude: 	numericDimension.getInitialBoundaries()
		Longitude:	numericDimension.getInitialBoundaries()
		isGlobal: 	true
	}
}

func (d *GPSDimension)getZeroCut(){
	//Lat
	//find max diff ->minden alatta +360
}

func (d *GPSDimension) getDimensionForStatistics(stat interface{}, firstRun bool) mondrianDimension {
	Latitude.getDimensionForStatistics().
	Longitude.getDimensionForStatistics()
}



func (d *GPSDimension) tryGetAllowableCut(k int, partition anonmodel.Partition, count int) (bool, anonmodel.Partition, anonmodel.Partition, error) {
	LatRange = d.currentArea.Latitude.max-d.currentArea.Latitude.min
	LonRange = d.currentArea.Longitude.max-d.currentArea.Longitude.min
	var Arr [2]string
	if LonRange> LatRange{
		Arr[0] ="Longitude"
		Arr[1] ="Latitude"  
	}
	else{
		Arr[0] ="Latitude"
		Arr[1] ="Longitude"  
	}
	for _, LL:=Arr{
		bool ok;
		ok, anonmodel.Partition, anonmodel.Partition, error = getCut(k, partition, "LL", count)
		if(ok&&(err==nil))
			return
	}	
	return
}

func (d *GPSDimension) getCut(k int, partition anonmodel.Partition, LL string, count int) (bool, anonmodel.Partition, anonmodel.Partition, error) {
	
}


func (d *GPSDimension) prepare(partition anonmodel.Partition, count int) {

}