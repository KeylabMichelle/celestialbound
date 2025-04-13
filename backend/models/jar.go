package models

type Jar struct {
	JarLevel       int `json:"jar_level"`        // Current level of the jar (0-10)
	StarsStored    int `json:"stars_stored"`     // Stars stored in the jar
	StarsPerSecond int `json:"stars_per_second"` // Stars gained per second (Passive action)
	UpgradeCost    int `json:"upgrade_cost"`     // Cost for next upgrade of the jar
	MaxCapacity    int `json:"max_capacity"`     // Maximum capacity of the jar

}
