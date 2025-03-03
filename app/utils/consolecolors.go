package utils

const (
	colorKey = "\x1b["

	Reset      = colorKey + "0m"
	Bright     = colorKey + "1m"
	Dim        = colorKey + "2m"
	Underscore = colorKey + "4m"
	Blink      = colorKey + "5m"
	Reverse    = colorKey + "7m"
	Hidden     = colorKey + "8m"

	FBlack   = colorKey + "30m"
	FRed     = colorKey + "31m"
	FGreen   = colorKey + "32m"
	FYellow  = colorKey + "33m"
	FBlue    = colorKey + "34m"
	FMagenta = colorKey + "35m"
	FCyan    = colorKey + "36m"
	FWhite   = colorKey + "37m"
	FGray    = colorKey + "90m"

	BBlack   = colorKey + "40m"
	BRed     = colorKey + "41m"
	BGreen   = colorKey + "42m"
	BYellow  = colorKey + "43m"
	BBlue    = colorKey + "44m"
	BMagenta = colorKey + "45m"
	BCyan    = colorKey + "46m"
	BWhite   = colorKey + "47m"
	BGray    = colorKey + "100m"
)
