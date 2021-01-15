package main

import(
    "fmt"
    "net/http"
    "io/ioutil"
    "log"
)

func basicAuth() string {
    var username string = "daudelins"
    var passwd string = "!Microbrasserie2021"
    client := &http.Client{}
    req, err := http.NewRequest("GET", "http://baton-2:8080/", nil)
    req.SetBasicAuth(username, passwd)
    resp, err := client.Do(req)
    if err != nil{
        log.Fatal(err)
    }
    bodyText, err := ioutil.ReadAll(resp.Body)
    s := string(bodyText)
    return s
}

func main(){
    fmt.Println("requesting...")
    S := basicAuth()
    fmt.Println(S)
}
 