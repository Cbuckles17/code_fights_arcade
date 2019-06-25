func commonCharacterCount(s1 string, s2 string) int {
    m1 := make(map[string]int)
    m2 := make(map[string]int)
    count := 0
    
    //make a char map of s1
    for _, c := range s1 {
        m1[string(c)] += 1
    }
    
    //make a char map of s2
    for _, c := range s2 {
        m2[string(c)] += 1
    }
    
    //loop through either of the maps, doing m1 in this case
    for key, val := range m1 {
        if val <= m2[key] {
            //add the char1 count as that is the lowest
            count += val
        } else if val > m2[key] && m2[key] != 0 {
            //add the char2 count as that is the lowest and not 0
            count += m2[key]
        }
    }
    
    return count
}