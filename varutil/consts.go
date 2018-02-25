package varutil

const (
	// ByteMask is mask to filter one byte with '&' and operator
	ByteMask = 0xFF
	// WordMask is mask to filter one word with '&' and operator
	WordMask = 0xFFFF
	// DWordMask is mask to filter one double-word with '&' and operator
	DWordMask = 0xFFFFFFFF
	// LongMask is mask to filter one long with '&' and operator
	LongMask = 0xFFFFFFFFFFFFFFFF

	// ThreeByteMask is mask to filter three bytes with '&' and operator
	ThreeByteMask = 0xFFFFFF
	// FiveByteMask is mask to filter five bytes with '&' and operator
	FiveByteMask = 0xFFFFFFFFFF
	// SixByteMask is mask to filter six bytes with '&' and operator
	SixByteMask = 0xFFFFFFFFFFFF
	// SevenByteMask is mask to filter seven bytes with '&' and operator
	SevenByteMask = 0xFFFFFFFFFFFFFF
)
