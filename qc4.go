package main

import (
	"fmt"
	"os/exec"
	//"strings"
)

func main() {
	out, _ := exec.Command("curl", "-i", "-X","GET","-u","daudels:!Microbrasserie2021","-H","'Content-type:application/x-www-form-urlencoded'","http://baton-2:8080/Baton/api/tasks/00000176e33b7de348c2b333000a0063008a0046/report?type=xml").Output()
	tcstring := string(out[:])
	fmt.Println(tcstring)
}