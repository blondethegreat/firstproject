package handlers

import (
	"fmt"
	"net/http"
	"os"
)

func Yarik(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "MACAN")
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "platosatlantis")
}

func MacCofee(w http.ResponseWriter, r *http.Request) {
	bytes, err := os.ReadFile("./assets/breaktime.jpeg")
	if err != nil {
		fmt.Println("Невозможно открыть файл:", err)
		return
	}
	w.Write(bytes)
}

func ShelbyWalk(w http.ResponseWriter, r *http.Request) {
	bytes, err := os.ReadFile("./assets/ShelbyWalk.gif")
	if err != nil {
		fmt.Println("Невозможно открыть файл:", err)
		return
	}
	w.Write(bytes)
}
