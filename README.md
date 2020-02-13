# GoLang Google Cloud Function Firestore Template

## Contents
1. [Stack](#stack)
2. [Useful Docs](#useful-docs)
3. [Description](#description)
4. [Responses](#responses)
5. [Getting Started](#getting-started)
6. [Testing](#testing)
7. [Deployment](#deployment)
8. [License](#license)

## Stack
* Written in GoLang
* Hosted on a [Google Cloud Function](https://console.cloud.google.com/functions)
* Backend/Database is hosted via [Google Cloud Firestore](https://console.cloud.google.com/firestore)

## Useful Docs
* Outputting JSON in GoLang https://gobyexample.com/json
* GoLang with Firestore Examples + Quick Start https://github.com/GoogleCloudPlatform/golang-samples/tree/master/firestore 
* official GoLang Package Docs for firestore package  https://pkg.go.dev/cloud.google.com/go/firestore?tab=doc


## Description
This is the skeleton of a google cloud function written in GoLang for a lookup/catalog service with a Google Firestore database backend.
For example.

GET request to https://region-account.cloudfunctions.net/FunctionName?object_id=1
 
POST request to https://region-account.cloudfunctions.net/FunctionName with form parameter of object_id = 1 

These can both return a flattened JSON object from the collection you specify
```
{"object_id":"1", "name":"Object 1", "description":"oooh.... shiny"}
```

## Responses
* 200 / http.StatusOK - This means the object was found, will also return the object in JSON format
* 422 / http.StatusUnprocessableEntity - The ID was not passed in, or wasn't a number
* 422 / http.StatusUnprocessableEntity - The query didn't return any results
* 500 / http.StatusInternalServerError - The function cannot connect to Firestore
* 500 / http.StatusInternalServerError - The Object retrieved cannot be converted to JSON format properly

## Getting Started
1. Clone/Download the repository and remove all git references
```
git clone https://github.com/hallindavid/go-gcpfunc-firestore-template /path/to/project
cd /path/to/project
rm -rf .git

# If desired, init your git repo here
# git init
```

2.  Follow the steps in this link to create a json file to use for credentials
https://cloud.google.com/docs/authentication/production#obtaining_and_providing_service_account_credentials_manually

3.  Create your `.env` file in the root directory
4.  In your `.env` file, put the following line of code
```
GOOGLE_APPLICATION_CREDENTIALS="/full/system/path/to/json-credentials.json"
```
5.  Update the gopher.go file (or copy/paste) to a more suitable name
eg: if you are you retrieving users, maybe make it `users.go` and `func Userlookup(...`

6.  download packages / build the module
```
go get
go mod init
```

7. Run it to make sure everything works - see [Testing Locally](#local-testing)
8. Make your tests
9. perform a `go test` to make sure everything is working
10.  Deploy it - see [Deployment](#deployment)

## Testing
#### Quick Test All
```
go test
```

#### Test All With Output
```
go test -v
```
#### Local Testing
in the cmd folder (after you've generated your go.mod file) update the main.go file to import the proper module
```
package main

import (
	"log"
	"net/http"

	"github.com/hallindavid/go-gcpfunc-firestore-template" //This line will need to be updated
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	http.HandleFunc("/", gopher.FirestoreLookup) //this line will need to point to the correct function (eg. userlookup.LookupUsers)

	log.Printf("Listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

```

in terminal/command prompt
```
go run ./cmd
```
This will listen on port 8080 and allow you try to postman/browse to your functions

## Deployment
Follow deployment steps found [here](https://cloud.google.com/functions/docs/concepts/go-runtime)
```
#eg
gcloud deploy UserLookup --region us-central1 --runtime go113 --trigger-http
```


## Licenses
* GoLange is licensed here: https://golang.org/LICENSE
* GoDotEnv is licensed (MIT) here: https://github.com/joho/godotenv/blob/master/LICENCE
* Firestore is licensed (Apache-2.0) here: https://pkg.go.dev/cloud.google.com/go/firestore?tab=licenses#LICENSE

## Use of the GoDotEnv Package
We are using godotenv (https://github.com/joho/godotenv) to set the default Firestore application credentials for testing - but it isn't required for production.
You can actually remove the package altogther and do a similar kind of thing in a bashfile
```
# build.sh in project root directory
export GOOGLE_APPLICATION_CREDENTIALS="/full/system/path/to/json-credentials.json"
go run ./cmd
```


```
# test.sh in project root directory
export GOOGLE_APPLICATION_CREDENTIALS="/full/system/path/to/json-credentials.json"
go test
#or go test -v
```