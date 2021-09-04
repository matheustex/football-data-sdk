package football

type AreasService service

type Area struct {
	ID           int    `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	CountryCode  string `json:"countryCode,omitempty"`
	ParentAreaID int    `json:"parentAreaId,omitempty"`
	ParentArea   string `json:"parentArea,omitempty"`
}
