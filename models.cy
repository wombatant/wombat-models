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

Sprite
	AnimLayers [][]AnimLayer
	SpriteType int
	PersonID   int
	Speed      int
	Name       string
	Path       string
	ScriptPath string

Tile
	Import       string
	TerrainType  int
	LowerAnim    AnimLayer
	UpperAnim    AnimLayer

TileInstance
	Tile      Tile
	Occupant  Sprite

World
	Zones []ZoneInstance

Zone
	Tiles [][][]TileInstance

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
	Overhead  []Animation
	FrontView Animation
	BackView  Animation

Person
	PersonClass PersonClass
