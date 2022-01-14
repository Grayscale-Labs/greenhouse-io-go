package greenhouseio

type QueryBuilder interface {
	Len() int
	WriteString(string) (int, error)
}

// addPrefixToken adds the correct preceding token given how many parameters exist.
func addPrefixToken(builder QueryBuilder) {
	if builder.Len() == 0 {
		builder.WriteString("?")
	} else {
		builder.WriteString("&")
	}
}
