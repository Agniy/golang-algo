package algo
import (
	"testing"
	"fmt"
)
/**
* Test sort algorithms 
*
*/


/** define compare func*/
var equals Equals = equalsRune; //checking rune type 

/** method to  generate sequence*/
type sequencer func () []interface{}
var seq sequencer = getRuneSequence; //generate sequence of rune 

/** method to generate expected result*/
type sortedSequancer func () []interface{}
var sorted sortedSequancer = getSortedRuneSequence;

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
	input := seq();
	/** sort sequence*/ 
	result := InsertionSort(input, equals);
	/** validate result*/
	validate(t, input, sorted(), result);
}
/** Tets merge*/
func TestMergeLeftEqualsRight(t *testing.T){
	var _seq sequencer = func ()[]interface{}{
		return []interface{}{Rune(1),Rune(3), Rune(5), Rune(2), Rune(4), Rune(6)};
	};
	sorted := []interface{}{Rune(1),Rune(2), Rune(3), Rune(4), Rune(5), Rune(6)};
	result := merge(0, 2, 5, _seq(), equals);
	validate(t, _seq(), sorted, result);
}
/** Tets merge*/
func TestMergeLeftOverRight(t *testing.T){
	var _seq sequencer = func ()[]interface{}{ 
		return []interface{}{Rune(1), Rune(2), Rune(3), Rune(5), Rune(4), Rune(6)};
	};
	sorted := []interface{}{Rune(1),Rune(2), Rune(3), Rune(4), Rune(5), Rune(6)};
	result := merge(0, 3, 5, _seq(), equals);
	validate(t, _seq(), sorted, result);
}

/** Tets merge*/
func TestMergeRightOverLeft(t *testing.T){
	var _seq sequencer = func ()[]interface{}{
		return []interface{}{Rune(1), Rune(3), Rune(2), Rune(4), Rune(5), Rune(6)};
	}
	sorted := []interface{}{Rune(1),Rune(2), Rune(3), Rune(4), Rune(5), Rune(6)};
	result := merge(0, 1, 5, _seq(), equals);
	validate(t, _seq(), sorted, result);
}

/** Test "merge sort" algo*/
func TestMergeSort(t *testing.T){
	input := seq();
	/** sort sequence*/ 
	result := MergeSort(seq(), equals);
	/** validate result*/
	validate(t, input, sorted(), result);
	
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

/** Benchmark "merge sort" algo*/
func BenchmarkMergeSort(b *testing.B){
	for i:= 0; i < b.N; i++ {
		/** get input sequence*/
		initSqnce := seq();
		/** sort sequence*/ 
		b.StartTimer()
		MergeSort(initSqnce, equals) 
		b.StopTimer();
	}
}

/** validate result*/
func validate(t *testing.T, input []interface{}, expected[]interface{}, result []interface{}){
	
	/**print input*/
	fmt.Println("input:\t\t", input);
	/**print expected*/
	fmt.Println("expected:\t", expected);
	/**priint result*/
	fmt.Println("result:\t\t", result);
	if len(input) != len(result){
		t.Logf("Sequence has invalid size = {%v}, expected size = {%v}", 
				len(result), len(input));
		t.FailNow();
	}
	/** check that expected result matches the actual result */
	for i:=0; i < len(result) ; i++ {
		if equals(expected[i], result[i]) != 0 {
			t.Logf("The expected result doesn't match the actual result. See index {%d}", i);
			t.FailNow();
		}
	}
}

