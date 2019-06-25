import "sort"

func sortByHeight(a []int) []int {
    peoplepos := 0
    people := make([]int, 0)
    
    //extract all the people
    for _, val := range a {
        if val != -1 {
            people = append(people, val)
        }  
    }
    
    //sort the people
    sort.Ints(people)
    
    //readd the people
    for i, val := range a {
        if val != -1 {
            a[i] = people[peoplepos]
            peoplepos += 1
        }  
    }
    
    return a
}