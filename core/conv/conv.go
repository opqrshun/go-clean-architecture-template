package conv

import (
	"math"
	"strconv"
)

func ToInterfaceSliceFromFloatSlice(s1 []float64) (s2 []interface{}) {
	s2 = make([]interface{}, len(s1))
	for i, v := range s1 {
		s2[i] = v
	}
	return
}

//ParseToFloatList  Parse To Float64 List
func ParseToFloatList(s []string) (s2 []float64, err error) {
	var v2 float64
	for _, v := range s {
		v2, err = strconv.ParseFloat(v, 64)
		if err != nil {
			break
		}
		s2 = append(s2, v2)
	}
	return
}

//RoundRange round to 6 decimal places
func RoundRange(v float64) float64 {
	return math.Round(v*1000000) / 1000000
}

func Contains(l []string, e string) bool {
	for _, v := range l {
		if e == v {
			return true
		}
	}
	return false
}

//SliceHeadString
func SliceHeadString(s string, h string) (string, bool) {
	sl := len(s)
	hl := len(h)
	if sl <= hl {
		return "", false
	}

	r := s[hl:]
	return r, true
}

// func (repo *Parent) ToInterfaceSliceFromIntSlice(s1 []int) (s2 []interface{}) {
//   s2 = make([]interface{}, len(s1))
//   for i, v := range s1 {
//      s2[i] = v
//   }
//   return
// }

// func (repo *Parent) ParseToIntList(s []string) (s2 []int,err error) {
//   var v2 float64
// 	for _,v := range s {
//     v2,err=strconv.Atoi(v)
//     if err != nil {
//       break
//     }
// 		s2 = append(s2,v2)
// 	}
//   return
// }
