package handlers

import (
	"fmt"
	"net/http"

	"example.com/m/controllers"
)

func DigiviceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "Error to parse form: %v", err)
			return
		}
		controllers.GetDigiByName(r.PostForm.Get("name"))
		// if err != nil {
		// 	fmt.Fprintf(w, "Error to get Digimon: %v", err)
		// 	return
		// }
	}
	http.ServeFile(w, r, "handlers/templates/digivice.html")
}
