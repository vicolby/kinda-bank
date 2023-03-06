package main

func main() {
	store, err := NewPostgresStorage()
	if err != nil {
		panic(err)
	}

	server := NewAPIServer(":8080", store)
	server.Start()
}
