package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)

	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var user usuario

	if err = json.Unmarshal(requestBody, &user); err != nil {
		w.Write([]byte("Falha ao converter usuário para struct"))
		return
	}

	db, err := banco.Conectar()

	if err != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("Insert into usuarios (nome, email) values (?, ?)")
	if err != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	insercao, err := statement.Exec(user.Nome, user.Email)
	if err != nil {
		w.Write([]byte("Erro ao inserir"))
		return
	}

	idInserido, err := insercao.LastInsertId()
	if err != nil {
		w.Write([]byte("Erro ao obter o id inserido"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuario inserido com sucesso! Id: %d", idInserido)))
}

// Traz os usuários que estão no banco
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, err := banco.Conectar()

	if err != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}

	defer db.Close()

	linhas, err := db.Query("select * from usuarios")
	if err != nil {
		w.Write([]byte("Erro ao buscar usuários"))
		return
	}

	defer linhas.Close()

	var usuarios []usuario

	for linhas.Next() {
		var usuario usuario

		if err := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); err != nil {
			w.Write([]byte("Erro ao ler usuários"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao converter usuários para json"))
		return
	}
}

// Traz um usuário específico
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}
	defer db.Close()

	linha, err := db.Query("select * from usuarios where id= ?", ID)
	if err != nil {
		w.Write([]byte("Erro ao buscar usuário"))
		return
	}

	defer linha.Close()

	var user usuario
	if linha.Next() {
		if err := linha.Scan(&user.ID, &user.Nome, &user.Email); err != nil {
			w.Write([]byte("Erro ao ler usuário"))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(user); erro != nil {
		w.Write([]byte("Erro ao converter usuário para json"))
		return
	}

}

// Atualiza um usuário existente
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	requestBody, err := io.ReadAll(r.Body)

	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var user usuario
	if err := json.Unmarshal(requestBody, &user); err != nil {
		w.Write([]byte("Falha ao converter usuário para struct"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao converter no banco de dados!"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar statement banco de dados!"))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(user.Nome, user.Email, ID); err != nil {
		w.Write([]byte("Erro ao atualizar usuário banco de dados!"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Deletar usuário deletará usuário no banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao converter no banco de dados!"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar statement banco de dados!"))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		w.Write([]byte("Erro ao deletar usuário banco de dados!"))
		return
	}

	w.WriteHeader(http.StatusOK)

}
