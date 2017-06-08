package gio

func isWhiteChar(c byte) bool {
	return c == ' ' || c == '\r' || c == '\n' || c == '\t'
}

func isNewLine(c byte) bool {
	return c == '\r' || c == '\n'
}
