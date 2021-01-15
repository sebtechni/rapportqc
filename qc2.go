package main

import (
    "fmt"
    "net/http"
    "net/url"
    "strconv"
    "strings"
)

func main() {
    apiUrl := "http://baton-2:8080"
    resource := "/Baton/api/tasks/00000176e33b7de348c2b333000a0063008a0046/report?type=xml"
    data := url.Values{}
    data.Set("name", "daudelins")
    data.Set("surname", "!Microbrasserie2021")

    u, _ := url.ParseRequestURI(apiUrl)
    u.Path = resource
    urlStr := u.String() // "https://api.com/user/"

    client := &http.Client{}
    r, _ := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode())) // URL-encoded payload
    r.Header.Add("Authorization", "auth_token=\"ZGF1ZGVsaW5zOiFNaWNyb2JyYXNzZXJpZQ==\"")
    r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

    resp, _ := client.Do(r)
    fmt.Println(resp)
}