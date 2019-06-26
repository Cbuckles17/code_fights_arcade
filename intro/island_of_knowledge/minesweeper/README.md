In the popular __Minesweeper__ game you have a board with some mines and those cells that don't contain a mine have a number in it that indicates the total number of mines in the neighboring cells. Starting off with some arrangement of mines we want to create a __Minesweeper__ game setup.

__Example__

For
```
matrix = [[true, false, false],
          [false, true, false],
          [false, false, false]]
```
the output should be
```
minesweeper(matrix) = [[1, 2, 1],
                       [2, 1, 1],
                       [1, 1, 1]]
```
Check out the image below for better understanding:

![Example](example.png)

__Input/Output__

+ __[execution time limit] 4 seconds (go)__

+ __[input] array.array.boolean matrix__<br><br>A non-empty rectangular matrix consisting of boolean values - `true` if the corresponding cell contains a mine, `false` otherwise.<br><br>_Guaranteed constraints:_<br>`2 ≤ matrix.length ≤ 5`,<br>`2 ≤ matrix[0].length ≤ 5`.

+ __[output] array.array.integer__<br><br>Rectangular matrix of the same size as `matrix` each cell of which contains an integer equal to the number of mines in the neighboring cells. Two cells are called neighboring if they share at least one corner.