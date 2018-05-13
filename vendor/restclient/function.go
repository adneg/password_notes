package restclient

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	//"flag"
	//"context"
	//"fmt"
	//"io"
	//"bufio"
	//"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"os"
	"encoding/json"
	"mstr"
	"time"
)

func CreateClientTLS(ca, key, crt string) {
	cert, err := tls.LoadX509KeyPair(crt, key)
	if err != nil {
		log.Fatal(err)
	}

	// Load CA cert
	caCert, err := ioutil.ReadFile(ca)
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		RootCAs:      caCertPool,
	}

	tlsConfig.BuildNameToCertificate()
	transport = &http.Transport{
		TLSClientConfig:     tlsConfig,
		TLSHandshakeTimeout: 0 * time.Second,
		MaxIdleConnsPerHost: 1024,
		//DisableKeepAlives:   true,
		//DisableCompression:  true,
	}

	client = &http.Client{Transport: transport, Timeout: time.Duration(5 * time.Second)}

}

func SendContentJson(typ string, url string, strdata string) (odp string, err error) {
	var jsonStr = []byte(strdata)
	var data = []byte{}
	req, err := http.NewRequest(typ, Adres+url, bytes.NewBuffer(jsonStr))
	if err != nil {

		//log.Fatal(err)
		return string(data), err
	}
	req.Header.Set("Content-Type", "application/json ; charset=utf-8")
	if len(mstr.AktywnaSesja.Sessiontopgoid) > 0 {
		req.Header.Add("Sessiontopgoid", mstr.AktywnaSesja.Sessiontopgoid)
	}
	resp, err := client.Do(req)
	//defer transport.CancelRequest(req)

	if err != nil {
		//fmt.print(odp)
		//log.Fatal(err)
		return string(data), err
	}
	defer resp.Body.Close()

	// Dump response

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		//log.Fatal(err)
		return string(data), err
	}

	//fmt.Println(resp.Close)
	return string(data), err
}

func GetContent(typ string, url string) (odp []byte, err error) {
	var jsonStr = []byte("")
	var data = []byte{}
	req, err := http.NewRequest(typ, Adres+url, bytes.NewBuffer(jsonStr))
	if err != nil {

		//log.Fatal(err)
		return odp, err
	}
	//req.Header.Set("Content-Type", "application/json ; charset=utf-8")
	if len(mstr.AktywnaSesja.Sessiontopgoid) > 0 {
		req.Header.Add("Sessiontopgoid", mstr.AktywnaSesja.Sessiontopgoid)
	}
	resp, err := client.Do(req)
	//defer transport.CancelRequest(req)

	if err != nil {
		//log.Fatal(err)
		return odp, err
	}
	defer resp.Body.Close()

	// Dump response

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		//log.Fatal(err)
		return odp, err
	}

	//fmt.Println(resp.Close)

	return data, err
}
func GetContent2(url string) (odp string) {
	resp, err := client.Get(Adres + url)
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

func Logon(url string, strdata string) (odp string, err error) {
	client.Timeout = time.Duration(5 * time.Second)
	return SendContentJson("POST", url, strdata)
}

func ZmienHaslo(url string, strdata string) (odp string, err error) {
	client.Timeout = time.Duration(5 * time.Second)
	return SendContentJson("POST", url, strdata)
}

func DodajRecord(url string, strdata string) (odp string, err error) {
	client.Timeout = time.Duration(5 * time.Second)
	return SendContentJson("POST", url, strdata)
}

func PobierzRekordy(url string) (odp []mstr.PasswordRecord, err error) {
	client.Timeout = time.Duration(5 * time.Second)
	daneSJ, err := GetContent("GET", url)
	//odp = []mstr.Record{}
	if err != nil {
		return
	}
	err = json.Unmarshal(daneSJ, &odp)
	return
}

func CheckVersion(url string) (odp []byte, err error) {
	return GetContent("GET", url)
}
