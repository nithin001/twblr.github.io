package data_structures
import "fmt"

func findPeopleWithCommonInterest(data map[string][]string, interest string) []string {
	var result []string;
	for k, v := range data {
        if contains(v, interest) {
			result = append(result,k)
		}
    }
    fmt.Printf("%v",result)
    return result
}

func contains(src []string, elem string) bool {
	for _,element := range src {
  		if element==elem{
  			return true
  		}
	}
	return false
}
