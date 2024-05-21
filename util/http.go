package util

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"encoding/pem"
	"encoding/xml"
	"fmt"
	"golang.org/x/crypto/pkcs12"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

const(
	BASE_URL = "/service/XChangeServlet"
)

//HTTPGet get request
func HTTPGet(uri string) ([]byte, error) {
	response, err := http.Get(uri)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	return ioutil.ReadAll(response.Body)
}

//HTTPPost post 请求
func HTTPPost(uri string, data string) ([]byte, error) {
	body := bytes.NewBuffer([]byte(data))
	response, err := http.Post(uri, "", body)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	return ioutil.ReadAll(response.Body)
}

//PostJSON post json data request
func PostJSON(uri string, obj interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	jsonData = bytes.ReplaceAll(jsonData, []byte("\\u003c"), []byte("<"))
	jsonData = bytes.ReplaceAll(jsonData, []byte("\\u003e"), []byte(">"))
	jsonData = bytes.ReplaceAll(jsonData, []byte("\\u0026"), []byte("&"))
	body := bytes.NewBuffer(jsonData)
	response, err := http.Post(uri, "application/json;charset=utf-8", body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	return ioutil.ReadAll(response.Body)
}

// PostJSONWithRespContentType post json data request , return data type
func PostJSONWithRespContentType(uri string, obj interface{}) ([]byte, string, error) {
	jsonData, err := json.Marshal(obj)
	if err != nil {
		return nil, "", err
	}

	jsonData = bytes.ReplaceAll(jsonData, []byte("\\u003c"), []byte("<"))
	jsonData = bytes.ReplaceAll(jsonData, []byte("\\u003e"), []byte(">"))
	jsonData = bytes.ReplaceAll(jsonData, []byte("\\u0026"), []byte("&"))

	body := bytes.NewBuffer(jsonData)
	response, err := http.Post(uri, "application/json;charset=utf-8", body)
	if err != nil {
		return nil, "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	contentType := response.Header.Get("Content-Type")
	return responseData, contentType, err
}

//PostFile upload file
func PostFile(fieldname, filename, uri string) ([]byte, error) {
	fields := []MultipartFormField{
		{
			IsFile:    true,
			Fieldname: fieldname,
			Filename:  filename,
		},
	}
	return PostMultipartForm(fields, uri)
}

//MultipartFormField save file and other field
type MultipartFormField struct {
	IsFile    bool
	Fieldname string
	Value     []byte
	Filename  string
}

//PostMultipartForm update file and other field
func PostMultipartForm(fields []MultipartFormField, uri string) (respBody []byte, err error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	for _, field := range fields {
		if field.IsFile {
			fileWriter, e := bodyWriter.CreateFormFile(field.Fieldname, field.Filename)
			if e != nil {
				err = fmt.Errorf("error writing to buffer , err=%v", e)
				return
			}

			fh, e := os.Open(field.Filename)
			if e != nil {
				err = fmt.Errorf("error opening file , err=%v", e)
				return
			}
			defer fh.Close()

			if _, err = io.Copy(fileWriter, fh); err != nil {
				return
			}
		} else {
			partWriter, e := bodyWriter.CreateFormField(field.Fieldname)
			if e != nil {
				err = e
				return
			}
			valueReader := bytes.NewReader(field.Value)
			if _, err = io.Copy(partWriter, valueReader); err != nil {
				return
			}
		}
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, e := http.Post(uri, contentType, bodyBuf)
	if e != nil {
		err = e
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}
	respBody, err = ioutil.ReadAll(resp.Body)
	return
}

//PostXML perform a HTTP/POST request with XML body
func PostXML(uri string, obj interface{}) ([]byte, error) {
	xmlData, err := xml.Marshal(obj)
	if err != nil {
		return nil, err
	}

	body := bytes.NewBuffer(xmlData)
	response, err := http.Post(uri, "application/xml;charset=utf-8", body)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http code error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	return ioutil.ReadAll(response.Body)
}


//httpWithTLS CA
func httpWithTLS(rootCa, key string) (*http.Client, error) {
	var client *http.Client
	certData, err := ioutil.ReadFile(rootCa)
	if err != nil {
		return nil, fmt.Errorf("unable to find cert path=%s, error=%v", rootCa, err)
	}
	cert := pkcs12ToPem(certData, key)
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	tr := &http.Transport{
		TLSClientConfig:    config,
		DisableCompression: true,
	}
	client = &http.Client{Transport: tr}
	return client, nil
}

//pkcs12ToPem  Pkcs12 to Pem
func pkcs12ToPem(p12 []byte, password string) tls.Certificate {
	blocks, err := pkcs12.ToPEM(p12, password)
	defer func() {
		if x := recover(); x != nil {
		}
	}()
	if err != nil {
		panic(err)
	}
	var pemData []byte
	for _, b := range blocks {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}
	cert, err := tls.X509KeyPair(pemData, pemData)
	if err != nil {
		panic(err)
	}
	return cert
}

//PostXMLWithTLS perform a HTTP/POST request with XML body and TLS
func PostXMLWithTLS(uri string, obj interface{}, ca, key string) ([]byte, error) {
	xmlData, err := xml.Marshal(obj)
	if err != nil {
		return nil, err
	}

	body := bytes.NewBuffer(xmlData)
	client, err := httpWithTLS(ca, key)
	if err != nil {
		return nil, err
	}
	response, err := client.Post(uri, "application/xml;charset=utf-8", body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http code error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	return ioutil.ReadAll(response.Body)
}
