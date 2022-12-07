package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pumpkinzomb/zombcoin/blockchain"
	"github.com/pumpkinzomb/zombcoin/utils"
)

var port string

type url string

func(u url) MarshalText() ([]byte, error) {
	uri := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(uri), nil
}

type urlDescription struct{
	URL url `json:"url"`
	Method string `json:"method"`
	Description string `json:"description"`
	Payload	*payload `json:"payload,omitempty"`
}

type payload struct {
	Data string `json:"data"`
}

type addBlockBody struct {
	Message string
}

type errorResponse struct {
	ErrorMessage string `json:"error_message"`
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	routes := []urlDescription{
		{
			URL: url("/"),
			Method: "GET",
			Description: "This url responses all of REST server's API documentation.",
		},
		{
			URL: url("/blocks"),
			Method: "GET",
			Description: "This url gets all of blocks list.",
		},
		{
			URL: url("/blocks/{:height}"),
			Method: "GET",
			Description: "This url gets inputted height of block detail informations.",
		},
		{
			URL: url("/block"),
			Method: "POST",
			Description: "If you send with data in this url, then you can create new block.",
			Payload: &payload{Data: "string *required"},
		},
	}
	codec := json.NewEncoder(w)
	err := codec.Encode(routes)
	utils.CheckErr(err)
}

func handleBlock(w http.ResponseWriter,r *http.Request) {
	switch r.Method {
		case "POST": 
		var addBlockBody addBlockBody
		codec := json.NewDecoder(r.Body)
		err := codec.Decode(&addBlockBody)
		utils.CheckErr(err)
		blockchain.GetBlockChain().AddBlock(addBlockBody.Message)
		w.WriteHeader(http.StatusCreated)
	}
}

func handleBlocks(w http.ResponseWriter,r *http.Request) {
	codec := json.NewEncoder(w)
	err := codec.Encode(blockchain.GetBlockChain().AllBlocks())
	utils.CheckErr(err)
}

func handleSingleBlock(w http.ResponseWriter,r *http.Request){
	getParams := mux.Vars(r)
	height, err := strconv.Atoi(getParams["height"])
	utils.CheckErr(err)
	block, err := blockchain.GetBlockChain().GetBlock(height)
	codec := json.NewEncoder(w)
	if(err == blockchain.ErrNotFound){
		codec.Encode(errorResponse{fmt.Sprint(err)})
	}else {
		codec.Encode(block)
	}
}

func setJSONheaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
		w.Header().Add("content-type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func Run(_port int) {
	router := mux.NewRouter()
	port = fmt.Sprintf(":%d", _port)
	router.Use(setJSONheaderMiddleware)
	router.HandleFunc("/", handleIndex).Methods("GET")
	router.HandleFunc("/block", handleBlock).Methods("POST")
	router.HandleFunc("/blocks", handleBlocks).Methods("GET")
	router.HandleFunc("/blocks/{height:[0-9]+}", handleSingleBlock).Methods("GET")
	
	fmt.Printf("Rest Server is on http://localhost%s \n", port)
	log.Fatal(http.ListenAndServe(port, router))
}