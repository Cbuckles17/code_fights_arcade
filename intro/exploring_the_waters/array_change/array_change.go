func arrayChange(inputArray []int) int {
    increased := 0
    
    //slice the loop 1 less so we don't go out of bounds on the i+1 calls
    for i, val := range inputArray[:len(inputArray) - 1] {
        if val >= inputArray[i+1] {
            //if it is not ascending
            increased += val - inputArray[i+1] + 1
            //subtract the two values and add 1 to find the lowest value to make it ascending
            inputArray[i+1] += val - inputArray[i+1] + 1
        }
    }
    
    return increased
}
