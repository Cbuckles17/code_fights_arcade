func arrayMaximalAdjacentDifference(inputArray []int) int {
    min := -30
    
    for i, val := range inputArray[:len(inputArray) - 1] {
        if (val - inputArray[i + 1]) > min {
            min = val - inputArray[i + 1]
        }
        
        if (inputArray[i + 1] - val) > min {
            min = inputArray[i + 1] - val
        }
    }
    
    return min
}