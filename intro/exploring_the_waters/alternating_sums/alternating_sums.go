func alternatingSums(a []int) []int {
    returnArray := make([]int, 2)
    
    for i, val := range a {
        returnArray[i%2] += val
    }
    
    return returnArray
}
