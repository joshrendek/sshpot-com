package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB gorm.DB

func main() {
	var err error
	dbString := os.Getenv("DO_POSTGRESQL")
	if dbString == "" {
		dbString = os.Getenv("HEROKU_POSTGRESQL_BLUE_URL")
		if dbString == "" {
			dbString = "user=joshrendek port=5432 dbname=ssh_honey sslmode=disable"
		}
	}
	DB, err = gorm.Open("postgres", dbString)

	if err != nil {
		panic(err)
	}

	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(50)
	DB.DB().Ping()
	DB.LogMode(true)

	DB.CreateTable(SshLogin{})

	http.HandleFunc("/", index)
	http.HandleFunc("/join", join)
	http.HandleFunc("/api/private/ssh", ssh)
	http.HandleFunc("/api/ssh_logins.json", sshLoginList)
	fmt.Println("listening...")

	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

func join(res http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("views/run.html", "views/header.html")
	if err != nil {
		panic(err)
	}
	err = t.Execute(res, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func sshLoginList(res http.ResponseWriter, req *http.Request) {
	var logins []SshLogin
	var page int64
	var resp []byte
	var per_page int64 = 50
	var total int64
	var err error

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

	fmt.Println(string(resp))

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

	DB.Save(&login)
	fmt.Fprintln(res, "{}")
}

func index(res http.ResponseWriter, req *http.Request) {
	var logins []SshLogin
	var total int64
	DB.Model(SshLogin{}).Count(&total)
	DB.Limit(25).Order("id desc").Find(&logins)
	//view, _ := ioutil.ReadFile("views/index.html")
	t, err := template.ParseFiles("views/index.html", "views/header.html")
	if err != nil {
		panic(err)
	}
	err = t.Execute(res, struct {
		Logins []SshLogin
		Total  int64
	}{logins, total})
	if err != nil {
		fmt.Println(err)
	}
}
