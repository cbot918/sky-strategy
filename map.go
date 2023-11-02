package main

type Map struct {
	Block []Block
}

type Block struct {
	Occupy       bool
	BuildingType string
	// user / bandit / beast / soldier / castle
	BuildingName string
}

func GetMap() [100]Block {

	// 2: 殤
	// 4: soldier
	// 13: 洛陽
	// 24: beast

	blocks := [100]Block{}

	blocks[23] = Block{
		Occupy:       true,
		BuildingType: "user",
		BuildingName: "公主",
	}

	blocks[24] = Block{
		Occupy:       true,
		BuildingType: "user",
		BuildingName: "天之殤",
	}

	blocks[57] = Block{
		Occupy:       true,
		BuildingType: "npc",
		BuildingName: "傭兵營",
	}

	blocks[32] = Block{
		Occupy:       true,
		BuildingType: "city",
		BuildingName: "洛陽",
	}

	blocks[44] = Block{
		Occupy:       true,
		BuildingType: "npc",
		BuildingName: "獸穴",
	}

	return blocks

}
