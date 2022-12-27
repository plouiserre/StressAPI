package configuration

type Configuration struct {
	StoreFolder string     `json:"StoreFolder"`
	Workflows   []Workflow `json:"Workflows"`
}

type Workflow struct {
	Id       int       `json:"Id"`
	Requests []Request `json:"Requests"`
}

type Request struct {
	Id         int      `json:"Id"`
	Verb       string   `json:"Verb"`
	Uri        string   `json:"Uri"`
	Parameters []string `json:"Parameters"`
	Body       string   `json:"Body"`
	Times      int      `json:"Times"`
}
