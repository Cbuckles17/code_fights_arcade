func areSimilar(a []int, b []int) bool {
    mismatchc := 0
    mismatches := make(map[int]int)
    
    for i, val := range a {
        if val != b[i] {
            
            mismatchc += 1
            mismatches[val] += 1
            mismatches[b[i]] += 1
        }       
    }
    
    if mismatchc == 0 {
        //if there were zero mismatches
        return true
    } else if mismatchc == 2 && len(mismatches) == 2 {
        //if there were 2 mismatches 
        //and 2 swappable chars aka there can be 2 mismatches but
        //a b and b c would not be swappable as they lead to a mismatch
        return true
    } else {
        //anythign else means failure
        return false
    }
}

func areSimilarV1(a []int, b []int) bool {
    swapped := false
    
    //loop through one of the arrays, doesn't matter which
    for i, val := range a {
        if val != b[i] {
            //we have a mismatch
            if swapped == true { 
                //if we have already swapped then fail
                return false 
            } else {
                //lets swap
                swapped = true
                
                //loop through the array you're doing the swapping in
                for j, jval := range a {
                    if jval == b[i] && jval != b[j] {
                        //if the value is what you need to swap 
                        //and the value at that index in both arrays doesn't already match
                        a[i], a[j] = a[j], a[i]
                    }
                }
            }
        }
    }
    
    return true
}
