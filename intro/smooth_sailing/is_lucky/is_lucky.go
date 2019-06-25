func isLucky(n int) bool {
    sum := 0
    ogn := n
    digitcount := 0
    counter := 0
    
    //count the number of digits
    for n > 0 {
        digitcount += 1
        n = n/10
    }
    
    //loop through the original n
    for ogn > 0 {
        if counter += 1; counter <= digitcount/2 {
            //if we are before the halfway then add
            sum += ogn%10
        } else {
            //if we are past the halfway then subtract
            sum -= ogn%10
        }
        
        ogn = ogn/10
    }
    
    if sum != 0 { 
        return false
    } else { 
        return true 
    }
}