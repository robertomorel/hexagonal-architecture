package handler

import "encoding/json"

/*
	Vai receber uma string e vai retornar um slice de byte, que é o formato padrão quando trabalhamos com json
*/
func jsonError(msg string) []byte {
	error := struct {
		Message string `json:"message"`
	}{
		msg,
	}
	// Conversando para json
	r, err := json.Marshal(error)
	if err != nil {
		return []byte(err.Error())
	}
	return r
}
