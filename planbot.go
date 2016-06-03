// This file is part of the Plan-B-ot package.
// Copyright (c) 2015 Martin Schenck
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/apheleia/plan-b-ot/bot"
)

// Runs the server listening for the requests
func main() {
	err := bot.ReadConfig("config.json")
	if err != nil {
		panic(err.Error())
	}

	http.HandleFunc(bot.Config.Route, planbotHandler)
	http.ListenAndServe(":"+bot.Config.Port, nil)
}

// planbotHandler handles incoming requests
// can be task, vote, or result requests
func planbotHandler(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("token") != bot.Config.Token {
		http.Error(w, "Invalid token in slashcommand setup!", http.StatusBadRequest)
		return
	}

	userName := r.FormValue("user_name")
	text := r.FormValue("text")
	words := strings.Fields(text)

	response, status := bot.HandleRequest(userName, words)
	responseJson, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Cannot marshal JSON in plan-b-ot", http.StatusInternalServerError)
		return
	}

	if status != http.StatusOK {
		http.Error(w, response.Text, status)
		return
	}

	w.WriteHeader(status)
	w.Write(responseJson)
}
