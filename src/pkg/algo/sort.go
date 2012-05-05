package algo


/**
* File to contain sorting algo implementation
* Attention: we modify the input slice
**/

func InsertionSort(input []interface{}, eq Equals) []interface{}{
	
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
			if eq(input[j], input[j-1]) != -1 {
				//input[j] >= input[j-1]
				break; //we found element < [j]
			}
			swap(j, j-1, input); //move element[j] left
		}
	}
	
	return input;
}

func MergeSort(input []interface{}, eq Equals) []interface{}{
	mergeSort(0, len(input) - 1, input, eq);
	return input;
}

func mergeSort(start int, end int, input []interface{}, eq Equals) (int, int){

	//array length
	length := end - start +1;
	if length == 1{
		// here we finish splitting
		return start, end;
	}
	media := start + length/2
	
	lstart, lend := mergeSort(start, media, input, eq)
	rstart, rend := mergeSort(media + 1, end, input, eq)
	
	merge(lstart, lend, rstart, rend, input, eq);
	return lstart, rend;
}


/** merge two sorted subsets, where
* whole unsorted array: A[start, end]
* sorted subset 1: A[start, media]
* sorted subset 2: A[media+1, end]
* and start <= media <= end
* @returns - merged array
*/
func merge(lstart int, lend int, rstart int, rend int, input []interface{}, eq Equals) []interface{}{
	
	l := lend - lstart + 1; //left hand length
	r := rend - rstart +1; //right hand length
	
	var left []interface{} = make([]interface{}, l); //left hand
	var right []interface{} = make([]interface{}, r); //right hand
	
	//copy sorted subsets into left and right hand respectively
	copy(left, input[lstart:lend + 1]);
	copy(right, input[rstart:rend + 1]);
	
	/** start merging*/
	var i, j, k int = 0, 0, lstart 
	for ;(i < l) && (j < r); k++{
		if eq(left[i], right[j]) == -1 {
			//input[i] < input[j], we take input [i] 
			input[k] = left[i];
			i++;
		} else {
			input[k] = right[j];
			j++;
		}
	}
	
	if k < rend {
		//some data remaind in the left or right hand
		//case: left and right hands have  different size
		if i == l {
			//data remained in the right hand
			copy(input[k:],right[j:]);
		}else{
			//data remained in the left hand hand
			copy(input[k:],left[i:]);
		}
	}
	return input;
}

/** swap elements in input*/
func swap(index1 int, index2 int,  input []interface{}){
	tmp := input[index1]
	input[index1] = input[index2]
	input[index2] = tmp
}
