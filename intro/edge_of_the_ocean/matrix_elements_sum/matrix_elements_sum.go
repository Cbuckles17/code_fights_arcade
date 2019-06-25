func matrixElementsSum(matrix [][]int) int {
    sum := 0
    
    for rowi := range matrix {
        for coli, colv := range matrix[rowi] {
            if rowi == 0 {
                //if first row then add everything
                sum += colv
            } else if matrix[rowi - 1][coli] == 0 {
                //if cell above is 0 then set current cell to 0
                matrix[rowi][coli] = 0
            } else {
                //cell above is not 0 so add to sum
                sum += colv
            }
        }
    }
    
    return sum
}