package algo

/** synonym for type rune*/
type Rune rune

/** compare func for the type rune*/
func equalsRune(arg0 interface{}, arg1 interface{}) int{
	runeArg0, runeArg1 := arg0.(Rune), arg1.(Rune);
	if runeArg0 < runeArg1 {
		return -1;
	} else if runeArg0 == runeArg1 {
		return 0;
	}
	return 1;
}

/** generate random Rune sequence*/
func getRuneSequence() []interface{} {
	return []interface{}{Rune(2), Rune(9), Rune(4), Rune(8), Rune(6), Rune(7), Rune(1), Rune(3), 
		Rune(0), Rune(5), Rune(8)}
}

/** generate a pair of Rune elemets, where more > less */
func getRuneMoreAndLess() (more interface{}, less interface{}){
	return Rune(3), Rune(2)
}


