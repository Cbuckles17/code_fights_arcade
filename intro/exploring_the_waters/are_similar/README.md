Two arrays are called _similar_ if one can be obtained from another by swapping at most one pair of elements in one of the arrays.

Given two arrays `a` and `b`, check whether they are similar.

__Example__

+ For `a = [1, 2, 3]` and `b = [1, 2, 3]`, the output should be <br>`areSimilar(a, b) = true`.<br><br>The arrays are equal, no need to swap any elements.

+ For `a = [1, 2, 3]` and `b = [2, 1, 3]`, the output should be<br>`areSimilar(a, b) = true`.<br><br>We can obtain b from a by swapping 2 and 1 in b.

+ For `a = [1, 2, 2]` and `b = [2, 1, 1]`, the output should be<br>`areSimilar(a, b) = false`.<br><br>Any swap of any two elements either in a or in b won't make a and b equal.

__Input/Output__

+ __[time limit] 4 seconds (go)__
+ __[input] array.integer a__<br><br>Array of integers.<br><br>_Guaranteed constraints:_<br>`3 ≤ a.length ≤ 105`,<br>`1 ≤ a[i] ≤ 1000`.

+ __[input] array.integer b__<br><br>Array of integers of the same length as `a`.<br><br>_Guaranteed constraints:_<br>`b.length = a.length`,<br>`1 ≤ a[i] ≤ 1000`.

+ __[output] boolean__<br><br>`true` if `a` and `b` are similar, `false` otherwise.