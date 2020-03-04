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

package logging

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var (
	trace, info, warn, errLogger       *log.Logger
	traceLog, infoLog, warnLog, errLog *log.Logger
)

// TraceLvl is fatal level log
const TraceLvl = 1

// InfoLvl is info level log
const InfoLvl = 2

// WarnLvl is warn level log
const WarnLvl = 3

// ErrorLvl is error level log
const ErrorLvl = 4

// Init initializes the loggers
func Init() {
	trace = log.New(os.Stdout, "[TRACE]:", log.LstdFlags)
	info = log.New(os.Stdout, "[INFO]:", log.LstdFlags)
	warn = log.New(os.Stdout, "[WARN]:", log.LstdFlags)
	errLogger = log.New(os.Stdout, "[ERROR]:", log.LstdFlags)

	filename := fmt.Sprintf("./logs/logs_%v", time.Now())

	if err := ioutil.WriteFile(filename, []byte{}, os.ModePerm); err != nil {
		panic(err)
	}

	f, err := os.OpenFile(filename, os.O_APPEND, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	traceLog = log.New(f, "[TRACE]:", log.LstdFlags)
	infoLog = log.New(f, "[INFO]:", log.LstdFlags)
	warnLog = log.New(f, "[WARN]:", log.LstdFlags)
	errLog = log.New(f, "[ERROR]:", log.LstdFlags)
}

// Log logs log to the indicated level
func Log(level byte, log string) {
	switch level {
	case TraceLvl:
		trace.Println(log)
		traceLog.Println(log)
	case InfoLvl:
		info.Println(log)
		infoLog.Println(log)
	case WarnLvl:
		warn.Println(log)
		warnLog.Println(log)
	case ErrorLvl:
		errLogger.Println(log)
		errLog.Println(log)
	}
}

// LogFatal logs to a crashes file
func LogFatal(v ...interface{}) {
	filename := fmt.Sprintf("./logs/crashes/[Crash]%v", time.Now())

	ioutil.WriteFile(filename, []byte(fmt.Sprintln(v...)), 0666)
}
