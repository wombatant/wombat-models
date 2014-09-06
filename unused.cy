# these models are not being used in the current code

CreatureType
	Name          map[string]string
	Special       bool
	StrongAgainst []string
	WeakAgainst   []string

CreatureClass
	Name          map[string]string
	Successor     string
	Predecessor   string
	Types         []string
	CanLearn      []string
	LearnsAtLevel map[int]string
	FrontView     Animation
	BackView      Animation

CreatureMoveInstance
	CreatureMove int
	PP           Fraction

Creature
	ID            int
	Name          map[string]string
	CreatureClass string
	Male          bool
	Level         int
	Health        Fraction
	Attack        int
	SpecAttack    int
	Defense       int
	SpecDefense   int
	Burned        bool
	Frozen        bool
	Poisoned      bool
	Asleep        bool
	Moves         []CreatureMoveInstance

Fraction
	Current   int
	Available int

ModelFile
	Type string
