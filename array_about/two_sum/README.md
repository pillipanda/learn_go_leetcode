### 收获

##### Go
1. 从之前还在使用基于$GOPATH的方式转换到基于go mod处理依赖
2. 学会写基本的测试
3. 有犯错在map的赋值上：
```
错误写法例子：
    var numMapIndex map[int]int
    numMapIndex[1] = 2
会报出错误： panic: assignment to entry in nil map 

错误说明：
创建返回的是引用、而map引用类型的零值是nil，即没有引用任何哈希表

正确写法（初始化引用到一个空map对象）：
    numMapIndex := map[int]int{}
    或：
    numMapIndex := make(map[int]int)
```

