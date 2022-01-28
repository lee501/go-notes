package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/url"

	xj "github.com/basgys/goxml2json"
)

var (
	 desKey = []byte("shenzhen")
	 phoneFiled = []byte("%22contactMobile%22")
	 nameFiled = []byte("%22contactName%22")
)
func parserB2CDESPayload(payload []byte, key []byte, field []byte) (string, error) {
	res := ""
	specialStr := []byte("%22")
	var encodeStr []byte
	start := bytes.Index(payload, field)
	if start > -1 {
		strLen := len(payload)
		start += len(field)
		for start <= (strLen - 3) {
			if bytes.Equal(payload[start:start+3], specialStr) {
				start += 3
				break
			}
			start++
		}

		end := start
		for end <= (strLen - 3) {
			if bytes.Equal(payload[end:end+3], specialStr) {
				break
			}
			end++
		}
		if end <= (strLen - 3) {
			encodeStr = payload[start:end]
		}
	}
	if len(encodeStr) == 0 {
		return res, errors.New("payload is not have ")
	}
	var err error
	res, err = DecryptDES_ECB(encodeStr, key)
	if err != nil {
		return res, err
	}
	return res, nil
}

func DecryptDES_ECB(encodeStr []byte, key []byte) (string, error) {
	decodeStr, ok := url.QueryUnescape(string(encodeStr))
	if ok != nil {
		return "", errors.New("urlEncode illegal")
	}
	data, _ := base64.StdEncoding.DecodeString(decodeStr)
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()
	if len(data)%bs != 0 {
		return "", errors.New("crypto/cipher: input not full blocks")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Decrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	out = PKCS5UnPadding(out)
	return string(out), nil
}

// PKCS5UnPadding Plaintext subtraction
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func JsonToMap(str []byte) (m map[string]interface{}) {
	err := json.Unmarshal(str, &m)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func Xml2Json(xml []byte) (json []byte, err error) {
	convert, err := xj.Convert(bytes.NewReader(xml))
	if err != nil {
		fmt.Println("111111111", err)
		return json, err
	}
	fmt.Println("--------", convert.Bytes())
	return convert.Bytes(), err
}

func Base64Decode(data []byte) []byte {
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	n, err := base64.StdEncoding.Decode(dst, data)
	if err != nil {
		log.Println("base64 decode string failed.", err)
		return nil
	}
	return dst[:n]
}

// PKCS7加填充/和PKCS5填充一样,只是填充字段多少的区别
func PKCS7Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

// PKCS7解填充/和PKCS5填充一样,只是填充字段多少的区别
func PKCS7UnPadding(encrypt []byte) []byte {
	length := len(encrypt)
	unPadding := int(encrypt[length-1])
	return encrypt[:(length - unPadding)]
}

// AES/ECB/PKCS7模式加密--签名加密方式
func AesECBEncrypt(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}
	ecb := NewECBEncryptEr(block)
	// 加PKCS7填充
	content := PKCS7Padding(data, block.BlockSize())
	encryptData := make([]byte, len(content))
	// 生成加密数据
	ecb.CryptBlocks(encryptData, content)
	return encryptData, nil
}

// AES/ECB/PKCS7模式解密--签名解密方式
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

func AesCBCDecrypt(data, key, iv []byte) ([]byte, error) {
	defer func() {
		if r := recover(); r != nil {

		}
	}()
	var block cipher.Block
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}
	cbc := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(data))
	cbc.CryptBlocks(plaintext, data)
	return PKCS7UnPadding(plaintext), nil
}

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbEncryptEr ecb

func NewECBEncryptEr(b cipher.Block) cipher.BlockMode {
	return (*ecbEncryptEr)(newECB(b))
}
func (x *ecbEncryptEr) BlockSize() int { return x.blockSize }
func (x *ecbEncryptEr) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

// ecb解密方法
type ecbDecryptEr ecb

