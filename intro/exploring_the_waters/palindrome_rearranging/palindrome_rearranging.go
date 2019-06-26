func palindromeRearranging(inputString string) bool {
    inputMap := make(map[rune]int)
    oddfound := false
    
    //map the chars and their counts
    for _, val := range inputString {
        inputMap[val] += 1
    }
    
    //loop through the map
    for _, val := range inputMap {
        if val%2 == 1 {
            //if the char count is odd
            if oddfound == true || len(inputString)%2 == 0{
                //if we already had a odd
                //or the string is of even length aka there cannot be a middle char
                return false
            } else {
                oddfound = true
            } 
        }
    }
    
    return true
}