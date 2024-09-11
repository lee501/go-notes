package main

/*
#include <errno.h>

static int div(int a, int b) {
    if(b == 0) {
        errno = EINVAL;
        return 0;
    }
    return a/b;
}
*/
import "C"
import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	v0, err0 := C.div(2, 1)
	fmt.Println(v0, err0)

	v1, err1 := C.div(1, 0)
	fmt.Println(v1, err1)

	target := `
	{"Headers":{"Connection":["keep-alive"],"Content-Length":["2648"],"Content-Type":["application/json"],"Date":["Fri, 17 May 2024 02:43:35 GMT"],"Server":["nginx"],"Vary":["Accept-Encoding"],"X-Powered-By":["qianji_apisix"]},"Status":200,"Body":{"code":200,"data":[{"children":[{"children":[],"label":"七里河营销VXLAN网络"},{"children":[],"label":"七里河侧数据中心"},{"children":[],"label":"七里河实验楼数据中心"},{"children":[],"label":"实验楼用电信息采集"},{"children":[],"label":"七里河营销"},{"children":[],"label":"七里河实验楼云平台"}],"label":"七里河侧数据中心"},{"children":[{"children":[],"label":"培训中心"},{"children":[],"label":"电力交易中心"},{"children":[],"label":"思极飞天"},{"children":[],"label":"同兴智能"},{"children":[],"label":"科源集团"},{"children":[],"label":"超高压公司"},{"children":[],"label":"电科院"},{"children":[],"label":"发展事业部"},{"children":[],"label":"营销事业部"},{"children":[],"label":"送变电公司"},{"children":[],"label":"计量中心"},{"children":[],"label":"物资事业部"}],"label":"直属单位"},{"children":[{"children":[],"label":"本部盐场堡数据中心"},{"children":[],"label":"电力调度"},{"children":[],"label":"本部侧数据中心"},{"children":[],"label":"本部营销采集"}],"label":"本部侧数据中心"},{"children":[{"children":[],"label":"平凉"},{"children":[],"label":"刘家峡"},{"children":[],"label":"张掖"},{"children":[],"label":"酒泉"},{"children":[],"label":"甘南"},{"children":[],"label":"天水"},{"children":[],"label":"金昌"},{"children":[],"label":"白银"},{"children":[],"label":"陇南"},{"children":[],"label":"庆阳"},{"children":[],"label":"嘉峪关"},{"children":[],"label":"兰州"},{"children":[],"label":"临夏"},{"children":[],"label":"武威"},{"children":[],"label":"定西"}],"label":"地市公司"},{"children":[{"children":[],"label":"新区传统数据中心"},{"children":[],"label":"新区内网云平台业务"},{"children":[],"label":"新区内网云平台管理"},{"children":[],"label":"新区内网运维专区"}],"label":"新区内网云"},{"children":[{"children":[],"label":"质监站"},{"children":[],"label":"电子评标基地"},{"children":[],"label":"全省72变电站（通信所）"},{"children":[],"label":"明珠集团"},{"children":[],"label":"祁连换流站5G"},{"children":[],"label":"黄河桥南百合超市"}],"label":"其他"},{"children":[{"children":[],"label":"七里河信通园区"}],"label":"七里河园区办公网"},{"children":[{"children":[],"label":"兰州新区用电信息采集2.0网络"},{"children":[],"label":"新区计量中心"},{"children":[],"label":"新区电力调度"}],"label":"兰州新区内网"},{"children":[{"children":[],"label":"省公司园区"}],"label":"本部园区办公网"}],"message":"请求成功"}}
	`
	re := regexp.MustCompile("(七里河)")
	results := re.FindAllStringSubmatch(target, -1)

	maskMap := make(map[string]string)

	if len(results) > 0 {
		unique := make(map[string]struct{})
		for _, res := range results {
			if len(res) < 1 {
				continue
			}
			var result string
			if len(res) >= 2 {
				result = res[1]
			} else {
				result = res[0]
			}
			if _, ok := unique[result]; !ok {
				unique[result] = struct{}{}
				maskMap[result] = convertMask(result)
			}
		}
	}
	fmt.Println(maskMap)
}

func convertMask(src string) string {
	var buff bytes.Buffer
	ru := []rune(src)
	lz := len(ru)
	for i := 0; i < lz; i++ {
		if i >= (lz+1)/3 && i <= 2*(lz+1)/3 {
			buff.WriteString("*")
			continue
		}
		buff.WriteRune(ru[i])
	}
	return buff.String()
}
