package DataStructure

import (
	"math"
	"strconv"
	"strings"
)

type Property struct {
	PId      string `json:"property_id"`
	PName    string `json:"property_name"`
	PDefault string `json:"property_default"`

	PValue string `json:"property_value"`
	PRange string `json:"property_range"` //as "x,y" or "none",use- for -max,+ for +max

	PMinSample               string `json:"property_minSample"`
	PRealSample              string `json:"property_realSample"`
	PCheckResult             string `json:"ccCheckResult"`
	PCheckRecord_builder     string `json:"property_checkRecord_builder"`
	PCheckRecord_supervisor  string `json:"property_checkRecord_supervisor"`
	PCheckRecord_constructor string `json:"property_checkRecord_constructor"`
}

func (p *Property) CheckRange() {

	var min float64 = -math.MaxFloat64
	var max float64 = math.MaxFloat64

	//if no range
	if p.PRange == "none" {
		p.PCheckResult = "合格"
		//return true
		return
	}

	nums := strings.Split(p.PRange, ",")
	if len(nums) > 0 && nums[0] != "-" {
		min, _ = strconv.ParseFloat(nums[0], 64)
	}
	if len(nums) > 1 && nums[1] != "+" {
		max, _ = strconv.ParseFloat(nums[1], 64)
	}

	//floatVal,_ :=strconv.ParseFloat(p.PValue, 64)

	valueVec := strings.Fields(p.PValue)
	for i, _ := range valueVec {
		valueVec[i] = strings.Trim(valueVec[i], p.PDefault)
	}

	result := "合格"

	for _, v := range valueVec {
		floatVal, _ := strconv.ParseFloat(v, 64)
		if floatVal < min || floatVal > max {

			result = "不合格"
		}
	}

	p.PCheckResult = result

	return
}

/*
	if floatVal>=min&&floatVal<=max{
		//p.PValue = "6"
		p.PCheckResult = "合格"
		//return true
		return
	} else{
		p.PCheckResult = "不合格"
		//return false
		return
	}

}
*/
