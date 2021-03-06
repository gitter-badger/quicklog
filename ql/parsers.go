package ql

// Parser defines the interface for parsing incoming data
type Parser interface {
	// Parse parses the incoming buffer into the line object
	Parse(buffer []byte, line *Line, config map[string]interface{}) error
}

var (
	parsers map[string]Parser
)

func init() {
	parsers = make(map[string]Parser)
}

// GetParser returns the parser registered under the name
func GetParser(name string) Parser {
	return parsers[name]
}

// RegisterParser registers a new parser
func RegisterParser(name string, parser Parser) {
	parsers[name] = parser
}
