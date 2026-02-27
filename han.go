package main

import (
	"fmt"
	"strconv"
)

func hanDemo() {
	//addrRe := regexp.MustCompile(`\p{Han}+省|\p{Han}+市|\p{Han}+州|\p{Han}+区|\p{Han}+县`)
	//str := `河北省廊坊市三河市`
	//re := addrRe.FindAllString(str, 5)
	//for _, item := range re {
	//	fmt.Println(item)
	//}
	//fmt.Println(len(re))
	//
	////军官
	//officeRe := regexp.MustCompile("([\u4E00-\u9FA5])(字第)([0-9a-zA-Z]{4,8})(号?)")
	//office := `2001988`
	//res := officeRe.FindAllString(office, -1)
	//fmt.Println(res)
	////护照
	//passPortCardRe := regexp.MustCompile(`1[45][0-9]{7}|([P|p|S|s]\d{7})|([S|s|G|g]\d{8})|([Gg|Tt|Ss|Ll|Qq|Dd|Aa|Ff]\d{8})|([H|h|M|m]\d{8，10})`)
	//passPortCard := "G12345678"
	//fmt.Println(passPortCardRe.FindString(passPortCard))
	//
	//a := []int{1, 2, 3, 4}
	//b := a
	//fmt.Println(a, b)
	//s := "北京市昌平区北七家镇卓悦路11号院1号楼"
	//if strings.HasSuffix(s, "院") || strings.HasSuffix(s, "楼") {
	//	fmt.Println("1")
	//}

	//m := fmt.Sprintf("%%%s%%", "你好")
	//fmt.Println(m)
	//
	//var l *string
	//l = nil
	//if *l == "" {
	//	fmt.Println("l is blank")
	//}
	//arr := [][]int{{1, 2}, {3, 4}, {5, 6}, {20, 15}}
	//fmt.Println(processGroup(arr))

	if true != false {
		fmt.Println(1)
	}
	fmt.Println(strconv.FormatBool(true))
}

func processGroup(indexes [][]int) [][]int {
	var result [][][]int
	var res [][]int
	var last int
	for i, val := range indexes {
		if i != 0 && val[0]-last > 2 {
			result = append(result, res)
			res = make([][]int, 0)
		}
		res = append(res, val)
		if i == len(indexes)-1 && len(res) > 0 {
			result = append(result, res)
		}
		last = val[1]
	}
	//获取最大的组
	max := 0
	for i, group := range result {
		if len(group) > len(result[max]) {
			max = i
		}
	}
	return result[max]
}
