package main

import (
	"github.com/resty"	
	"log"
	"encoding/json"	
	// "io/ioutil"
	// "fmt"
	"reflect" //fmt.Println(reflect.TypeOf(resp.String()))
)
type Fields struct {
	Titre							string		`json:"titre"`
	// Version							string		`json:"version"`
	// Provider						string		`json:"provider"`
	// Qc_type							string		`json:"qc.type"`
	// Validation_type					string		`json:"validation.type"`
	// Date							string		`json:"date"`
	// Production_year					string		`json:"production.year"`
	// Qc_operator						string		`json:"qc.operator"`
	// Facility						string		`json:"facility"`
	// Source_filename					string		`json:"source.filename"`
	// Coordinator						string		`json:"coordinator"`
	// Asset_type						string		`json:"asset.type"`
	// Creation_de_rapports			[]string	`json:"création.de.rapports"`
	// Baton_taskid					string		`json:"baton.taskid"`
	// Narrative_text_presence_1		string		`json:"narrative.text.presence.1"`
	// Burnedin_dialogs_presence		string		`json:"burned-in.dialogs.presence"`
	// Nom_original					string		`json:"nom.original"`
}	

type Data struct  {
	Name		string		`json:"name"`
	Fields		Fields		`json:"fields"`
}

type CatDV struct  {
	Data		Data		`json:"data"`
}

type CatRest struct {
	Client		*resty.Client
	ClipID		string
	CatDV		CatDV
	JsonData	string
}

func main() {
//436847 - 420127 - 420108
	CatInfos := CatRest{Client: resty.New(), ClipID: "420108"}
	CatInfos.Open()
	Clip := CatInfos.Get()
	
	
	log.Println(string(Clip))
	
	

	// var InfosCatDV CatDV
	// err := json.Unmarshal(Clip, &InfosCatDV)

	err := json.Unmarshal(Clip, &CatInfos.CatDV)
	if err != nil {
		log.Println("JSON error:", err)
	}
	
	//log.Println(CatInfos.CatDV.Data.Fields)

	// log.Println(CatInfos.CatDV.Data.Fields.Creation_de_rapports[0])
	// CatInfos.CatDV.Data.Fields.Creation_de_rapports[0] = ""
	log.Println(CatInfos.CatDV.Data.Fields.Titre)
	CatInfos.CatDV.Data.Fields.Titre = "New Pony TITLE"
	log.Println(CatInfos.CatDV.Data.Fields.Titre)


	jsonData, err := json.Marshal(CatInfos.CatDV)
	if err != nil {
		log.Println(err)
	}
	CatInfos.JsonData = string(jsonData)
	log.Println(CatInfos.CatDV)
	log.Println(string(jsonData))
	log.Println(reflect.TypeOf(jsonData))

	CatInfos.Update()
	CatInfos.Delete()
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

func (CatInfos CatRest) Update() {
	//SetBody(`{"username":"testuser", "password":"testpass"}`).
	resp, _ := CatInfos.Client.R().
		//SetBody(CatInfos.CatDV).
		SetBody(`{"fields":{"création.de.rapports" : false}}`).
		Put("http://10.99.139.220:8080/api/9/clips/"+CatInfos.ClipID)	
	log.Println(resp)
}

















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