package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"

	xj "github.com/basgys/goxml2json"

	log "github.com/sirupsen/logrus"
)

func decodeDemo() {
	payload := "MmQyY2U0OGY0NjhjMjdkZWlQL7Y5h2UQueQP4YX95/vXtozl1fxedITwF/M6akeLEpvt6pYjw7fo\nKLGh/niKcBnPv4KG6vN5MSeD5ykX08amuYNLB137JIxT4GkFIQz1psfsRl4sxkLqzBJ36+yurCQm\nmPzony8fyBzKzBFYr9dfiB/kq0cOgsoDw2VFMZOh7Ee4xZzseJQLeOjUabPfmoGRtXaGKLq+60Cr\nVIRVvJ/yfHx0LUHhe33YyHKmj/4H+Yg8nZf7LH5e2wxJ80tQghuwSAEtDpfxsXdrGiv8zwS6jdaf\nuGP97DiRyvSxRuX1TuQ981mxgvoJ65hTYl2BCLtMCxMPiILn9EjxKz3ySxZUv17ENIb77oxqGW+J\npr8MuYQRr/pG0IOu4ReDXysOxHGCKaIKeH4AmzAD4EKmZ3YswisTMLASgqtPpJg+N7VkDqB0Mt8v\n97+hFkYI4G+pQEAw9DNsZb5JPTgEA81VlyWt9rX7AEIHvPVsXTyg4hONyMaIbh0cO3muL905Y91H\n+iBjFHh12zTHyTVQjnj67uFQCl5uGzVGcq25iF1+d1vBln4zYAeP9ncFKGz+ZP7IyoiLbIax6LoI\nULWxsliOB86L3oIYrPn5gztYifmcySA90TyrI2x0GLFT2IAxPCG8okdl2BA7rtzGpaAjNf+79rmC\nl7dGWp39FfAPbAdbVl7ID8vE0a4ZjtsCNtKIMZpg3Z+hGNgVlreceMb3o6zGDW27lj6GuDf1toxU\nUxVDsmCsxuy/EiTKf3G6RVoSfC4ZbSPjMUyIwy7YH09SrC1KgCUOqB688l7PMsGbQPnQv43mpYxN\nNiX2KnT/BhdqmbY8GcxdryZB/8mlqDyhVmPToIb7RpiqxbH71LY3QLiHYQskokfpJlRRW0G5dJr/\nx7xvNw==\n"
	res := decrypt([]byte(payload), "11111")

	if bytes.HasPrefix([]byte(payload), []byte("<")) {
		fmt.Println("2222")
		return
	}
	json, err := Xml2Json(res)
	if err != nil {
		fmt.Println(err)
	}
	toMap := JsonToMap(json)
	fmt.Printf("-------------%#v\n", toMap)
}

func JsonToMap(str []byte) (m map[string]interface{}) {
	_ = json.Unmarshal(str, &m)
	return
}

func Xml2Json(xml []byte) (json []byte, err error) {
	convert, err := xj.Convert(bytes.NewReader(xml))
	if err != nil {
		return json, err
	}
	return convert.Bytes(), err
}

func decrypt(data []byte, traceID string) (res []byte) {
	decode := Base64Decode(data)
	// fix array out of bound issue
	if len(decode) <= 16 {
		return []byte("")
	}
	decrypt, err := AesECBDecrypt(decode[16:], decode[0:16])
	if err != nil {
		log.WithFields(log.Fields{
			"trace_id": traceID,
			"data":     data,
			"err":      err,
		}).Error("decrypt failed.")
		return []byte("")
	}
	return decrypt
}

func Base64Decode(data []byte) []byte {
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	n, err := base64.StdEncoding.Decode(dst, data)
	if err != nil {
		log.Error("base64 decode string failed.", err)
		return nil
	}
	return dst[:n]
}

func AesECBDecrypt(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}
	ecb := NewECBDecryptEr(block)
	retData := make([]byte, len(data))
	ecb.CryptBlocks(retData, data)
	// 解PKCS7填充
	retData = PKCS7UnPadding(retData)
	return retData, nil
}

func PKCS7UnPadding(encrypt []byte) []byte {
	length := len(encrypt)
	unPadding := int(encrypt[length-1])
	return encrypt[:(length - unPadding)]
}

type ecbDecryptEr ecb

type ecb struct {
	b         cipher.Block
	blockSize int
}

func (x *ecbDecryptEr) BlockSize() int { return x.blockSize }
func (x *ecbDecryptEr) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

func NewECBDecryptEr(b cipher.Block) cipher.BlockMode {
	return (*ecbDecryptEr)(newECB(b))
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}
