Ticket numbers usually consist of an even number of digits. A ticket number is considered _lucky_ if the sum of the first half of the digits is equal to the sum of the second half.

Given a ticket number `n`, determine if it's _lucky_ or not.

__Example__

* For `n = 1230`, the output should be <br />`isLucky(n) = true;`
* For `n = 239017`, the output should be<br />`isLucky(n) = false`.

__Input/Output__

* __[time limit] 4 seconds (go)__
* __[input] integer n__ <br /><br />A ticket number represented as a positive integer with an even number of digits.<br /><br />_Guaranteed constraints:_<br />`10 ≤ n < 106.`<br />

* __[output] boolean__<br />`true` if `n` is a lucky ticket number, `false` otherwise.