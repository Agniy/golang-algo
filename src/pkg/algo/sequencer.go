package algo
/**
* Class to provide the sequence of runes to test algorithms
*/

type sequencer struct{
}

/** get sequence of random runes*/
func (s *sequencer)getSequence() []rune{
	return []rune{2, 9, 4, 8, 6, 7, 1, 3, 0, 5};
}
/** get the worst case, i.e sequence is reversed odered*/
func (s *sequencer)getWorstCase() []rune{
	return []rune{9, 8, 7, 6, 5, 4, 3, 2, 1, 0};
}