/*
Copyright 2020 Yi Fan Song

This file is part of mini-octo-giggle.

mini-octo-giggle is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

mini-octo-giggle is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with mini-octo-giggle. If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"fmt"
	"html/template"
	"mini-octo-giggle/errorhandling"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func loadDbEnv() (username, password, hostname string, port int) {

	username = os.Getenv("dbusername")
	if username == "" {
		panic("Database username cannot be empty")
	}

	password = os.Getenv("dbpassword")
	if password == "" {
		panic("Database password cannot be empty")
	}

	hostname = os.Getenv("dbhostname")
	if hostname == "" {
		panic("Database hostname cannot be empty")
	}

	port, err := strconv.Atoi(os.Getenv("dbport"))
	if err != nil {
		panic("Database port parsing failed: \n" + err.Error())
	}

	return
}

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
	</head>
	<body>
		{{range .Items}}<div>{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
	</body>
</html>`

	// t, err := template.New("webpage").Parse(tpl)
	// check(err)

	tmpl, err := template.ParseFiles(
		"./templates/home.html",
		"./templates/components/header.html",
		"./templates/components/footer.html",
	)
	errorhandling.CheckFatal(err, nil)

	data := struct {
		Title string
		Items []string
	}{
		Title: "My page",
		Items: []string{
			"My photos",
			"My blog",
		},
	}

	// err = tmpl.Execute(os.Stdout, data)

	// err = t.Execute(os.Stdout, data)
	// check(err)

	// noItems := struct {
	// 	Title string
	// 	Items []string
	// }{
	// 	Title: "My another page",
	// 	Items: []string{},
	// }

	// err = t.Execute(os.Stdout, noItems)
	// check(err)

	// username, password, hostname, dbPort := loadDbEnv()

	// repository.GetSQLDb(username, password, hostname, dbPort)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello World")
	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		// _, _ = fmt.Fprintf(w, "what")

		err = tmpl.Execute(w, data)
		errorhandling.CheckFatal(err, nil)
	})

	if port := os.Getenv("port"); port != "" {
		if err := http.ListenAndServe(port, nil); err != nil {
			panic(err)
		}
	} else {
		if err := http.ListenAndServe(":9990", nil); err != nil {
			panic(err)
		}
	}
}
