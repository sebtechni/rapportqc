package main

import (
	"fmt"
	// "encoding/xml"
	// "time"	
	// "regexp"		
	// "strings"
	// "xlsx"	
	//"os"
	//"strconv"
	// "path"
	//"encoding/json"	
	//"io/ioutil"
	"github.com/resty"	
	//"os/exec"
	//"net/http"
	//"log"
)

func main() {
	fmt.Println("hello")

	client := resty.New()

	resp, _ := client.R().
		SetBasicAuth("dev_user", "jhg45W&sd_18").
		SetHeader("Content-Type", "application/x-www-form-urlencoded").		
		Get("http://baton-2:8080/Baton/api/tasks/00000176e33b7de348c2b333000a0063008a0046/report?type=xml")

		fmt.Println(resp)
		
	

}

