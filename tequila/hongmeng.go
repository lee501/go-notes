package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
)

const (
	AES128ECBPkcs5AlgName = "AES128|ECB|PKCS5"
	AES128ECBPkcs5Key     = "Wearebest!123456"
	AES128                = "AES128"
)

// CryptUtil 加密工具类

var (
	ErrInvalidPKCS7Padding = errors.New("invalid padding")
	ErrInvalidBlockSize    = errors.New("invalid block size")
)

// PKCS7加填充/和PKCS5填充一样,只是填充字段多少的区别
func PKCS7Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

// PKCS7解填充/和PKCS5填充一样,只是填充字段多少的区别
func PKCS7UnPadding(data []byte, blockSize int) ([]byte, error) {
	if len(data) == 0 {
		return nil, ErrInvalidPKCS7Padding
	}
	if len(data)%blockSize != 0 {
		return nil, ErrInvalidPKCS7Padding
	}
	c := data[len(data)-1]
	n := int(c)
	if n == 0 || n > len(data) {
		return nil, ErrInvalidPKCS7Padding
	}
	for i := 0; i < n; i++ {
		if data[len(data)-n+i] != c {
			return nil, ErrInvalidPKCS7Padding
		}
	}
	return data[:len(data)-n], nil
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
	return PKCS7UnPadding(retData, block.BlockSize())
}

func AesCBCDecrypt(data, key, iv []byte) ([]byte, error) {
	var block cipher.Block
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}

	if iv == nil {
		iv = make([]byte, block.BlockSize())
	} else if len(iv) != block.BlockSize() {
		return nil, ErrInvalidBlockSize
	}
	cbc := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(data))
	cbc.CryptBlocks(plaintext, data)
	return PKCS7UnPadding(plaintext, block.BlockSize())
}

func AesCBCEncrypt(data, key, iv []byte) ([]byte, error) {
	// 创建AES加密器
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if iv == nil {
		iv = make([]byte, block.BlockSize())
	} else if len(iv) != block.BlockSize() {
		return nil, ErrInvalidBlockSize
	}

	// 创建CBC模式的加密器
	mode := cipher.NewCBCEncrypter(block, iv)
	data = PKCS7Padding(data, mode.BlockSize())
	ciphertext := make([]byte, len(data))
	mode.CryptBlocks(ciphertext, data)
	return ciphertext, nil
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

// ExampleAESDecrypt 演示AES解密功能的示例函数
func ExampleAESDecrypt() {
	// 示例：解密
	encrypted := "gq6K18BHWVqq7m0wgW3AUw=="
	// Base64 解码
	encryptedBytes, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		log.Printf("Base64 解码失败: %v", err)
		return
	}

	decrypted, err := AesECBDecrypt(encryptedBytes, []byte(AES128ECBPkcs5Key))
	if err != nil {
		log.Printf("解密失败: %v", err)
		return
	}
	fmt.Printf("解密结果: %s\n", decrypted)
}
