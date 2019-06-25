import "strings"

func reverseInParentheses(inputString string) string {
    leftp, rightp := 0, 0
    
    //loop until there are no more (
    for strings.Index(inputString, "(") != -1 {
        //loop through the string
        for j := range inputString {
            if inputString[j] == '(' {
                //set the last ( position
                leftp = j
            } else if inputString[j] == ')' {
                //set the first ) position
                rightp = j
                break
            }
        }
        
        //convert to rune array so we can mutate it
        runes := []rune(inputString)
        
        //reverse the runes
        for left, right := leftp + 1, rightp - 1; left < right; left, right = left+1, right-1 {
            runes[left], runes[right] = runes[right], runes[left]
        }
        
        //slice the rune array into 3 parts
        firstSection := runes[0:leftp]
        reverseSection := runes[leftp+1 : rightp]
        endSection := runes[rightp+1 :]
        
        //overwrite the orginal string so our for loops will update accordingly
        inputString = string(firstSection) + string(reverseSection) + string(endSection)
        //not necessary but cannot hurt to reset these
        leftp, rightp = 0, 0
    }
    
    return inputString
}