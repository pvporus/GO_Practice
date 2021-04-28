package main

func main() {
	a := App{}

	a.Initialize("booksDB", "postgres", "root")

	a.Run(":8080")
}
