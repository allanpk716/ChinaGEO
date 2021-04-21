package main

import (
	"encoding/json"
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/gocarina/gocsv"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

// LocationInfo 解析出 CSV 的数据内容
type LocationInfo struct {
	AreaId	string `csv:"areaid"`
	Name	string `csv:"name"`
}

var hzRegexp = regexp.MustCompile("^[\u4e00-\u9fa5]$")
// StrFilterNonChinese 去除特殊字符
func StrFilterNonChinese(src *string) {
	strn := ""
	for _, c := range *src {
		if hzRegexp.MatchString(string(c)) {
			strn += string(c)
		}
	}
	*src = strn
}

func main() {

	csvFile, err := os.OpenFile("geo.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()
	// 把原来ANSI格式的文本文件里的字符，用gbk进行解码。
	decoder := mahonia.NewDecoder("gbk")
	csvInput := []*LocationInfo{}
	if err := gocsv.UnmarshalFile(csvFile, &csvInput); err != nil {
		panic(err)
	}

	chinaGEO := make(map[string]GEOInfo, 4000)
	for _, One := range csvInput {
		nowName := decoder.ConvertString(One.Name)
		StrFilterNonChinese(&nowName)
		// 一共是 6 位数
		// 一级区域 省级		后4为是 0000
		// 二级区域 市级		只有最后两位是 00
		// 三级区域 县级
		// 切割编号
		last4Code := One.AreaId[2:]
		tmpGeoInfo := NewGEOInfo()
		tmpGeoInfo.AreaId = One.AreaId
		tmpGeoInfo.Name = nowName
		if last4Code == "0000" {
			// 那么代表他是一级区域，无需赋值
			tmpGeoInfo.FullName = nowName
			tmpGeoInfo.PlaceType = "1"
		} else {
			// 先判断是不是二级区域
			last2COde := last4Code[2:]
			if last2COde == "00" {
				// 是二级区域
				// 那么就需要赋值他有一级区域，那么推算来说就是 前两位 后面补0000
				tmpGeoInfo.FirstId = One.AreaId[0:2] + "0000"
				tmpGeoInfo.FullName = chinaGEO[tmpGeoInfo.FirstId].Name + nowName
				tmpGeoInfo.PlaceType = "2"
			} else {
				// 是三级区域
				// 那么就需要赋值一级、二级区域ID
				tmpGeoInfo.FirstId = One.AreaId[0:2] + "0000"
				tmpGeoInfo.SecondId = One.AreaId[0:4] + "00"
				tmpGeoInfo.FullName = chinaGEO[tmpGeoInfo.FirstId].Name + chinaGEO[tmpGeoInfo.SecondId].Name + nowName
				tmpGeoInfo.PlaceType = "3"
			}
		}
		chinaGEO[One.AreaId] = *tmpGeoInfo
	}

	mason, err :=json.Marshal(chinaGEO)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("geo_no.json", mason, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Read CSV Done.")

	// 获取对应地区的经纬度
	gHelper := NewGEOHelper()
	for _, info := range chinaGEO {
		tmpResult, err := info.Query(gHelper)
		if err != nil {
			fmt.Println(info.Name + "google geo locate error")
			continue
		}
		tmpGeoInfo := chinaGEO[info.AreaId]
		strLat := strconv.FormatFloat(tmpResult.Lat, 'E', -1, 64)
		tmpGeoInfo.Lat = strLat
		strLng := strconv.FormatFloat(tmpResult.Lng, 'E', -1, 64)
		tmpGeoInfo.Lng = strLng
		chinaGEO[info.AreaId] = tmpGeoInfo
	}

	mason, err =json.Marshal(chinaGEO)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("geo_yes.json", mason, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Done.")
}