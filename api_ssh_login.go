package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func sshLoginList(res http.ResponseWriter, req *http.Request) {
	var logins []SshLogin
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

	limit_string := strconv.FormatInt(per_page, 10)
	offset_string := strconv.FormatInt(((page * per_page) - per_page), 10)
	DB.Order("id desc").Limit(limit_string).Offset(offset_string).Find(&logins)
	DB.Model(SshLogin{}).Count(&total)

	data := struct {
		Page  int64      `json:"page"`
		Total int64      `json:"total"`
		Data  []SshLogin `json:"data"`
	}{
		page,
		total,
		logins,
	}

	resp, err = json.Marshal(data)

	if err != nil {
		fmt.Println(err)
		resp, _ = json.Marshal(struct{ Message string }{fmt.Sprintf("Invalid json: %s", err)})
	}

	fmt.Fprintln(res, string(resp))
}

func ssh(res http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	login := SshLogin{}
	fmt.Println(string(body))
	err = json.Unmarshal(body, &login)
	login.CreatedAt = time.Now()
	fmt.Println(login)
	if err != nil {
		panic(err)
	}

	honeyPot := HoneypotServer{}
	var remoteAddr string
	_remote := req.Header["X-Forwarded-For"]
	if len(_remote) > 0 {
		tmp := strings.Split(_remote[len(_remote)-1], ",")
		remoteAddr = tmp[0]
		fmt.Println(fmt.Sprintf("************* %s", remoteAddr))
	} else {
		remoteAddr = strings.Split(req.RemoteAddr, ":")[0]
	}

	DB.Where(HoneypotServer{RemoteAddr: remoteAddr}).Assign(HoneypotServer{RemoteAddr: remoteAddr, UpdatedAt: time.Now()}).FirstOrCreate(&honeyPot)

	DB.Save(&login)
	fmt.Fprintln(res, "{}")
}
