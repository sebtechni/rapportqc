package main

import (
	"fmt"
	"html/template"
	"os"
	"defaults"
	"go-wkhtmltopdf"
	"log"
	"io/ioutil"	
	"strings"
	
)

func main() {

	type GENCOM struct {
		CODE	[]string		`default:"[\"fesg\", \"t5jhyhgd\", \"eth5\"]"`
		VALUE	[]string		`default:"[\"htrjryj\", \"sae4rg45\", \"sgfnty\"]"`
	}

	type QCInfos struct {
		TITLE		string			`default:"John Smith"`
		FILENAME	[]string		`default:"[\"HAHA\", \"fefe\", \"efe\"]"`
		GENCOM		GENCOM			
	}


	type Enfant struct {
		UnEnfant	[]string
	}
	type User struct {
		Name  string
		Email string
		Enfant Enfant 
	}

	t := template.Must(template.New("").Parse(`<div>Foobar {{ (index .Enfant.UnEnfant 1) }}</div>`))
	tdeux := template.Must(template.ParseFiles("C:/Users/daudels/go/src/template_test/test.txt"))
	

	// m := map[string]interface{}{
	// 		"Users": []User{
	// 			{Name: "Bob", Email: "bob@myco.com"},
	// 			{Name: "Alice", Email: "alice@myco.com"},
	// 		},
	// 	}

	testtest := map[string]*User{"UsersDeux": &User{Name: "Bob", Email: "bob@myco.com"}}
	//testtest["UsersDeux"].Name = "HOLLO"
	testtest["UsersDeux"].Enfant.UnEnfant = append(testtest["UsersDeux"].Enfant.UnEnfant, "ouf, yap")
	testtest["UsersDeux"].Enfant.UnEnfant = append(testtest["UsersDeux"].Enfant.UnEnfant, "deux")
	fmt.Println(testtest["UsersDeux"].Enfant.UnEnfant[0])



	test := []User{
		{Name: "BobDU1er", Email: "bob@myco.com"},
		{Name: "AliceDU1er", Email: "alice@myco.com"},
	}
	test[0].Enfant.UnEnfant = append(test[0].Enfant.UnEnfant, "allo")
	test[0].Enfant.UnEnfant = append(test[0].Enfant.UnEnfant, "deuxieme")

	fmt.Println(test[0])
	fmt.Println(testtest["UsersDeux"])
	
	
	// fmt.Println(m["Users"])
	//fmt.Println(m)
	fmt.Println(t.Execute(os.Stdout, test[0]))

	fmt.Println(tdeux.Execute(os.Stdout, test[0]))

	fmt.Println("-----------------------")
	obj := &QCInfos{}
	if err := defaults.Set(obj); err != nil {
		panic(err)
	}
	fmt.Println(obj)


	fmt.Println("-----------------------")
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
	  log.Fatal(err)
	}
	// Set global options
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)
	pdfg.Grayscale.Set(true)

	htmlfile, err := ioutil.ReadFile("C:/Users/daudels/go/src/rapport-qc/template_qc_report.htm")
	if err != nil {
	log.Fatal(err)
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(string(htmlfile))))
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}
	err = pdfg.WriteFile("C:/Users/daudels/go/src/template_test/simplesample.pdf")
	if err != nil {
	  log.Fatal(err)
	}
}



// m := map[string]interface{}{
// 	"Users": []User{
// 		{Name: "Bob", Email: "bob@myco.com"},
// 		{Name: "Alice", Email: "alice@myco.com"},
// 	},
// }