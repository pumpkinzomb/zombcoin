package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/pumpkinzomb/zombcoin/explorer"
	"github.com/pumpkinzomb/zombcoin/rest"
)

func usage(){
	fmt.Println("################################################")
	fmt.Println("#                                              #")
	fmt.Println("#           Welcome to Zomb Chain CLI          #")
	fmt.Println("#                                              #")
	fmt.Println("#  -mode  Start with 'rest' or 'explorer'      #")
	fmt.Println("#  -port  Start with own port number           #")
	fmt.Println("#  -b     Start with both mode                 #")
	fmt.Println("#                                              #")
	fmt.Println("################################################")
	os.Exit(0)
}

func Run(){
	mode := flag.String("mode", "rest", "You can start with 'rest' or 'explorer' server")
	port := flag.Int("port", 4000, "You can set your port number")
	both := flag.Bool("b", false, "You can start both server")

	flag.Parse()

	if len(os.Args) < 2 {
		usage()
	}
	if len(os.Args) >= 2 && (*both == true){
		go rest.Run(*port)
		explorer.Run(*port+1)
	}
	switch(*mode){
		case "rest":
			rest.Run(*port)
		case "explorer":
			explorer.Run(*port)
		default:
			usage()
	}
}