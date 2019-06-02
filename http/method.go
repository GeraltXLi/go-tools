package http

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
)

func SGet(reqUrl string) ([]byte, error) {
	var res []byte
	resp, err := http.Get(reqUrl)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	res, err = ioutil.ReadAll(resp.Body)
	return res, err
}
func Get(reqUrl string, data map[string]string) ([]byte, error) {
	var res []byte
	params := url.Values{}

	//set get params
	rUrl, err := url.Parse(reqUrl)
	if err != nil {
		return res, err
	}
	for k, v := range data {
		params.Set(k, v)
	}

	//urlencode
	rUrl.RawQuery = params.Encode()
	urlPath := rUrl.String()
	req, err := http.NewRequest("GET", urlPath, nil)
	if err != nil {
		return res, err
	}

	resp, err := HttpClient.Do(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	res, err = ioutil.ReadAll(resp.Body)
	return res, err
}

func Post(url string, data []byte) ([]byte, error) {
	var res []byte
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return res, err
	}
	resp, err := HttpClient.Do(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	res, err = ioutil.ReadAll(resp.Body)
	return res, err
}

func UploadFile(url string, params map[string]string, nameField, fileName string, file io.Reader) ([]byte, error) {
	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)

	formFile, err := writer.CreateFormFile(nameField, fileName)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(formFile, file)
	if err != nil {
		return nil, err
	}

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	//req.Header.Set("Content-Type","multipart/form-data")
	req.Header.Add("Content-Type", writer.FormDataContentType())

	resp, err := HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return content, nil
}
