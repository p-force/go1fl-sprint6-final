package main

<<<<<<< HEAD
import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, `http-server: `, log.LstdFlags)

	serv := server.CreateServer(logger)
	err := serv.ListenAndServe()

	if err != nil {
		logger.Fatal(err)
	}
=======
func main() {

>>>>>>> 451cb3d (Initial commit)
}
