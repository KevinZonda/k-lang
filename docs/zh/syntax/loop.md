# 循环

K 语言支持多种风格的循环。

## 循环整个数组

对于一个数组，我们可以使用 `for (element : list)` 语法来遍历整个数组。

```k
a = [1, 2, 3, 4, 5]
for (i : a) {
    println(i)
}
# Output:
# 1
# 2
# 3
# 4
# 5
```


## C语言风格的循环

```k
for (i := 0; i < 5; i = i + 1) {
    println(i)
}
# Output:
# 0
# 1
# 2
# 3
# 4
```