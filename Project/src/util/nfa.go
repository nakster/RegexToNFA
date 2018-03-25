package util

type state struct {
	symbol rune
	edge1  *state
	edge2  *state
}

//this is the helper state
type nfa struct {
	initial *state
	accept  *state
}

//returns a pointer to nfa
func ReturnNFA(pofix string) *nfa {

	nfaStack := []*nfa{}

	return nfaStack[0]
}
