### 收获
##### Go
1. 使用math库获取类似inf的最大最小值

##### 算法
用简单思维最好，核心就是找到后面的最大的来减去前面的最小的，
这种情况下同理找后面最大的等价于找max(后面 - 前面最小)
所以不断在循环中保持阶段性的 前面最小 与 max(后面 - 前面最小) 即可
