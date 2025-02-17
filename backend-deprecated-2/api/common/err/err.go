package err

// contains common error messages and useful funcs

type Error struct {
	Error string `json:"error"`
}

//FIXME: i dont like plural for strings...
type Errors struct {
	Errors []string `json:"errors"`
}
