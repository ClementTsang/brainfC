package ctobf

// Type names pulled from here: http://www.quut.com/c/ANSI-C-grammar-l-2011.html

// o represents a regex of [0-7]
type o struct{}

func (lexRegex o) getRegexValue() string {
	return "[0-7]"
}

// d represents a regex of [0-7]
type d struct{}

func (lexRegex d) getRegexValue() string {
	return "[0-9]"
}

type nz struct{}

func (lexRegex nz) getRegexValue() string {
	return "[1-9]"
}

type l struct{}

func (lexRegex l) getRegexValue() string {
	return "[a-zA-Z_]"
}

type a struct{}

func (lexRegex a) getRegexValue() string {
	return "[a-zA-Z_0-9]"
}

type h struct{}

func (lexRegex h) getRegexValue() string {
	return "[a-fA-F0-9]"
}

type hp struct{}

func (lexRegex hp) getRegexValue() string {
	return "(0[xX])"
}

type e struct{}

func (lexRegex e) getRegexValue() string {
	return "([Ee][+-]?{D}+)"
}

type p struct{}

func (lexRegex p) getRegexValue() string {
	return "([Pp][+-]?{D}+)"
}

type fs struct{}

func (lexRegex fs) getRegexValue() string {
	return "(f|F|l|L)"
}

type is struct{}

func (lexRegex is) getRegexValue() string {
	return "(((u|U)(l|L|ll|LL)?)|((l|L|ll|LL)(u|U)?))"
}

type cp struct{}

func (lexRegex cp) getRegexValue() string {
	return "(u|U|L)"
}

type sp struct{}

func (lexRegex sp) getRegexValue() string {
	return "(u8|u|U|L)"
}

type es struct{}

// TODO: may need to check this one
func (lexRegex es) getRegexValue() string {
	return "(\\\\(['\"\\?\\\\abfnrtv]|[0-7]{1,3}|x[a-fA-F0-9]+))"
}

type ws struct{}

func (lexRegex ws) getRegexValue() string {
	return "[ \\t\\v\\n\\f]"
}
