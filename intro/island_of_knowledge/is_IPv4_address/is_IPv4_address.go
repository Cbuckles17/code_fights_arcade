import "strings"

func isIPv4Address(inputString string) bool {
    sections := strings.Split(inputString, ".")
    
    if len(sections) != 4 {
        return false
    }
    
    for _, val := range sections {
        numbersection, atoierr := strconv.Atoi(val)
        if atoierr != nil || numbersection < 0 || numbersection > 255 {
            return false
        }
    }
    
    return true
}

func isIPv4AddressV1(inputString string) bool {
    start, sectioncount := 0, 0
    
    for i, val := range inputString {
        if sectioncount > 3 {
            return false
        } else if val == '.' {
            numbersection, atoierr := strconv.Atoi(string(inputString[start: i]))
            if atoierr != nil || numbersection < 0 || numbersection > 255 {
                return false
            } else {
                sectioncount += 1
                start = i + 1
            }
        }
    }
    
    numbersection, atoierr := strconv.Atoi(string(inputString[start: len(inputString)]))
    if atoierr != nil || numbersection < 0 || numbersection > 255 {
        return false
    } else if sectioncount != 3 {
        return false
    }
    
    return true
}
