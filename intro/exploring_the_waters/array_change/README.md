You are given an array of integers. On each move you are allowed to increase exactly one of its element by one. Find the minimal number of moves required to obtain a strictly increasing sequence from the input.

__Example__

For `inputArray = [1, 1, 1]`, the output should be<br>`arrayChange(inputArray) = 3`.

__Input/Output__

+ __[time limit] 4 seconds (go)__
+ __[input] array.integer inputArray__<br><br>_Guaranteed constraints:_<br>`3 ≤ inputArray.length ≤ 105`,<br>`-105 ≤ inputArray[i] ≤ 105`.

+ __[output] integer__<br><br>The minimal number of moves needed to obtain a strictly increasing sequence from `inputArray`.<br>It's guaranteed that for the given test cases the answer always fits signed `32`-bit integer type.