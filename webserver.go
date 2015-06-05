// Package to handle and route other web request as a server
package main

import (
	"github.com/eyetoe/jabberwebby"
	"github.com/eyetoe/landing"
	"github.com/eyetoe/style"
	"github.com/eyetoe/technotes"
	"net/http"
)

func main() {
	// Handlers
	http.HandleFunc("/", landing.Handler)
	http.HandleFunc("/technotes", technotes.Handler)
	http.HandleFunc("/jabberwebby", jabberwebby.Handler)
	http.HandleFunc("/rpg", jabberwebby.Handler)
	// CSS
	http.HandleFunc("/style.css", style.Default)
	http.HandleFunc("/landing.css", style.Landing)
	http.HandleFunc("/technotes.css", style.Technotes)
	http.HandleFunc("/jabberwebby.css", style.Jabberwebby)
	http.ListenAndServe(":80", nil)
}
