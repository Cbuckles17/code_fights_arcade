Some people are standing in a row in a park. There are trees between them which cannot be moved. Your task is to rearrange the people by their heights in a non-descending order without moving the trees.

__Example__

For `a = [-1, 150, 190, 170, -1, -1, 160, 180]`, the output should be <br />
`sortByHeight(a) = [-1, 150, 160, 170, -1, -1, 180, 190]`.

__Input/Output__

* __[time limit] 4 seconds (go)__
* __[input] array.integer a__<br /><br />If `a[i] = -1`, then the `ith` position is occupied by a tree. Otherwise `a[i]` is the height of a person standing in the `ith position`.<br /><br />_Guaranteed constraints:_<br />`5 ≤ a.length ≤ 15`,<br />`-1 ≤ a[i] ≤ 200`.<br />

* __[output] array.integer__<br />Sorted array `a` with all the trees untouched.