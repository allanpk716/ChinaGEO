package main

import (
	"testing"
)

func TestGEOHelper_queryLocationName(t *testing.T) {

	gHelper := NewGEOHelper()
	//result, err := gHelper.queryLocationName("安徽省蚌埠市")
	//if err != nil {
	//	t.Error(err)
	//}
	//println("Lat", result.Lat, "Lng", result.Lng)
	//result, err = gHelper.queryLocationName("新疆维吾尔自治区昌吉回族自治州呼图壁县")
	//if err != nil {
	//	t.Error(err)
	//}
	//println("Lat", result.Lat, "Lng", result.Lng)
	result, err := gHelper.queryLocationName("海南省三沙市")
	if err != nil {
		t.Error(err)
	}
	println("Lat", result.Lat, "Lng", result.Lng)
}
