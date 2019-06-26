func avoidObstacles(inputArray []int) int {
    obstaclemap := make(map[int]int)
    goodstep := false
    maxobstacle := 0
    
    //create a map of the obstacles
    for _, val := range inputArray {
        obstaclemap[val] = val
        //find the max distance obstacle
        if val > maxobstacle { maxobstacle = val }
    }
    
    //loop up to size max obstacle + 1 step
    for i := 1; i <= maxobstacle + 1; i++ {
        //loop through the map and see if there are any collisions with the step size
        for _, val := range obstaclemap {
            if val%i == 0 {
                goodstep = false
                break
            } else {
                goodstep = true
            }
        }
        
        if goodstep == true {
            //if we ended up with a goodstep == true that means we got past the max obstacle
            return i
        }
    }
    
    return -1
}