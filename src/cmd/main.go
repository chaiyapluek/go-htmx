package main

const portNumber = "localhost:8080"

func main() {
	e := NewRouter()
	e.Logger.Fatal(e.Start(portNumber))
}
