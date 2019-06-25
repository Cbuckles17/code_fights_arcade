func shapeArea(n int) int {
    area := 1
    
    for i := 0; i < n; i++ {
        area += i * 4
    }
    
    return area
}