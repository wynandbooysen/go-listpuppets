package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/gorilla/mux"
)

func listpuppets() string {
	//bash command excludes header and puppet master duplicate
  //use absolute path otherwise no data will be returned when running under Systemd
	out, err := exec.Command("bash", "-c", "/opt/puppetlabs/bin/puppetserver ca list --all | awk 'NR>2{print $1}'").Output() 

	if err != nil {
		fmt.Printf("%s", err)
	}

	//fmt.Println("Command Successfully Executed")
	output := string(out[:])
	//fmt.Println(output)

	servers := strings.Fields(output)
	list := "["
	for i, server := range servers {
		fmt.Println(i, " => ", server)
		list = list + " {\"Server\": \"" + server + "\"},"
	}
	list = strings.TrimSuffix(list, ",")
	list = list + "]"
	return list

}

func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello")
	fmt.Println("Endpoint Hit: homePage")
}

func getPuppets(w http.ResponseWriter, r *http.Request) {
	list := listpuppets()
	fmt.Fprintf(w, list)
	fmt.Println("Endpoint Hit: getPuppets")
}

func main() {
	listpuppets()
	router := mux.NewRouter()
	router.HandleFunc("/", homePage)
	router.HandleFunc("/listpuppets", getPuppets)
	log.Fatal(http.ListenAndServe(":8000", router))
}
