package main

import (
	"net/http"
	"fmt"
)

func FireAndForget(method string, url string) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(fmt.Sprintf("Error while building url request [%s]", url))
	}
	req.SetBasicAuth(GetXlaConfig().Xld.Login, GetXlaConfig().Xld.Password)

	client := &http.Client{}
	client.Do(req)
}
