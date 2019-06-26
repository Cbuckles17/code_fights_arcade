Call two arms _equally strong_ if the heaviest weights they each are able to lift are equal.

Call two people _equally strong_ if their strongest arms are equally strong (the strongest arm can be both the right and the left), and so are their weakest arms.

Given your and your friend's arms' lifting capabilities find out if you two are equally strong.

__Example__

+ For `yourLeft = 10`,` yourRight = 15`, `friendsLeft = 15` and `friendsRight = 10`, the output should be<br>`areEquallyStrong(yourLeft, yourRight, friendsLeft, friendsRight) = true;`
+ For `yourLeft = 15`, `yourRight = 10`, `friendsLeft = 15` and `friendsRight = 10`, the output should be<br>`areEquallyStrong(yourLeft, yourRight, friendsLeft, friendsRight) = true;`
+ For `yourLeft = 15`, `yourRight = 10`, `friendsLeft = 15` and `friendsRight = 9`, the output should be<br>`areEquallyStrong(yourLeft, yourRight, friendsLeft, friendsRight) = false.`

__Input/Output__

+ __[time limit] 4 seconds (go)__
+ __[input] integer yourLeft__<br><br>A non-negative integer representing the heaviest weight you can lift with your left arm.<br><br>_Guaranteed constraints:_<br>`0 ≤ yourLeft ≤ 20`.

+ __[input] integer yourRight__<br><br>A non-negative integer representing the heaviest weight you can lift with your right arm.<br><br>_Guaranteed constraints:_<br>`0 ≤ yourRight ≤ 20`.

+ __[input] integer friendsLeft__<br><br>A non-negative integer representing the heaviest weight your friend can lift with his or her left arm.<br><br>_Guaranteed constraints:_<br>`0 ≤ friendsLeft ≤ 20`.

+ __[input] integer friendsRight__<br><br>A non-negative integer representing the heaviest weight your friend can lift with his or her right arm.<br><br>_Guaranteed constraints:_<br>`0 ≤ friendsRight ≤ 20`.

+ __[output] boolean__<br><br>`true` if you and your friend are equally strong, `false` otherwise.