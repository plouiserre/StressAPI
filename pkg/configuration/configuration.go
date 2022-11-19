package configuration

type Configuration struct {
	Verb       string   `json:"Verb"`
	Uri        string   `json:"Uri"`
	Parameters []string `json:"Parameters"`
	Body       string   `json:"Body"`
}
