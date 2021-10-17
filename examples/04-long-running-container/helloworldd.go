// Copyright 2019 Jomoespe. All rights reserved.
// Use of this source code is governed by a The Unlicense license that can be
// found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var instanceName = "unamed"

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "server instance name: %s\n", instanceName)
	name := req.URL.Query().Get("name")
	log.Printf("URI: %s, query params=%s", req.RequestURI, req.URL.Query())
	fmt.Fprintf(w, "Hello %s!\n", name)
}

func main() {
	if len(os.Args) >= 2 {
		instanceName = os.Args[1]
	}
	log.Printf("Starting server instance name: %s\n", instanceName)

	http.HandleFunc("/", hello)
	http.ListenAndServe(":80", nil)
}
