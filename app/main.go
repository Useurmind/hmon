package main


func main() {
	address := "127.0.0.1"
	port := "8080"

	server := Server{
		Address: address,
		Port: port,
	}

	err := server.Run()
	if err != nil {
		panic(err)
	}
}