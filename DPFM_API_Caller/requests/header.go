package requests

type Header struct {
	Shop				int     `json:"Shop"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
