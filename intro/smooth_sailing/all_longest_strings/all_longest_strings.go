func allLongestStrings(inputArray []string) []string {
    retArray := make([]string,0)
    maxlen := len(inputArray[0])
    
    for i, v := range inputArray {
        if len(inputArray[i]) > maxlen {
            maxlen = len(inputArray[i])
            retArray = []string{v}
        } else if len(inputArray[i]) == maxlen {
            retArray = append(retArray, v)
        }
    }
    
    return retArray
}