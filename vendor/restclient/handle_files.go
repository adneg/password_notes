package restclient

import (
	"bytes"
	//"crypto/tls"
	//	"crypto/x509"
	//"flag"
	//"fmt"
	//"io"
	"bufio"
	//	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func Wget(url string, filepath string) (odp string) {
	return GetContentAndSave(url, filepath)
}

func PutFile(url string, strdata string) (odp string) {
	client.Timeout = time.Duration(5 * time.Hour)
	//return "nie gotowa funkcja"
	//return SendContentJson("POST", url, strdata)
	var jsonStr = []byte(strdata)
	req, err := http.NewRequest("POST", Adres+url, bytes.NewBuffer(jsonStr))
	//req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json ; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// Dump response
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func PutFile2(url string, filepath string) (odp string) {
	client.Timeout = time.Duration(5 * time.Hour)
	//return "nie gotowa funkcja"
	//return SendContentJson("POST", url, strdata)
	//bytes.NewBuffer()
	f, _ := os.Open(filepath)

	//var jsonStr = []byte(strdata)
	req, err := http.NewRequest("POST", Adres+url, bufio.NewReader(f))
	//req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json ; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// Dump response
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func GetContentAndSave(url string, filepath string) (odp string) {

	resp, err := client.Get(Adres + url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	// Dump response
	//	data, err := ioutil.ReadAll(resp.Body)
	//	if err != nil {
	//		log.Fatal(err)
	//	}

	out, err := os.Create(filepath)
	defer out.Close()
	if err != nil {
		return "Nie Pobrano: " + filepath + " " + err.Error()
	}
	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadBytes('\n')
		out.Write(line)
		if err != nil {
			if err.Error() == "EOF" {
				return "Pobrano: " + filepath
			} else {
				return "Nie Pobrano: " + err.Error()
			}
		}
	}
}
