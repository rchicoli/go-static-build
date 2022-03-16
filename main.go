package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"time"
)

func main() {

	fmt.Println("hello world")

	//
	// User
	//
	usr, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("User: %s\n", usr.Name)
	fmt.Printf("UID: %s\n", usr.Uid)

	//
	// Time
	//
	now1 := time.Now()
	fmt.Println(now1)
	//
	tz := os.Getenv("TZ")
	time.Local, _ = time.LoadLocation(tz)
	now2 := time.Now()
	fmt.Println(now2)

	//
	// Net
	//
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://icanhazip.com")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s", body)

}
