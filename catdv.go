package main

import (
	"github.com/resty"	
	"log"
	"encoding/json"	
	// "io/ioutil"
	// "fmt"
)

func main() {

	// client := resty.New()
	// fmt.Println(fmt.Sprintf("%T", client))
	// resp, _ := client.R().	
	// 	Get("http://10.99.139.220:8080/api/9/session?usr=api_admin&pwd=Brute-Fence8-Backboned")
		
	// 	//fmt.Println(resp.String())
	// 	//fmt.Println(reflect.TypeOf(resp.String()))
	// //log.Println(resp.String())
	
	// resp, _ = client.R().	
	// 	Get("http://10.99.139.220:8080/api/9/clips/420127")
	// //log.Println("Clips",resp.String())
	// Clip := resp.Body()	
	
	// resp, _ = client.R().	
	// 	Delete("http://10.99.139.220:8080/api/9/session")
	// log.Println("Delete",resp.String())

	CatInfos := CatRest{Client: resty.New(), ClipID: "420127"}
	CatInfos.Open()
	Clip := CatInfos.Get()
	CatInfos.Delete()

	type Fields struct {
		Titre							string
		Version							string
		Provider						string
		Qc_type							string
		Validation_type					string		`json:"Validation.type"`
		Date							string
		Production_year					string		`json:"production.year"`
		Qc_operator						string		`json:"qc.operator"`
		Facility						string		
		Source_filename					string		`json:"source.filename"`
		Coordinator						string		
		Asset_type						string		`json:"asset.type"`
		Creation_de_rapports			[]string	`json:"création.de.rapports"`
		Baton_taskid					string		`json:"baton.taskid"`
		Narrative_text_presence_1		string		`json:"narrative.text.presence.1"`
		Burnedin_dialogs_presence		string		`json:"burned-in.dialogs.presence"`
		Nom_original					string		`json:"nom.original"`
	}	

	type Data struct  {
		Name		string
		Fields		Fields
	}

	type CatDV struct  {
		Data		Data
	}

	var InfosCatDV CatDV
	err := json.Unmarshal(Clip, &InfosCatDV)
	if err != nil {
		log.Println("JSON error:", err)
	}
	
	log.Println(InfosCatDV.Data.Fields)
	
}

type CatRest struct {
	Client		*resty.Client
	ClipID		string
}


func (CatInfos CatRest) Open() {
	CatInfos.Client.R().Get("http://10.99.139.220:8080/api/9/session?usr=api_admin&pwd=Brute-Fence8-Backboned")
}

func (CatInfos CatRest) Get() []byte {
	resp, _ := CatInfos.Client.R().Get("http://10.99.139.220:8080/api/9/clips/"+CatInfos.ClipID)
	
	return resp.Body()
}

func (CatInfos CatRest) Delete() {
	CatInfos.Client.R().	
		Delete("http://10.99.139.220:8080/api/9/session")
}


