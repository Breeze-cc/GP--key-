package main

type SteamDB struct {
	Hits []SteamDBResult `json:"hits"`
}

type SteamDBResult struct {
	LastUpdated     int    `json:"lastUpdated"`
	ObjectID        string `json:"objectID"`
	HighlightResult struct {
		Name struct {
			Value            string   `json:"value"` // 名字
			MatchLevel       string   `json:"matchLevel"`
			FullyHighlighted bool     `json:"fullyHighlighted"`
			MatchedWords     []string `json:"matchedWords"`
		} `json:"name"`
	} `json:"_highlightResult"`
}
