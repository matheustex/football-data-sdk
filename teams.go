package football

type TeamsService service

type Team struct {
	ID          int      `json:"id,omitempty"`
	Area        Area     `json:"area,omitempty"`
	Name        string   `json:"name,omitempty"`
	ShortName   string   `json:"shortName,omitempty"`
	Tla         string   `json:"tla,omitempty"`
	CrestURL    string   `json:"crestUrl,omitempty"`
	Address     string   `json:"address,omitempty"`
	Phone       string   `json:"phone,omitempty"`
	Website     string   `json:"website,omitempty"`
	Email       string   `json:"email,omitempty"`
	Founded     int      `json:"founded,omitempty"`
	ClubColors  string   `json:"clubColors,omitempty"`
	Venue       string   `json:"venue,omitempty"`
	Coach       Coach    `json:"coach"`
	Captain     Player   `json:"captain"`
	Lineup      []Player `json:"lineup"`
	Bench       []Player `json:"bench"`
	LastUpdated string   `json:"lastUpdated,omitempty"`
}
