package radix_tree

import (
    "strings"
)

// 判断str2是否是str1的前缀
func isPrefix(str1, str2 string) bool {
    return strings.HasPrefix(str1, str2)
}

// 判断str2是否是str1的部分前辍
func isPartPrefix(str1, str2 string) bool {
    return str1[0] == str2[0]
}

// 求两个字符串的公共前缀
func calcPrefix(str1, str2 string) int {
    idx := 0
    str2length := len(str2) - 1
    for ;idx < len(str1); idx++ {
        if idx > str2length {
            break
        }
        if str1[idx] != str2[idx] {
            break
        }
    }
    return idx
}
