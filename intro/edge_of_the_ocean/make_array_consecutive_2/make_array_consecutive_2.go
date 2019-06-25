func makeArrayConsecutive2(statues []int) int {
    maxheight := statues[0]
    minheight := statues[0]
    
    for _, val := range statues {
        if val > maxheight {
            maxheight = val
        }
        
        if val < minheight {
            minheight = val
        }
    }
    
    return (maxheight - minheight + 1) - len(statues)
}