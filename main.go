package main

import "net/http"

func main() {
	http.HandleFunc("/", HandleIndex)
	http.ListenAndServe(":10000", nil)
}

// HandleIndex renders the home page using our RenderIndex ego template.
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	// Our list of available widgets.
	var widgets = []*Widget{
		{Name: "Blue Widget", Price: 100},
		{Name: "Red Widget", Price: 20},
	}

	// Generate the template and write it to the response.
	RenderIndex(w, widgets)
}

type Widget struct {
	Name  string
	Price int
}
