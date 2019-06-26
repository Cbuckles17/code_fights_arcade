func sumSquare(matrix [][]bool, coli int, rowi int) int {
    sum := 0
    
    //row above
    if rowi-1 >= 0 {
        if coli-1 >= 0 && matrix[rowi-1][coli-1] == true {
            sum += 1
        }
        if matrix[rowi-1][coli] == true {
            sum += 1
        }
        if coli+1 < len(matrix[rowi]) && matrix[rowi-1][coli+1] == true {
            sum += 1
        }
    }

    //current row
    if coli-1 >= 0 && matrix[rowi][coli-1] == true {
        sum += 1
    }
    if coli+1 < len(matrix[rowi]) && matrix[rowi][coli+1] == true {
        sum += 1
    }

    //row below
    if rowi+1 < len(matrix) {
        if coli-1 >= 0 && matrix[rowi+1][coli-1] == true {
            sum += 1
        }
        if matrix[rowi+1][coli] == true {
            sum += 1
        }
        if coli+1 < len(matrix[rowi]) && matrix[rowi+1][coli+1] == true {
            sum += 1
        }
    }
    
    return sum
}

func minesweeper(matrix [][]bool) [][]int {
    returnMatrix := make([][]int, len(matrix))
                         
    for rowi := 0; rowi < len(matrix); rowi++ {  
        for coli := 0; coli < len(matrix[rowi]); coli++ {
            returnMatrix[rowi] = append(returnMatrix[rowi], sumSquare(matrix, coli, rowi))
        }
    }
    
    return returnMatrix
}
