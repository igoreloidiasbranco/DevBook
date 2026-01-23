package controllers

import (
	"api/autenticacao"
	"api/banco"
	"api/models"
	"api/repositorios"
	"api/respostas"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func CriarPublicacao(w http.ResponseWriter, r *http.Request) {

	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publicacao models.Publicacao
	if erro = json.Unmarshal(corpoRequest, &publicacao); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	publicacao.AutorID = usuarioID

	if erro = publicacao.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.ConectarBanco()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.RepositorioDePublicacoes(db)
	publicacao.CriadaEm = time.Now()
	publicacao.ID, erro = repositorio.Criar(publicacao)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, publicacao)
}

func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {

}

func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {

}

func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {

}

func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {

}
