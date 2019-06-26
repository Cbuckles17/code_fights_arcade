Given a rectangular matrix of characters, add a border of asterisks(`*`) to it.

__Example__

For
```
picture = ["abc",
           "ded"]
```
the output should be

```
addBorder(picture) = ["*****",
                      "*abc*",
                      "*ded*",
                      "*****"]
```
                      
__Input/Output__

* __[time limit] 4 seconds (go)__
* __[input] array.string picture__  <br /><br />A non-empty array of non-empty equal-length strings.<br /><br />_Guaranteed constraints:_<br />`1 ≤ picture.length ≤ 5`,<br />`1 ≤ picture[i].length ≤ 5`.

* __[output] array.string__ <br /><br />The same matrix of characters, framed with a border of asterisks of width `1`.