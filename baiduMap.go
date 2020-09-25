package baiduMap


const (

	defaultAppKey string = ""

	reqURLForGEO string = "http://api.map.baidu.com/reverse_geocoding/v3/?ak="

	reqURLForRoute string = "http://api.map.baidu.com/directionlite/v1/driving?coord_type=wgs84&ret_coordtype=gcj02&"
)

func GetAddressViaGEO(lat, lng string) (*StructGEOToAddress, error) {
	bc := NewBaiduMapClient(GetDefaultAK())
	return bc.GetAddressViaGEO(lat, lng)
}

func GetRoute(lat1, lng1, lat2, lng2 string) (*StructRoute, error) {
	bc := NewBaiduMapClient(GetDefaultAK())
	return bc.GetRoute(lat1, lng1, lat2, lng2)
}
