func areEquallyStrong(yourLeft int, yourRight int, friendsLeft int, friendsRight int) bool {
    //just check the two possible cases for true and if not it's false
    if yourLeft == friendsLeft && yourRight == friendsRight {
        return true
    } else if yourLeft == friendsRight && yourRight == friendsLeft {
        return true 
    } else {
        return false
    }
}