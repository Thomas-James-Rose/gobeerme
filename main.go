package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/thomas-james-rose/gobeerme/truerand"
)

type Beer struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Tagline     string `json:"tagline"`
	FirstBrewed string `json:"first_brewed"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}

func beerHandler(w http.ResponseWriter, r *http.Request) {
	beerID := truerand.Int(1, 325)
	resp, _ := http.Get("https://api.punkapi.com/v2/beers/" + strconv.Itoa(beerID))

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		var beers []Beer
		err = json.Unmarshal(bodyBytes, &beers)

		beer := beers[0]

		fmt.Println("Serving a nice cold", beer.Name)

		fmt.Fprintf(w, "<h1>"+beer.Name+"<h1>")
		fmt.Fprintf(w, "<h2>"+beer.Tagline+"<h2>")
		fmt.Fprintf(w, "<p>First brewed in "+beer.FirstBrewed+"</p>")
		fmt.Fprintf(w, "<p>"+beer.Description+"</p>")
		fmt.Fprintf(w, "<img src="+beer.ImageURL+">")
	}
}

func main() {
	http.HandleFunc("/beer", beerHandler)

	fmt.Println("Serving on port 3000...")
	http.ListenAndServe(":3000", nil)
}
