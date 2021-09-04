package football

type Season struct {
	ID              int      `json:"id,omitempty"`
	StartDate       string   `json:"startDate,omitempty"`
	EndDate         string   `json:"endDate,omitempty"`
	CurrentMatchday int      `json:"currentMatchday,omitempty"`
	Winner          Winner   `json:"winner,omitempty"`
	AvailableStages []string `json:"availableStages,omitempty"`
}

type Standing struct {
	Stage string      `json:"stage,omitempty"`
	Type  string      `json:"type,omitempty"`
	Group interface{} `json:"group,omitempty"`
	Table []Table     `json:"table,omitempty"`
}

type StandingType string

const (
	StandingTypeTotal StandingType = "TOTAL"
	StandingTypeHome  StandingType = "HOME"
	StandingTypeAway  StandingType = "AWAY"
)

type Status string

const (
	StatusScheduled Status = "SCHEDULED"
	StatusLive      Status = "LIVE"
	StatusInPlay    Status = "IN_PLAY"
	StatusPaused    Status = "PAUSED"
	StatusFinished  Status = "FINISHED"
	StatusPostPoned Status = "POSTPONED"
	StatusSuspended Status = "SUSPENDED"
	StatusCanceled  Status = "CANCELED"
)

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

type Coach struct {
	ID             int    `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	CountryOfBirth string `json:"countryOfBirth,omitempty"`
	Nationality    string `json:"nationality,omitempty"`
}

type Time struct {
	HomeTeam int `json:"homeTeam,omitempty"`
	AwayTeam int `json:"awayTeam,omitempty"`
}

type Score struct {
	Winner    string `json:"winner,omitempty"`
	Duration  string `json:"duration,omitempty"`
	FullTime  Time   `json:"fullTime,omitempty"`
	HalfTime  Time   `json:"halfTime,omitempty"`
	ExtraTime Time   `json:"extraTime,omitempty"`
	Penalties Time   `json:"penalties,omitempty"`
}
type Goals struct {
	Minute    int         `json:"minute,omitempty"`
	ExtraTime interface{} `json:"extraTime,omitempty"`
	Type      string      `json:"type,omitempty"`
	Team      Team        `json:"team,omitempty"`
	Scorer    Player      `json:"scorer,omitempty"`
	Assist    Player      `json:"assist,omitempty"`
}

type Bookings struct {
	Minute int    `json:"minute,omitempty"`
	Team   Team   `json:"team,omitempty"`
	Player Player `json:"player,omitempty"`
	Card   string `json:"card,omitempty"`
}

type Substitutions struct {
	Minute    int    `json:"minute,omitempty"`
	Team      Team   `json:"team,omitempty"`
	PlayerOut Player `json:"playerOut,omitempty"`
	PlayerIn  Player `json:"playerIn,omitempty"`
}
type Referees struct {
	ID          int         `json:"id,omitempty"`
	Name        string      `json:"name,omitempty"`
	Nationality interface{} `json:"nationality,omitempty"`
}
type Head2Head struct {
	NumberOfMatches int       `json:"numberOfMatches,omitempty"`
	TotalGoals      int       `json:"totalGoals,omitempty"`
	HomeTeam        TeamStats `json:"homeTeam,omitempty"`
	AwayTeam        TeamStats `json:"awayTeam,omitempty"`
}

type TeamStats struct {
	Wins   int `json:"wins,omitempty"`
	Draws  int `json:"draws,omitempty"`
	Losses int `json:"losses,omitempty"`
}
