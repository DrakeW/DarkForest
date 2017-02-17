package utils

// ########## Civilizaitons Constants #################

const (
	LIGHTSPEED_x0001 int = 1 + iota
	LIGHTSPEED_x001
	LIGHTSPEED_x1
	LIGHTSPEED_x2
	LIGHTSPEED_x10
)
const (
	CONQUERER    string = "conquerer"
	CONSERVATIVE string = "conservative"
)

const (
	LIGHTSPEED float64 = 1
)

// ########## Universe Constants #################

// assume a fixed size universe
const (
	WIDTH  int64 = 10e+5 // unit: lightyear
	HEIGHT int64 = 10e+5 // 10e6
)

const (
	TOTAL_MATTER float64 = 10e+5
)

// ########## Gaming Window ####################
const (
	G_WIDTH  int     = 150
	G_HEIGHT int     = 100
	G_SCALE  float64 = 5
)
