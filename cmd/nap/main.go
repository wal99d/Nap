package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	nap "github.com/wal99d/Nap"
)

var api = nap.NewAPI("https://httpbin.org")

func main() {
	list := flag.Bool("list", false, "Get a list of api resources")
	flag.Parse()
	if *list {
		fmt.Println("Available resources..")
		for _, name := range api.ResourceNames() {
			fmt.Println(name)
		}
		return
	}
	resource := os.Args[1]
	if err := api.Call(resource, nil, nil); err != nil {
		log.Fatalln("Unable to get resources!!")
	}
}

func init() {
	router := nap.NewRouter()
	router.RegisterFunc(200, func(resp *http.Response, _ interface{}) error {
		defer resp.Body.Close()

		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		fmt.Println(string(content))
		return nil
	})
	api.AddResource("get", nap.NewResource("/get", "GET", router))
}
