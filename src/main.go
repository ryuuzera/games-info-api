package main

import (
	"games-info-api/src/server"
)

func main() {
		app := server.New();

		app.Run();
};