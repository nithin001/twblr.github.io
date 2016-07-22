package control_structures

func collatzChainLength(n int) int {
	var count int=0;
	for{
		if(n==1){
			break
		} 
		if(n%2==0){ 
			n/=2
			}else{
				n=(n*3)+1
			}
		count++
	}
	return count
}
