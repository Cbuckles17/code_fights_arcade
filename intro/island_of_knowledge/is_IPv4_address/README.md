An IP address is a numerical label assigned to each device (e.g., computer, printer) participating in a computer network that uses the Internet Protocol for communication. There are two versions of the Internet protocol, and thus two versions of addresses. One of them is the _IPv4 address_.

_IPv4 addresses_ are represented in dot-decimal notation, which consists of four decimal numbers, each ranging from `0` to `255`, separated by dots, e.g., `172.16.254.1`.

Given a string, find out if it satisfies the _IPv4 address_ naming rules.

__Example__

+ For `inputString = "172.16.254.1"`, the output should be<br>`isIPv4Address(inputString) = true`;

+ For `inputString = "172.316.254.1"`, the output should be<br>`isIPv4Address(inputString) = false`.<br><br>`316` is not in `range [0, 255]`.

+ For `inputString = ".254.255.0"`, the output should be<br>`isIPv4Address(inputString) = false`.<br><br>There is no first number.

__Input/Output__

+ __[time limit] 4 seconds (go)__
+ __[input] string inputString__<br><br>A string consisting of digits, full stops and lowercase Latin letters.<br><br>_Guaranteed constraints:_<br>`5 ≤ inputString.length ≤ 15`.

+ __[output] boolean__<br><br>`true` if `inputString` satisfies the `IPv4 address` naming rules, `false` otherwise.