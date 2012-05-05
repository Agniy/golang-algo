package algo
import "testing"
/**
* Test sort algorithms 
*
*/


/** define compare func*/
var equals Equals = equalsRune; //checking rune type 

/** method to  generate sequence*/
type sequencer func () []interface{}
var seq sequencer = getRuneSequence; //generate sequence of rune 

/** method to generate more & less values*/
type moreless func()(interface{}, interface{}) 
var mrls moreless = getRuneMoreAndLess; //generate a pair of [more,less] rune values 


/** Test equals method*/
func TestEquals(t *testing.T){
	arg0, arg1 := mrls(); //get pair of [more,less] values, where arg0 > arg 1
	if equals(arg0, arg1) != 1 {
		t.Fatal();
	}
	if equals(arg1, arg0) != -1 {
		t.Fatal();
	}
	if equals(arg0, arg0) != 0 {
		t.Fatal();
	}
}

/** Test "insertion sort" algo*/
func TestInsertionSort(t *testing.T){
	/** sort sequence*/ 
	sortedSqnce := InsertionSort(seq(), equals);
	/** validate result*/
	validate(t, sortedSqnce);
}
/** Tets merge*/
func TestMergeLeftEqualsRight(t *testing.T){
	input := []interface{}{Rune(1),Rune(3), Rune(5), Rune(2), Rune(4), Rune(6)};
	validate(t, merge(0, 2, 3, 5, input, equals));
}
/** Tets merge*/
func TestMergeLeftOverRight(t *testing.T){
	input := []interface{}{Rune(1), Rune(2), Rune(3), Rune(5), Rune(4), Rune(6)};
	validate(t, merge(0, 3, 4, 5, input, equals));
}

/** Tets merge*/
func TestMergeRightOverLeft(t *testing.T){
	input := []interface{}{Rune(1), Rune(3), Rune(2), Rune(4), Rune(5), Rune(6)};
	validate(t, merge(0, 1, 2, 5, input, equals));
}

/** Test "merge sort" algo*/
func TestMergeSort(t *testing.T){
	/** sort sequence*/ 
	sortedSqnce := MergeSort(seq(), equals);
	/** validate result*/
	validate(t, sortedSqnce);
}

/** Benchmark "insertion sort" algo*/
func BenchmarkInsertionSort(b *testing.B){
	for i:= 0; i < b.N; i++ {
		/** get input sequence*/
		initSqnce := seq();
		/** sort sequence*/ 
		b.StartTimer()
		InsertionSort(initSqnce, equals) 
		b.StopTimer();
	}
}

/** validate result*/
func validate(t *testing.T, result []interface{}){
	/** check that a1 < a2 <...<aN */
	end := len(result) -1
	for i:=0; i < end ; i++ {
		// check if result[i] <= result[i+1]
		if equals(result[i], result[i+1]) == 1 {
			// result[i] > result[i+1], this is the issue !
			t.Logf("Sequence is not sorted. See index {%d}; sequence {%v}", 
				i, result);
			t.FailNow();
		}
	}
}

