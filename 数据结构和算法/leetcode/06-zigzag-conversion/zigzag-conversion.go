package zigzag

import "bytes"

/*
	按索引规律处理
		p = numRows * 2 - 2
		第0行 0*p, 1*p
		第r行 r, 1*p -r, 2*p + r
		第n行 numRows - 1, numRows - 1 + p
*/
func ZigZagConvert(str string, numRows int) string {
	if numRows <= 1 || len(str) < numRows {
		return str
	}
	p := numRows * 2 - 2
	res := bytes.Buffer{}
	//处理第一行
	for i := 0; i < len(str); i += p {
		res.WriteByte(str[i])
	}
	//处理中间行
	for r := 1; r < numRows - 1; r++ {
		//写入每行的第一个元素
		res.WriteByte(str[r])
		for k := p; k - r < len(str); k += p {
			res.WriteByte(str[k-r])
			if k + r < len(str) {
				res.WriteByte(str[k+r])
			}
		}
	}

	//处理最后一行
	for i:=numRows -1; i<len(str); i += p {
		res.WriteByte(str[i])
	}
	return res.String()
}
