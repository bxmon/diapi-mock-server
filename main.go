package main

func main() {
	server := SetUpEngine()
	server.Run(":5555")
}
