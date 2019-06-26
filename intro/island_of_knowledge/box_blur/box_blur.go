func boxBlur(image [][]int) [][]int {
    returnImage := make([][]int, len(image) - 2)
    
    for rowi := 0; rowi <= len(image) - 3; rowi++ {  
        for coli := 0; coli <= len(image[rowi]) - 3; coli++ {
            blocksum := image[rowi][coli] + image[rowi][coli + 1] + image[rowi][coli + 2] +
            image[rowi + 1][coli] + image[rowi + 1][coli + 1] + image[rowi + 1][coli + 2] +
            image[rowi + 2][coli] + image[rowi + 2][coli + 1] + image[rowi + 2][coli + 2]
            
            returnImage[rowi] = append(returnImage[rowi], blocksum/9)
        }
        
    }
    
    return returnImage
}