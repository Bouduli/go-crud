package types

type ProblemJson struct {
	Type   string `json:"type"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Status int    `json:"status"`

	Context map[string]any `json:"context"`
}
