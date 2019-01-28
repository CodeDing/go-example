package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"adeaz.com/rtb/adx/wifi"
	proto "github.com/golang/protobuf/proto"
)

var MediaFile_Map = map[string][]string{
	"wifi": []string{"json/wifi_big.json", "json/wifi_single.json", "json/wifi_three.json"},
}

var URL_Map = map[string]string{
	//"wifi": "http://sx.rtb.fastapi.net/?pid=17117",
	"wifi": "http://sx.rtb.ggxt.net/?pid=17117",
}

var Res_Map = map[string]interface{}{
	"wifi": wifi.RTBResponse{},
}

type Media struct {
	Pid   int
	Files []string
	URL   string
}

func init() {

}

var (
	ErrReadFile       = errors.New("failed to read file")
	ErrMarshalJson    = errors.New("failed to marshal json")
	ErrUnmarshalJson  = errors.New("failed to unmarshal json")
	ErrMarshalProto   = errors.New("failed to marshal proto")
	ErrUnmarshalProto = errors.New("failed to unmarshal proto")
	ErrNotExistMedia  = errors.New("media not exist")
	ErrNotExistURL    = errors.New("url not exist")
)

func json2pb(file, media string) ([]byte, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, ErrReadFile
	}
	var wifiReq wifi.RTBRequest
	err = json.Unmarshal(data, &wifiReq)
	if err != nil {
		return nil, ErrUnmarshalJson
	}
	pb, err := proto.Marshal(&wifiReq)
	if err != nil {
		return nil, ErrMarshalProto
	}
	fmt.Printf("[%s-request] pb => %s\n\n", strings.ToUpper(media), string(pb))
	bs, err := json.Marshal(wifiReq)
	if err != nil {
		return nil, ErrMarshalJson
	}
	fmt.Printf("[%s-request] js => %s\n\n", strings.ToUpper(media), string(bs))
	return pb, nil
}

func sendPbMessage(media string, b []byte) (string, error) {
	if url, ok := URL_Map[media]; ok {
		resp, err := http.Post(url,
			"application/octet-stream",
			strings.NewReader(string(b)))
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		fmt.Printf("[%s-response] pb => %s\n\n", strings.ToUpper(media), string(body))
		if res, exist := Res_Map[media]; exist {
			switch res.(type) {
			case wifi.RTBResponse:
				var wifiRes wifi.RTBResponse
				if err := proto.Unmarshal(body, &wifiRes); err != nil {
					return "", err
				}
				bs, err := json.Marshal(wifiRes)
				if err != nil {
					return "", err
				}
				fmt.Printf("[%s-response] js => %s\n", strings.ToUpper(media), string(bs))
				return string(bs), nil
			}
		}
	}
	return "", ErrNotExistURL
}

const (
	Format = "==================    %s(%s)     ================\n"
	Footer = `==============================================================`
)

func main() {

	medias := [...]string{"wifi"}
	for _, m := range medias {
		if files, ok := MediaFile_Map[m]; ok {
			for _, f := range files {
				fmt.Printf(Format, m, f)
				bs, err := json2pb(f, m)
				if err != nil {
					panic(err)
				}
				sendPbMessage(m, bs)
				fmt.Println(Footer)
			}
		}
	}
}

func httpDo() {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://www.01happy.com/demo/accept.php", strings.NewReader("name=cjb"))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))
}
