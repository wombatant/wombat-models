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

CreatureMove
	Name            map[string]string
	Type            string
	Power           int
	RequiresRegarge bool
	Script          string
	Burn            StatusEffect
	Freeze          StatusEffect
	Paralyze        StatusEffect
	Poison          StatusEffect
	Sleep           StatusEffect

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

StatusEffect
	AttackerEffect int
	EnemyEffect    int

Fraction
	Current   int
	Available int

ModelFile
	Type string

PersonClass
	ID        int
	Name      map[string]string
	Creatures []int
	Overhead  []Animation
	FrontView Animation
	BackView  Animation

Person
	PersonClass PersonClass

User
	PersonID int
	World    string
	Zone     string
	XAddress int
	YAddress int
	Layer    int

SaveFile
	Vars map[string]unknown
	User User
