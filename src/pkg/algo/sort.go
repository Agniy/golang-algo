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
	
	split(0, len(input) - 1, input, eq);
	return input;
}

func split(start int, end int, input []interface{}, eq Equals){
	length := end -start +1
	merge(start, start + length/2, start + length, input, eq)
}

/** merge two sorted subsets, where
* whole unsorted array: A[start, end]
* sorted subset 1: A[start, media]
* sorted subset 2: A[media+1, end]
* and start <= media <= end
* @returns - merged array
*/
func merge(start int, media int, end int, input []interface{}, eq Equals){
	
	l := media -start + 1; //left hand length
	r := end - media; //right hand length
	
	if (l != 1) && (r != 1){
		split(start, media, input, eq);
		split(media +1, media + 1, input, eq);
	}	
	
	var left []interface{} = make([]interface{}, l); //left hand
	var right []interface{} = make([]interface{}, r); //right hand
	
	left = input[start:media]; //copy the first subset into the left hand
	right = input[media+1:end]; //copy teh second subset into the right hand
	
	/** start merging*/
	var i, j, k int = start, media, start 
	for ;(i < media) && (j < end); k++{
		if eq(left[i], right[j]) == -1 {
			//input[i] < input[j], we take input [i] 
			input[k] = left[i];
			i++;
		} else {
			input[k] = right[j];
			j++;
		}
	}
	
	if i == media {
		copy(input[k:],right[j:end]);
	}else{
		copy(input[k:],right[i:media]);
	}
}

/** swap elements in input*/
func swap(index1 int, index2 int,  input []interface{}){
	tmp := input[index1]
	input[index1] = input[index2]
	input[index2] = tmp
}
