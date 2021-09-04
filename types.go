package football

type Season struct {
	ID              int    `json:"id,omitempty"`
	StartDate       string `json:"startDate,omitempty"`
	EndDate         string `json:"endDate,omitempty"`
	CurrentMatchday int    `json:"currentMatchday,omitempty"`
	Winner          Winner `json:"winner,omitempty"`
}

type Standing struct {
	Stage string      `json:"stage,omitempty"`
	Type  string      `json:"type,omitempty"`
	Group interface{} `json:"group,omitempty"`
	Table []Table     `json:"table,omitempty"`
}

type Scorer struct {
	Player        Player `json:"player,omitempty"`
	Team          Team   `json:"team,omitempty"`
	NumberOfGoals int    `json:"numberOfGoals,omitempty"`
}

type Table struct {
	Position       int  `json:"position,omitempty"`
	Team           Team `json:"team,omitempty"`
	PlayedGames    int  `json:"playedGames,omitempty"`
	Won            int  `json:"won,omitempty"`
	Draw           int  `json:"draw,omitempty"`
	Lost           int  `json:"lost,omitempty"`
	Points         int  `json:"points,omitempty"`
	GoalsFor       int  `json:"goalsFor,omitempty"`
	GoalsAgainst   int  `json:"goalsAgainst,omitempty"`
	GoalDifference int  `json:"goalDifference,omitempty"`
}

type Winner struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	ShortName string `json:"shortName,omitempty"`
	TLa       string `json:"tla,omitempty"`
	CrestURL  string `json:"crestUrl,omitempty"`
}
