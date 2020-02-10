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

package repository

import (
	"database/sql"
	"fmt"
	"net/url"

	// mssql driver
	_ "github.com/denisenkom/go-mssqldb"
)

// GetSQLDb tries connects to mssql db.
// Panics with sql.Open() errors if it fails.
func GetSQLDb(username, password, hostname string, port int) *sql.DB {
	query := url.Values{}
	query.Add("Authentication API", "AuthenticationAPI")

	u := url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(username, password),
		Host:     fmt.Sprintf("%s:%d", hostname, port),
		RawQuery: query.Encode(),
	}

	db, err := sql.Open("sqlserver", u.String())

	if err == nil {
		return db
	}

	panic(err)
}
