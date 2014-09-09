Animation
	Import string
	Images []AnimationSlide

AnimationSlide
	Interval int
	Image    Image

AnimLayer
	Point     Point
	Animation string # Animation import

Point
	X int
	Y int

Size
	Width  int
	Height int

Bounds
	X      int
	Y      int
	Width  int
	Height int

SaveVariables
	Vars map[string]unknown

SpriteSheetImage
	SrcBounds Bounds

SpriteSheet
	TilesWide         int
	TilesHigh         int
	TileWidth         int
	TileHeight        int
	SrcFile           string
	Images            map[int]SpriteSheetImage
	ImageIdIterator   int
	RecycledImageIds  []int

Image
	# Path to the sprite sheet the image is from
	SpriteSheet string
	ImgId       int
	DefaultSize Size

Interaction
	Type int # (Text = 0, Script = 1)
	Next unknown

SpriteClass
	AnimLayers [][]AnimLayer
	SpriteType int # (Inanimate = 1, Person = 2, Creature = 3)
	Attributes string # content differs depending on SpriteType

Sprite
	Id          string
	SpriteClass string
	Motion      int # (Still = 0, Walking = 1)
	Facing      int # (North = 0, South = 1, East = 2, West = 3)
	SpriteType  int # (Inanimate = 0, Person = 1, Creature = 2)
	Interaction Interaction
	Data        string # content differs depending on SpriteType

TileClass
	Import      string
	TerrainType int
	LowerAnim   AnimLayer
	UpperAnim   AnimLayer

Tile
	TileClass TileClass
	Occupant  Sprite

World
	Zones []ZoneInstance

Zone
	Tiles [][][]Tile

ZoneHeader
	Zone      string
	TilesWide int
	TilesHigh int
	Layers    int

ZoneInstance
	AccessorID string
	ZoneHeader string
	Address    Point

CreatureMove
	Name              map[string]string
	Type              string
	Power             int
	RequiresRegarge   bool
	Script            string
	Burn              StatusEffect
	Freeze            StatusEffect
	Paralyze          StatusEffect
	Poison            StatusEffect
	Sleep             StatusEffect
	WorldAbilityFlags int # 1 = Swim, 2 = Whirlpool, 4 = Waterfall, 8 = Cut, 16 = Light

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
	CreatureMove string
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

StatusEffect
	AttackerEffect int
	EnemyEffect    int

PersonClass
	Import    string
	Title     map[string]string
	# [x][y]: x - (North/South/East/West), y - (Still = 0, Walking = 1)
	Overhead  [][][]AnimLayer
	FrontView Animation
	BackView  Animation

Person
	PersonClass PersonClass
	Name        map[string]string
	Creatures   []string # paths to the Creatures this Person owns

Font
	TtfPath   string
