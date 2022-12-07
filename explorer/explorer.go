package explorer

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/pumpkinzomb/zombcoin/blockchain"
)

const (
	templateDir = "explorer/template/"
)

var templates *template.Template 

type homeData struct {
	PageTitle string
	Blocks []*blockchain.Block
}

func addBlockHandler(w http.ResponseWriter, r *http.Request) {
	switch(r.Method){
		case "GET":
			templates.ExecuteTemplate(w, "addblock", nil)
		case "POST":
			r.ParseForm()
			data := r.Form.Get("blockData")
			blockchain.GetBlockChain().AddBlock(data)
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

func homeHandler (w http.ResponseWriter, r *http.Request) {
	data := homeData{"Main page", blockchain.GetBlockChain().Blocks}
	templates.ExecuteTemplate(w, "home", data)
}

func Run (_port int) {
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))

	handler := http.NewServeMux()

	handler.HandleFunc("/", homeHandler)
	handler.HandleFunc("/add", addBlockHandler)
	
	fmt.Printf("Server is inited on localhost:%d\n", _port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", _port), handler))
}