### 收获
##### Go
对于嵌套slice，内部的slice还需要再一次进行初始化!
```
dp := make([][]int, m+1)
for i := range dp {
    dp[i] = make([]int, n+1)
}
```

##### 算法
![command_pattern](/statics/dynamic_range__unique_paths.jpg)