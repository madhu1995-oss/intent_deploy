Skip to content
Search or jump to…

Pull requests
Issues
Marketplace
Explore
 
@madhu1995-oss 
Learn Git and GitHub without any code!
Using the Hello World guide, you’ll start a branch, write comments, and open a pull request.


45
460521GoogleCloudPlatform/kubernetes-engine-samples
 Code Issues 6 Pull requests 2 Actions Projects 0 Wiki Security Insights
kubernetes-engine-samples/hello-app/main.go / 
@ahmetb ahmetb Clarify hello-app/main.go with comments (#47)
f9c5196 on Mar 8, 2018
53 lines (46 sloc)  1.4 KB
  
 You're using jump to definition to discover and navigate code.
Learn more or give us feedback
/**
 * Copyright 2017 Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// [START all]
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// use PORT environment variable, or default to 8080
	port := "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	// register hello function to handle all requests
	server := http.NewServeMux()
	server.HandleFunc("/", hello)

	// start the web server on port and accept requests
	log.Printf("Server listening on port %s", port)
	err := http.ListenAndServe(":"+port, server)
	log.Fatal(err)
}

// hello responds to the request with a plain-text "Hello, world" message.
func hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving request: %s", r.URL.Path)
	host, _ := os.Hostname()
	fmt.Fprintf(w, "Hello, world! i am intent classifier\n")
	fmt.Fprintf(w, "Version: 1.0.0\n")
	fmt.Fprintf(w, "Hostname: %s\n", host)
}

