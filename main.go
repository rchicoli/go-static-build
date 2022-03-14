package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/user"
	"time"
)

func main() {

	usr, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(usr.Name)
	fmt.Println(usr.Uid)
	fmt.Println(usr.Username)

	now := time.Now()
	fmt.Println(now)
	location := now.Location()
	fmt.Println(location.String())

	var brazil, _ = time.LoadLocation("Brazil/West")
	// time.Now().In(brazil).Format("2006-01-02 15:04:05")
	fmt.Println(time.Now().In(brazil))

	resp, err := http.Get("https://icanhazip.com/")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s", body)

}
