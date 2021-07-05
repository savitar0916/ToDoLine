package models

type List struct {
	Name     string
	Title    string
	DateTime string
	Done     bool
}

func main() {
	collection := client.Database("test").Collection("trainers")
}
