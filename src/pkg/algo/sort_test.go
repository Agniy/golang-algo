package algo
import "testing"
/**
* Test sort algorithms 
*
*/

/** Test "insertion sort" algo*/
func TestInsertionSort(t *testing.T){
	/** get input sequence*/
	sqncer := new(sequencer);
	var sortedSqnce, initSqnce []rune;
	initSqnce = sqncer.getSequence();
	/** sort sequence*/ 
	sortedSqnce = InsertionSort(initSqnce);
	/** validate result*/
	validate(t, sortedSqnce);
}

/** Test "merge sort" algo*/
func TestMergeSort(t *testing.T){
	/** get input sequence*/
	sqncer := new(sequencer);
	var sortedSqnce, initSqnce []rune;
	initSqnce = sqncer.getSequence();
	/** sort sequence*/ 
	sortedSqnce = MergeSort(initSqnce);
	/** validate result*/
	validate(t, sortedSqnce);
}

/** Benchmark "insertion sort" algo*/
func BenchmarkInsertionSort(b *testing.B){
	for i:= 0; i < b.N; i++ {
		b.StopTimer();
		/** get input sequence*/
		sqncer := new(sequencer);
		var initSqnce []rune;
		initSqnce = sqncer.getSequence();
		b.StartTimer()
		/** sort sequence*/ 
		InsertionSort(initSqnce);
	}
}

/** validate result*/
func validate(t *testing.T, result []rune){
	/** check that a1 < a2 <...<aN */
	end := len(result) -1
	for i:=0; i < end ; i++ {
		if(result[i] > result[i+1]){
			t.Logf("Sequence is not sorted. See index {%d}; sequence {%v}", 
				i, result);
			t.FailNow();
		}
	}
}

