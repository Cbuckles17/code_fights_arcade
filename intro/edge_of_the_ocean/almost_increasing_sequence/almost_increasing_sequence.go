func almostIncreasingSequence(sequence []int) bool {
    removed := false
    last := -1000001
    
    for i := range sequence[:len(sequence) - 1] {
        //not increasing
        if sequence[i + 1] <= sequence[i] {
            //if we already removed 1 then fail
            if removed { 
                return false 
            } else { 
                //remove 1
                removed = true 
                //remove the smaller value if the i-1->i+1 is not increasing
                if sequence[i + 1] <= last {
                    sequence[i + 1] = sequence[i]
                }
                
            }
        }
        
        last = sequence[i]
    }
    
    return true
}