import "strings"

func addBorder(picture []string) []string {
    returnArray := []string{strings.Repeat("*", len(picture[0]) + 2)}
    
    for _, val := range picture {
        returnArray = append(returnArray, "*" + val + "*")
    }
    
    returnArray = append(returnArray, strings.Repeat("*", len(picture[0]) + 2))
    
    return returnArray
}