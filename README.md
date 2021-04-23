# 中国省、市、县级经纬度列表

在使用 EChar 老版本的时候，国内的地理位置补完整，网上给的都不全，就整理了下。

Release 有整理好的。

## 解析

```go
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
```

## Update

* 2021年4月21日 -- 首次，[参考](http://www.mca.gov.cn/article/sj/xzqh/2020/2020/202101041104.html)

## 依赖

* [2020年行政区划代码](http://www.mca.gov.cn/article/sj/xzqh/2020/)
* Google Geocoding
* [百度地图 -- 拾取坐标系统](https://api.map.baidu.com/lbsapi/getpoint/index.html)



