package util

import "fmt"

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
	for _, r := range pofix {
		switch r {
		case '.':
			frag2 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			frag1 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			frag1.accept.edge1 = frag2.initial

			nfaStack = append(nfaStack, &nfa{initial: frag1.initial, accept: frag2.accept})
		case '|':

		case '*':

		default:

		}

	}
	//handles errors
	if len(nfaStack) != 1 {
		fmt.Println("Sorry more than 1 nfa found", len(nfaStack), nfaStack)
	}
	return nfaStack[0]
}

/////////////////////////////MATCHING/////////////////////

func pomatch(po string, s string) bool {

	return false
}

func addState(l []*state, s *state, a *state) []*state {

	l = append(l, s)

	return l
}
