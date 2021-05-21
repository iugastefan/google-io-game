package main

type ArenaUpdate struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Arena Arena `json:"arena"`
}

type Arena struct {
	Dimensions []int                  `json:"dims"`
	State      map[string]PlayerState `json:"state"`
}

type Dir string

const (
	N Dir = "N"
	W Dir = "W"
	S Dir = "S"
	E Dir = "E"
)

type PlayerState struct {
	X         int `json:"x"`
	Y         int `json:"y"`
	Direction Dir `json:"direction"`
	// WasHit    bool   `json:"wasHit"`
	Score int `json:"score"`
}
