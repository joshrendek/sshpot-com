package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func ssh_http(res http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	ssh_http := SshHttp{Request: string(body)}
	fmt.Println(string(body))
	ssh_http.CreatedAt = time.Now()
	fmt.Println(ssh_http)
	if err != nil {
		panic(err)
	}

	DB.Save(&ssh_http)
	fmt.Fprintln(res, "{}")
}

type ProxyResponse struct {
	Id        int64             `json:"id"`
	Data      map[string]string `json:"data"`
	CreatedAt time.Time         `json:"created_at"`
}

func sshProxyList(res http.ResponseWriter, req *http.Request) {
	var http_requests []SshHttp
	var page int64
	var resp []byte
	var per_page int64 = 50
	var total int64
	var err error

	UpdateApiCounters()

	if len(req.URL.Query()["page"]) > 0 {
		page, err = strconv.ParseInt(req.URL.Query()["page"][0], 10, 64)
		if err != nil {
			resp, _ = json.Marshal(struct{ Message string }{"Invalid page parameter"})
		}
	} else {
		page = 1
	}

	data := []ProxyResponse{}

	limit_string := strconv.FormatInt(per_page, 10)
	offset_string := strconv.FormatInt(((page * per_page) - per_page), 10)
	DB.Order("id desc").Limit(limit_string).Offset(offset_string).Find(&http_requests)
	DB.Model(SshHttp{}).Count(&total)

	for _, sshReq := range http_requests {
		json_data := make(map[string]string)

		json.Unmarshal([]byte(sshReq.Request), &json_data)
		data = append(data, ProxyResponse{
			Id:        sshReq.Id,
			Data:      json_data,
			CreatedAt: sshReq.CreatedAt,
		})
	}

	resp_data := struct {
		Page  int64           `json:"page"`
		Total int64           `json:"total"`
		Data  []ProxyResponse `json:"data"`
	}{
		page,
		total,
		data,
	}

	resp, err = json.Marshal(resp_data)

	if err != nil {
		fmt.Println(err)
		resp, _ = json.Marshal(struct{ Message string }{fmt.Sprintf("Invalid json: %s", err)})
	}

	fmt.Fprintln(res, string(resp))
}
