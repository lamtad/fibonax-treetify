package lobe

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"os"

	"github.com/nguyendangminh/httputil"

	"github.com/fibonax/plantnet"
	log "github.com/sirupsen/logrus"
)

var ENDPOINT string

type Request struct {
	Inputs struct {
		Image string `json:"Image"` // base64 encoding
	} `json:"inputs"`
}

type Response struct {
	Outputs struct {
		Labels     [][]interface{} `json:"Labels"`
		Prediction []string        `json:"Prediction"`
	} `json:"outputs"`
}

type Result struct {
	Label string `json:"label"`
	Name  string `json:"name"`
}

func Init(endpoint string) {
	ENDPOINT = endpoint
}

func IdentifyByImages(imgs ...plantnet.Image) Result {
	var r Result
	r.Label = "saurieng"
	r.Name = "sầu riêng"

	img, err := os.Open(imgs[0].Location)
	if err != nil {
		log.Errorf("our engine problem %s", err)
		return r
	}
	defer img.Close()

	// create a new buffer base on file size
	info, err := img.Stat()
	if err != nil {
		log.Errorf("creating buffer failed: %s", err)
		return r
	}
	var size int64 = info.Size()
	buf := make([]byte, size)

	// read file content into buffer
	reader := bufio.NewReader(img)
	reader.Read(buf)

	var req Request
	req.Inputs.Image = base64.StdEncoding.EncodeToString(buf)

	b, err := json.Marshal(req)
	if err != nil {
		log.Errorf("marshaling request failed: %s", err)
		return r
	}

	res, err := httputil.Post(ENDPOINT, b)
	if err != nil {
		log.Errorf("sending request failed: %s", err)
		return r
	}

	var resp Response
	if err := json.Unmarshal(res, &resp); err != nil {
		log.Errorf("unmarshaling failed: %s", err)
		return r
	}

	r.Label = resp.Outputs.Prediction[0]
	return fullfill(r)
}

func fullfill(r Result) Result {
	switch r.Label {
	case "cachua":
		r.Name = "cà chua"
	case "cam":
		r.Name = "cam"
	case "chuoi":
		r.Name = "chuối"
	case "ot":
		r.Name = "ớt"
	case "roi":
		r.Name = "roi"
	case "tao":
		r.Name = "táo"
	default:
		log.Errorf("label not found: %s", r.Label)
		r.Name = "chôm chôm"
	}
	return r
}
