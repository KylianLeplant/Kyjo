package game

const (
	// Card values range from -2 to 12, inclusive.
	MinCardValue    = -2
	MaxCardValue    = 12
	HiddenCardValue = 13
)

func isValidCardValue(value int) bool {
	return value >= MinCardValue && value <= MaxCardValue
}
