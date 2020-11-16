package main

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	{
		// GBK TO UTF8
		gbkBytes, err := ioutil.ReadFile(os.ExpandEnv("$HOME/gbk.in"))
		if err != nil {
			log.Fatal(err)
		}
		gbkReader := bytes.NewReader(gbkBytes)
		utf8Reader := transform.NewReader(gbkReader, simplifiedchinese.GBK.NewDecoder())
		utf8Bytes, err := ioutil.ReadAll(utf8Reader)
		if err != nil {
			log.Fatal(err)
		}
		if err = ioutil.WriteFile(os.ExpandEnv("$HOME/utf8.out"), utf8Bytes, 0600); err != nil {
			log.Fatal(err)
		}
	}

	{
		// UTF8 TO GBK
		utf8Bytes, err := ioutil.ReadFile(os.ExpandEnv("$HOME/utf8.in"))
		if err != nil {
			log.Fatal(err)
		}
		utf8Reader := bytes.NewReader(utf8Bytes)
		gbkReader := transform.NewReader(utf8Reader, simplifiedchinese.GBK.NewEncoder())
		gbkBytes, err := ioutil.ReadAll(gbkReader)
		if err != nil {
			log.Fatal(err)
		}
		if err = ioutil.WriteFile(os.ExpandEnv("$HOME/gbk.out"), gbkBytes, 0600); err != nil {
			log.Fatal(err)
		}
	}
}