package baiduMap

// get the route
type StructRoute struct {
	Status int64 `json:"status"`
	Message string `json:"message"`
	Result *StructResult `json:"result"`
}

type StructResult struct {
	Origin struct {
		Lng float64 `json:"lng"`
		Lat float64 `json:"lat"`
	} `json:"orign"`
	Destination struct {
		Lng float64 `json:"lng"`
		Lat float64 `json:"lat"`
	} `json:"destination"`
	Routes []*StructRoutes `json:"routes"`
}

type StructRoutes struct {
		Distance float64 `json:"distance"`
		Duration float64 `json:"duration"`
		Toll int64 `json:"toll"`
		Traffic_condition int64 `json:"traffic_condition"`
		Steps []*StructSteps `json:"steps"`
}

type StructSteps struct {
	Leg_index int64 `json:"leg_index"`
	Direction int64 `json:"direction"`
	Turn int64 `json:"turn"`
	Distance float64 `json:"distance"`
	Duration float64 `json:"duration"`
	Road_types string `json:"road_types"`
	Instruction string `json:"instruction"`
	Start_loaction struct {
		Lng string `json:"lng"`
		Lat string `json:"lat"`
	} `json:"start_location"`
	End_location struct {
		Lng string `json:"lng"`
		Lat string `json:"lat"`
	} `json:"end_location"`
	Path string `json:"path"`
	Traffic_condition []*StructTraffic `json:"traffic_condition"`
}

type StructTraffic struct {
	Status int64 `json:"status"`
	Geo_cnt int64 `json:"geo_cnt"`
}

// StructGEOToAddress
type StructGEOToAddress struct {
	Status int64 `json:"status"`
	Result struct {
		AddressComponent struct {
			Country  string `json:"country_code_iso2"`
		} `json:"addressComponent"`
	} `json:"result"`
}
