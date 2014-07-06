package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type MaskedServerOutput struct {
	MaskedIP  string
	UpdatedAt string
}

func listHoneypotServers(res http.ResponseWriter, req *http.Request) {
	var servers []HoneypotServer
	var maskedServers []MaskedServerOutput
	const time_layout = "01/02/2006 3:04pm"
	DB.Find(&servers)

	for _, s := range servers {
		mask := strings.Split(s.RemoteAddr, ".")
		tmp := MaskedServerOutput{
			MaskedIP:  fmt.Sprintf("%s.x.x.%s", mask[0], mask[3]),
			UpdatedAt: s.UpdatedAt.Format(time_layout),
		}
		maskedServers = append(maskedServers, tmp)
	}

	t, err := template.ParseFiles("views/honeypot_list.html", "views/header.html", "views/footer.html")
	if err != nil {
		panic(err)
	}
	err = t.Execute(res, struct{ Servers []MaskedServerOutput }{maskedServers})
	if err != nil {
		fmt.Println(err)
	}
}

func join(res http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("views/run.html", "views/header.html", "views/footer.html")
	if err != nil {
		panic(err)
	}
	err = t.Execute(res, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func index(res http.ResponseWriter, req *http.Request) {
	var logins []SshLogin
	var total int64
	DB.Model(SshLogin{}).Count(&total)
	DB.Limit(25).Order("id desc").Find(&logins)
	//view, _ := ioutil.ReadFile("views/index.html")
	t, err := template.ParseFiles("views/index.html", "views/header.html", "views/footer.html")
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
