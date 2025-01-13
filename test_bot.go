package main

import (
	"github.com/JackHumphries9/dapper-go/server"
)

const PUBLIC_KEY = "e9573621727df5e8b915f2a52b481d262f0d26e6d429913e1b960062ca6d4ab3"

func main() {
	server := server.NewInteractionServer(PUBLIC_KEY)

	server.Listen(3000)
}
