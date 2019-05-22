package types

// Gdpr defines the structure for assets/gdpr-en.json
type Gdpr struct {
	Division []struct {
		Title struct {
			TI struct {
				P struct {
					HT struct {
						Type string `json:"Type"`
						Text string `json:"Text"`
					} `json:"HT"`
				} `json:"P"`
			} `json:"TI"`
			STI struct {
				P struct {
					HT struct {
						Type string `json:"Type"`
						HT   struct {
							Type string `json:"Type"`
							Text string `json:"Text"`
						} `json:"HT"`
					} `json:"HT"`
				} `json:"P"`
			} `json:"STI"`
		} `json:"Title"`
		ARTICLE []Article
	} `json:"Division,omitempty"`
}

type QuotStart struct {
	Code   string `json:"Code,omitempty"`
	ID     string `json:"ID,omitempty"`
	RefEnd string `json:"RefEnd,omitempty"`
}

type QuotEnd struct {
	Code     string `json:"Code,omitempty"`
	ID       string `json:"ID,omitempty"`
	RefStart string `json:"RefStart,omitempty"`
}

type Txt struct {
	QuotStart QuotStart `json:"QuotStart,omitempty"`
	QuotEnd   QuotEnd   `json:"QuotEnd,omitempty"`
	Text      string    `json:"Text,omitempty"`
}

type Article struct {
	Identifier string `json:"Identifier,omitempty"`
	TiArt      string `json:"TiArt,omitempty"`
	StiArt     string `json:"StiArt,omitempty"`
	Parag      []struct {
		Identifier string    `json:"Identifier,omitempty"`
		NoParag    string    `json:"NoParag,omitempty"`
		Text       string    `json:"Text,omitempty"`
		TextLines  []string  `json:"TextLines,omitempty"`
		Alinea     Alinea    `json:"Alinea,omitempty"`
		Elements   []Alinea  `json:"Elements,omitempty"`
		QuotStart  QuotStart `json:"QuotStart,omitempty"`
		QuotEnd    QuotEnd   `json:"QuotEnd,omitempty"`
	} `json:"Parag,omitempty"`
}

type Alinea struct {
	Text string `json:"Text,omitempty"`
	List struct {
		Type string `json:"Type,omitempty"`
		Item []struct {
			NP struct {
				NOP  string `json:"NOP,omitempty"`
				TXT  Txt    `json:"TXT,omitempty"`
				Text string `json:"Text,omitempty"`
				P    struct {
					List struct {
						Type string `json:"Type,omitempty"`
						Item []struct {
							NP struct {
								NOP  string `json:"NOP,omitempty"`
								TXT  string `json:"TXT,omitempty"`
								Text string `json:"Text,omitempty"`
							} `json:"NP,omitempty"`
						} `json:"Item,omitempty"`
					} `json:"List,omitempty"`
				} `json:"P,omitempty"`
			} `json:"NP,omitempty"`
		} `json:"Item,omitempty"`
	} `json:"List,omitempty"`
}
