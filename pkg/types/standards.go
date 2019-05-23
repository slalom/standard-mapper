package types

// Standard defines type and revion of a Standard along with it's elements
type Standard struct {
	Type     string `json:"type"`
	Rev      string `json:"rev"`
	Children []Children
	Links    []interface{} `json:"links"`
}

// Children represent the elements (body) of a Standard
type Children struct {
	ID              string   `json:"id"`
	Section         string   `json:"section"`
	Body            string   `json:"body"`
	ComplianceLevel int      `json:"compliance_level"`
	Links           []string `json:"links"`
	Children        []Children
}
