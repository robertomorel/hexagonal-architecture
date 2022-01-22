package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codeedu/go-hexagonal/adapters/web/handler"
	"github.com/codeedu/go-hexagonal/application"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

// Criando um novo tipo WebServer
type Webserver struct {
	Service application.ProductServiceInterface
}

func MakeNewWebserver() *Webserver {
	return &Webserver{}
}

/*
	Quando esse método for executado, ele vai começar a servir o WebServer
*/
func (w Webserver) Serve() {
	// O "mux" trabalhar como um roteador. Tratamento de rotas
	r := mux.NewRouter()
	// O "negroni" trabalha como um middleware de logs
	n := negroni.New(
		negroni.NewLogger(),
	)
	handler.MakeProductHandlers(r, n, w.Service)
	http.Handle("/", r)
	// Configurações padrão do server
	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":9000",
		Handler:           http.DefaultServeMux,                        // Vamos trabalhar com o handler do Mux
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile), // Como iremos trabalhar com o log
	}
	// Subindo o servidor
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
