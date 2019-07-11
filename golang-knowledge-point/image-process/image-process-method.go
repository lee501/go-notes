package image_process

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"io/ioutil"
	"os"
)

/*
	图片文件的读写
	图片在go缓存中与base64相互转换
	图片裁剪
*/

//base64 -> file
func Base64ToFile(data string) {
	//读取string数据到缓存中
	buf, _ := base64.StdEncoding.DecodeString(data)
	//直接写入到jpg文件中
	ioutil.WriteFile("./output.jpg", buf, 0666)
}

//base64 -> buffer -> imageBuf -> file
//图片裁剪
func Base64ToBuffer(datasource string) {
	//base64 -> buffer
	data, _ := base64.StdEncoding.DecodeString(datasource) //图片文件并把文件写入到buffer
	buffer := bytes.NewBuffer(data) 	//创建buffer

	//图片裁剪
	img, _, _ := image.Decode(buffer) //传入io.Reader类型， 图片文件解码
	rgbimg := img.(*image.YCbCr)
	subimg := rgbimg.SubImage(image.Rect(0,0,200,200)).(*image.YCbCr) //图片裁剪x0 y0 x1 y1

	//img -> file
	f, _ := os.Create("test.jpg")     //创建文件
	defer f.Close()                   //关闭文件
	jpeg.Encode(f, subimg, nil)       //写入文件

	//img -> base64
	emptybuf := bytes.NewBuffer(nil) //开辟空的buff
	jpeg.Encode(emptybuf, subimg, nil) //裁剪后img写入到buff
	dist := make([]byte, 50000) //开辟存储空间
	base64.StdEncoding.Encode(dist, emptybuf.Bytes()) //buff转成base64
	ioutil.WriteFile("./base64.pic", dist, 0666)
}

//img -> base64
func ImgToBase64(imgname string) {
	img, _ := ioutil.ReadFile(imgname)
	bufstore := make([]byte, 5000000)                     //数据缓存
	base64.StdEncoding.Encode(bufstore, img)               // 文件转base64
	ioutil.WriteFile("output.jpg", bufstore, 0666)
}
