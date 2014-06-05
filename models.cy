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

SpriteClass
	AnimLayers [][]AnimLayer
	SpriteType int # (Inanimate = 0, Person = 1, Creature = 2)
	Attributes string # content differs depending on SpriteType

Sprite
	SpriteClass string

TileClass
	Import       string
	TerrainType  int
	LowerAnim    AnimLayer
	UpperAnim    AnimLayer

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

PersonClass
	Import    string
	Name      map[string]string
	Creatures []string # paths to the Creatures this Person owns
	Overhead  [][]AnimLayer
	FrontView Animation
	BackView  Animation

Person
	PersonClass PersonClass
