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
package templates

import (
	"io/ioutil"
	"os"
)

func CreatePage(name string, header, footer bool) error {
	// TODO: Add to db

	baseTmpl, err := ioutil.ReadFile("base.tmpl")

	ioutil.WriteFile(name+".html", baseTmpl, os.ModePerm)
}
