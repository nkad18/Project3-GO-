# Project3-GO-
A nondeterministic finite automaton (NFA) is defined by a set of states, symbols in an
alphabet, and a transition function. A state is represented by an integer. A symbol is
represented by a rune, i.e., a character. Given a state and a symbol, a transition function
returns the set of states that the NFA can transition to after reading the given symbol. This
set of next states could be empty.


A given final state is said to be reachable from a given start state via a given input
sequence of symbols if there exists a sequence of transitions such that if the NFA starts
at the start state it would reach the final state after reading the entire sequence of input
symbols.
