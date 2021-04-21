package main

import geo "github.com/allanpk716/google-geolocate"

type GEOInfo struct {
	PlaceType string `json:"place_type,omitempty"` // 1 是 省级，2 是市级，3 是县级
	FirstId   string `json:"first_id,omitempty"`   // 父级的一级编码
	SecondId  string `json:"second_id,omitempty"`  // 父级的二级编码
	AreaId    string `json:"area_id,omitempty"`    // 当前区域编号
	Name      string `json:"name,omitempty"`       // 直接的地区描述，无需包含父级区域名称
	FullName  string `json:"full_name,omitempty"`  // 完整的区域名称，省市县 这样
	Lat       string `json:"lat"`        // 纬度
	Lng       string `json:"lng"`        // 经度
}

func NewGEOInfo() *GEOInfo {
	g := GEOInfo{}
	return &g
}

func (g GEOInfo) Query(gHelper *GEOHelper) (*geo.Point, error) {
	return gHelper.queryLocationName(g.FullName)
}

func (g GEOInfo) ToEChartOnePlaceString() string {
	out := ""
	out = "'" + g.Name + "': " + "[" + g.Lng + ", " + g.Lat + "],\n"
	return out
}
