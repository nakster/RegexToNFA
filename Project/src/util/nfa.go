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
			//pop everything except the last one
			frag2 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			frag1 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]
			//join states
			frag1.accept.edge1 = frag2.initial
			//append the changes(states)
			nfaStack = append(nfaStack, &nfa{initial: frag1.initial, accept: frag2.accept})
		case '|':
			//pop everything except the last one
			frag2 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			frag1 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]
			//join states
			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			accept := state{}

			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept
			//append the changes(states)
			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})
		case '*':
			//pop everything except the last one
			frag := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]
			//join states
			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}

			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})
		case '+':
			//pop everything except the last one
			frag := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]
			//join states
			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}
			//append the changes(states)
			frag.accept.edge1 = &initial
			frag.accept.edge2 = &accept

			nfaStack = append(nfaStack, &nfa{initial: frag.initial, accept: &accept})
		case '?':
			//pop everything except the last one
			frag := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]
			//join states
			initial := state{edge1: frag.initial, edge2: frag.accept}
			//append the changes(states)
			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: frag.accept})
		default:
			accept := state{}
			initial := state{symbol: r, edge1: &accept}
			//append the changes(states)
			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})
		}
	}
	//handles errors
	if len(nfaStack) != 1 {
		fmt.Println("Sorry more than 1 nfa found", len(nfaStack), nfaStack)
	}
	return nfaStack[0]
}

/////////////////////////////MATCHING/////////////////////

func Pomatch(po string, s string) bool {
	//this function evaluates a string using a ReturnNFA
	ismatch := false

	ponfa := ReturnNFA(po)

	current := []*state{}
	next := []*state{}

	current = addState(current[:], ponfa.initial, ponfa.accept)

	for _, r := range s {
		for _, c := range current {
			if c.symbol == r {
				next = addState(next[:], c.edge1, ponfa.accept)
			}
		}

		current, next = next, []*state{}
	}

	for _, c := range current {
		if c == ponfa.accept {
			ismatch = true
			break
		}
	}
	return ismatch
}

func addState(l []*state, s *state, a *state) []*state {

	l = append(l, s)

	if s != a && s.symbol == 0 {
		l = addState(l, s.edge1, a)
		if s.edge2 != nil {
			l = addState(l, s.edge2, a)
		}
	}

	return l
}
