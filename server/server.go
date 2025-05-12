package server

import (
	"encoding/json"
	"fmt"
	"industrialApplication/database"
	"io/ioutil"
	"net/http"
)

/*
This struct is related with the data that exist on the database
*/
type machines struct {
	Id                    uint32 `json:"id"`
	Name                  string `json:"name"`
	Brand                 string `json:"brand"`
	Description           string `json:"description"`
	Serial_number         uint64 `json:"serial_number"`
	Installation_location string `json:"installation_location"`
}

/*
	On the next steps we will create the functions related with the methods:
		- POST (CREATE AN DATA)
		- GET (READ DATA FROM THE DATABASE)
		- PUT (UPDATE THE DATA)
		- DELETE (DELETE THE DATA)
*/

// POST
func CreateMachine(w http.ResponseWriter, r *http.Request) {
	/*
		When we call the function 'CreateMachine' through http.HandleFunc(),
		it is necessary to pass two parameters: 'w' and 'r'.

		r: *http.Request
		- A pointer to a struct that contains everything about the HTTP request.
		- It includes:
			• Method (GET, POST, etc.)
			• Headers (ex: Content-Type)
			• URL parameters and query strings
			• Body (data sent by the client, usually in JSON)
			• Cookies, form data, etc.
		- You use 'r' to **read** what the client sent.

		w: http.ResponseWriter
		- An interface used to **write** a response back to the client.
		- It allows you to:
			• Set HTTP status code (e.g., 200, 201, 404)
			• Set response headers
			• Write content (text, JSON, HTML, etc.)
		- Example: w.Write([]byte("Created"))

		Why are they passed as parameters?
		- Because the Go HTTP server (net/http) automatically provides them
		  when it calls your handler.
		- Each HTTP request gets its own 'w' and 'r'.
		- This ensures concurrent safety and keeps handlers stateless and modular.
	*/

	fmt.Println("Creating the user...")

	/*
		r.Body:
		- Represents the **body of the HTTP request** — it's where the client sends data (e.g., JSON, form data).
		- It is of type `io.ReadCloser`, meaning it behaves like a stream that must be read and then closed.
		- Common in POST or PUT requests, where the client sends structured data in the request body.

		io.ReadAll(r.Body):
		- Reads the entire content of r.Body and returns it as a byte slice ([]byte).
		- This allows you to later convert it to a string or parse it as JSON.
		- The result is stored in 'bodyRequest'.
		- 'erro' captures any error that might occur during reading (e.g., connection closed, malformed request).
	*/
	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler corpo da requisição"))
		return
	}

	var machines machines

	/*
		The json.Unmarshal function is used to convert JSON data (received as []byte)
		into a Go struct.

		It receives two parameters:
		1. bodyRequest → the raw JSON data, usually read from the HTTP request body
		2. &machine → a pointer to the struct where the data will be stored

		Why a pointer (&machine)?
		- json.Unmarshal needs to modify the original variable and store the parsed data inside it.
		- Using a pointer allows json.Unmarshal to access the memory location and fill the struct fields.

		If an error occurs during the conversion (e.g., invalid JSON or mismatch in field types),
		it is captured in the 'erro' variable, and an error message is returned to the client.
	*/
	if erro = json.Unmarshal(bodyRequest, &machines); erro != nil {
		w.Write([]byte("Erro ao converter o dado para struct"))
		return
	}

	// Opening the connection with the database
	db, erro := database.Connection()
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}

	// The last function to be called after all the code has been executed
	defer db.Close()

	/*
		db.Prepare is used to prepare an SQL instruction to improve performance and security.
		It allows reusing the same SQL statement with different values, and helps prevent SQL injection.
	*/
	statement, erro := db.Prepare("INSERT INTO machines (name, brand, description, serial_number, installation_location) VALUES (?, ?, ?, ?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar statement!"))
		return
	}

	defer statement.Close()

	/*
		statement is responsible for preparing the SQL instruction with placeholders (the ? symbols),
		which allows us to insert values securely and efficiently into the database.

		And, when we use .Exec, it represents the execution of that prepared instruction, where the actual
		values are passed as arguments and bound to the placeholders defined in the SQL.

		.Exec will be on the variable 'insert', which stores the result of the execution. This result can be
		used to retrieve information like the number of affected rows or the ID of the last inserted record.
	*/
	insert, erro := statement.Exec(machines.Name, machines.Brand, machines.Description, machines.Serial_number, machines.Installation_location)
	if erro != nil {
		w.Write([]byte("Erro ao inserir os dados no banco!"))
		return
	}

	/*
		idInsert is responsible for storing the ID generated by the database after the insert operation.
		This ID is usually the value of an auto-increment primary key.
	*/
	idInsert, erro := insert.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o ID Inserido"))
	}

	fmt.Println(machines)
	fmt.Println(machines.Name)
	fmt.Println(machines.Brand)
	fmt.Println(machines.Description)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! ID: %d", idInsert)))
	fmt.Println("Usuario criado com sucesso...")

}
