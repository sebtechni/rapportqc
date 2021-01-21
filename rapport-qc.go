package main

import (
	// "fmt"
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
	"log"
	"encoding/xml"
	//"reflect"
	"rapport-qc/lib"
)

type StartTimeCodeSMPTE struct {
	Value			string				`xml:"value,attr"`
}

type DurationSMPTE struct {
	Value			string				`xml:"value,attr"`
}

type DropFrameFlag struct {
	Value			bool				`xml:"value,attr"`
}

type ChromaBitDepth struct {
	Value			int8				`xml:"value,attr"`
}

type TimeCodeTrack struct {
	DropFrameFlag			DropFrameFlag
	DurationSMPTE			DurationSMPTE
	StartTimeCodeSMPTE		StartTimeCodeSMPTE
}

type ActualTracks struct {
	Value			string				`xml:"value,attr"`
}

type Field struct {
	Name			string				`xml:"name,attr"`	
	Value			string				`xml:"value,attr"`
	ActualTracks	ActualTracks				
	TimeCodeTrack	TimeCodeTrack	
	ChromaBitDepth	ChromaBitDepth
}

type Info struct {
	Field			[]Field				`xml:"field"`	
}

type Error struct {
	FrameDuration				int				`xml:",attr"`
	SMPTETimecodeDuration		string			`xml:",attr"`
	Description					string			`xml:"description,attr"`
	Endsmptetimecode			string			`xml:"endsmptetimecode,attr"`
	Item						string			`xml:"item,attr"`
	Startsmptetimecode			string			`xml:"startsmptetimecode,attr"`
	Synopsis					string			`xml:"synopsis,attr"`
}

type Decodedaudiochecks struct {
	Name			string				`xml:"name,attr"`	
	Error			[]Error				`xml:"error"`	
}

type Decodedvideochecks struct {
	Name			string				`xml:"name,attr"`	
	Error			[]Error				`xml:"error"`	
}

type Customchecks struct {
	Decodedvideochecks			Decodedvideochecks				`xml:"decodedvideochecks"`	
	Decodedaudiochecks			Decodedaudiochecks				`xml:"decodedaudiochecks"`	
}

type Errors struct {
	Customchecks			Customchecks				`xml:"customchecks"`	
}

type Streamnode struct {
	EncodedDuration			string				`xml:",attr"`
	Name					string				`xml:"name,attr"`
	Info					Info				`xml:"info"`
	Errors					Errors				`xml:"errors"`	

}

type Toplevelinfo struct {
	Error			string				`xml:",attr"`	
	Format			string				`xml:",attr"`
}

type TaskReport struct {	
	XMLName 			xml.Name					`xml:"taskReport"`
	Toplevelinfo		Toplevelinfo				`xml:"toplevelinfo"`
	Streamnode  		[]Streamnode				`xml:"streamnode"`
}

var TEST string = "LLL"

func main() {
//00000176e33b7de348c2b333000a0063008a0046
//0000016e3d423cb80b853502000a0063008a0046
		var v TaskReport
		xml.Unmarshal(RestCall("0000016e3d423cb80b853502000a0063008a0046"), &v)
	
		log.Println(v)

		lib.WriteHTML()

		log.Println(TEST)
}

func RestCall(BatonTaskId string) []byte {
	client := resty.New()
	resp, _ := client.R().
		SetBasicAuth("dev_user", "jhg45W&sd_18").
		SetHeader("Content-Type", "application/x-www-form-urlencoded").		
		Get("http://baton-2:8080/Baton/api/tasks/"+BatonTaskId+"/report?type=xml")
		
		//fmt.Println(resp.String())
		//fmt.Println(reflect.TypeOf(resp.String()))
	return []byte(resp.String())

}
