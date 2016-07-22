package data_structures

type mapOperation func(int32) int32
type filterOperation func(int32) bool

func mapInts(op mapOperation, vals []int32) []int32 {
	var result []int32
	for _,elem := range vals{
		result = append(result,op(elem))
	}
	return result

}

func filterInts(op filterOperation, vals []int32) []int32 {
	var result []int32
	for _,elem := range vals{
		if op(elem){
			result = append(result,elem)
		}
	}
	return result
}

func concatenate(dest []string, newValues ...string) []string {
	for _,elem := range newValues{
		dest = append(dest,elem)
	}
	return dest
}

func equals(list1 []string, list2 []string) bool {
	if len(list1)!=len(list2){
		return false
	}else{
		for index,_:=range list1{
			if list1[index]!=list2[index]{
				return false
			}
		}
		return true
	}
}

func partialReverse(src []int, from, to int) []int {
	var result []int
	for i := to; i < from; i-- {
		result = append(result,src[i])
	}
	return result
}
