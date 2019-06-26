You are given an array of integers representing coordinates of obstacles situated on a straight line.

Assume that you are jumping from the point with coordinate 0 to the right. You are allowed only to make jumps of the same length represented by some integer.

Find the minimal length of the jump enough to avoid all the obstacles.

__Example__

For `inputArray = [5, 3, 6, 7, 9]`, the output should be<br>
`avoidObstacles(inputArray) = 4`.

Check out the image below for better understanding:

![Example](example.png)

__Input/Output__

+ __[time limit] 4 seconds (go)__

+ __[input] array.integer inputArray__<br><br>Non-empty array of positive integers.<br><br>_Guaranteed constraints:_<br>`2 ≤ inputArray.length ≤ 10`,<br>`1 ≤ inputArray[i] ≤ 40`.

+ __[output] integer__<br><br>The desired length.