func NewECBDecryptEr(b cipher.Block) cipher.BlockMode {
	return (*ecbDecryptEr)(newECB(b))
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

func main()  {
	payload := "passengerJson=%7B%22passengerList%22%3A%5B%7B%22id%22%3A%22passenger0%22%2C%22psgrType%22%3A%220%22%2C%22psgrName%22%3A%2245Sf2zWHsqN0P26wa9mdCw%3D%3D%22%2C%22certType%22%3A%220%22%2C%22certNo%22%3A%22IXGBdeZZz656Ax5XJOgHzwiHcAQQmfOw%22%2C%22psgrPhone%22%3A%22oE95FLaPzS79PoHeu1dbHA%3D%3D%22%2C%22psgrBirth%22%3A%22gxXFbrT7QTGecgByXQJaWw%3D%3D%22%2C%22psgrAge%22%3A30%2C%22products%22%3A%5B%5D%7D%5D%2C%22contactName%22%3A%22ntIqkSjWMvs%3D%22%2C%22contactMobile%22%3A%22oE95FLaPzS79PoHeu1dbHA%3D%3D%22%2C%22contactTelphone%22%3A%22gmuEq2Cv4kw%3D%22%2C%22contactEmail%22%3A%22gmuEq2Cv4kw%3D%22%2C%22isMemberContact%22%3Atrue%2C%22lxrId%22%3A%221103832%22%2C%22isDefault%22%3A%220%22%2C%22totalPrice%22%3A2370%2C%22onlyChild%22%3Afalse%2C%22selectNoticeType%22%3Anull%2C%22oriContactEm%22%3A%22%22%7D&postMsg=%7B%22orderPostType%22%3A%221%22%7D"
	concatPhone, _ := parserB2CDESPayload([]byte(payload), desKey, phoneFiled)
	contactName, _ := parserB2CDESPayload([]byte(payload), desKey, nameFiled)
	fmt.Println(concatPhone, contactName)
	json := `%7B%22passengerList%22%3A%5B%7B%22id%22%3A%22passenger0%22%2C%22psgrType%22%3A%220%22%2C%22psgrName%22%3A%22Q379VANxwaA%3D%22%2C%22certType%22%3A%220%22%2C%22certNo%22%3A%22iTDxj6kO8Hr7xTBtA77Ec2MaRCOSXsfS%22%2C%22psgrPhone%22%3A%22NfVPnowD9wtWsbtxRVNUsA%3D%3D%22%2C%22psgrBirth%22%3A%22u%2F3DPguinZdKg7qMCmESsQ%3D%3D%22%2C%22psgrAge%22%3A29%2C%22products%22%3A%5B%7B%22productName%22%3A%22couponUse%22%2C%22productInstanceId%22%3A%224C75F70DA73C%22%2C%22productId%22%3A%22HAKYHQ2022012502%22%2C%22properties%22%3A%5B%7B%22propertyCode%22%3A%22couponType%22%2C%22propertyType%22%3A%223%22%2C%22propertyValue%22%3A%221%22%7D%2C%7B%22propertyCode%22%3A%22couponValue%22%2C%22propertyType%22%3A%221%22%2C%22propertyValue%22%3A%2260%22%7D%5D%7D%5D%7D%2C%7B%22id%22%3A%22passenger1%22%2C%22psgrType%22%3A%222%22%2C%22psgrName%22%3A%22V6QPy3yLP0z9FrSAcd%2Felw%3D%3D%22%2C%22certType%22%3A%229%22%2C%22certNo%22%3A%220F396IM%2B38jpmZYpMTghTw%3D%3D%22%2C%22psgrPhone%22%3A%22HYgBDcA7FD0Ji%2Fz3ttVf8Q%3D%3D%22%2C%22psgrBirth%22%3A%22QuxdtPeCdzEuTEq30EZv2A%3D%3D%22%2C%22psgrAge%22%3A0%2C%22boundPsgId%22%3A%22passenger0%22%2C%22products%22%3A%5B%5D%7D%5D%2C%22contactName%22%3A%22Q379VANxwaA%3D%22%2C%22contactMobile%22%3A%22HYgBDcA7FD0Ji%2Fz3ttVf8Q%3D%3D%22%2C%22contactTelphone%22%3A%22gmuEq2Cv4kw%3D%22%2C%22contactEmail%22%3A%22gmuEq2Cv4kw%3D%22%2C%22isMemberContact%22%3Atrue%2C%22totalPrice%22%3A880%2C%22onlyChild%22%3Afalse%2C%22selectNoticeType%22%3A1%2C%22oriContactEm%22%3A%22%22%7D&postMsg=%7B%22orderPostType%22%3A%221%22%7D`
	//JsonToMap([]byte(json))
	res := decrypt([]byte(json), "")
	Xml2Json(res)
}