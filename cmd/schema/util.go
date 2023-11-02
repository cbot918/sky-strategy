package main

type Stats struct {
	Int   int `json:"int"`
	Rule  int `json:"rule"`
	Force int `json:"force"`
	Charm int `json:"charm"`
}

type EQs struct {
	Head   string `json:"head"`
	Weapon string `json:"weapon"`
	Armor  string `json:"armor"`
	Pants  string `json:"pants"`
	Boots  string `json:"boots"`
}

type General struct {
	Userid int    `json:"user_id"`
	Name   string `json:"name"`
	Level  int    `json:"3"`
	Stats  Stats

	SKPosition int    `json:"sk_position"`
	Growth     string `json:"growth"`
	EQs        EQs
	Pets       string `json:"pets"`
	Talent     []int  `json:"talent"`
}

func newGeneral() General {
	return General{
		Userid: 1,
		Name:   "小小將",
		Level:  3,
		Stats: Stats{
			Int:   1,
			Rule:  1,
			Force: 1,
			Charm: 1,
		},
		SKPosition: 2,
		Growth:     "A",
		EQs: EQs{
			Head:   "h001",
			Weapon: "w001",
			Armor:  "a001",
			Pants:  "p001",
			Boots:  "b001",
		},
		Pets:   "none",
		Talent: []int{1, 2, 3},
	}
}
