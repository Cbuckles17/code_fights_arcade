func checkPalindrome(inputString string) bool {
    
    for left, right := 0, len(inputString) - 1; left < len(inputString)/2; left, right = left+1, right-1 {
        if inputString[left] != inputString[right] {
            return false
        }
    }
    
    return true
}
