package algo
/**
* File to contain sorting algo implementation
* Attention: we modify the input slice
**/

func InsertionSort(input []rune) []rune{
	
	length := len(input);
	/** input is already sorted*/
	if length < 2{
		return input;
	}
		
	/** let's sort */
	for i :=1; i < length; i++ {
		for j:= i; j > 0; j--{
			/** here we consider all elements [0, j -1] to be sorted out
			* We should correctly (i.e. sorted)place [j] into [0, j-1] set.
			* We shift all elemsnt in [0, j-1] to the right, until we find the first 
			* elemen <= [j]
			*/
			if input[j] >= input[j -1]{
				break; //we found element < [j]
			}
			swap(j, j-1, input); //move element[j] left
		}
	}
	
	return input;
}

func MergeSort(input []rune) []rune{
	
	/** split input first*/
	sp1, sp2 := split(input)
	
	return []rune{1,2,3};
}

/** split slice into two slices*/
func split(input []rune) []rune, rune[]{
	/** split point*/
	sp := len(input)/2
	return input[:sp], input[sp + 1:]
}


/** swap elements in input*/
func swap(index1 int, index2 int,  input []rune){
	tmp := input[index1]
	input[index1] = input[index2]
	input[index2] = tmp
}
