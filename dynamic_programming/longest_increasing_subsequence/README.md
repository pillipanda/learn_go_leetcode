### 收获
##### Go

##### 算法
这道题还挺简单的，相比凑硬币
1. 假设数组的构成是x1,x2,x3, ..., xn，令index为i的数字的最长上升子串长度为x，即f(i) = x 
2. 那么与子问题的关系就是 f(i) = max([f(index) for index in range(0, i) if xindex < xi ]) + 1
3. 故令从子问题开始记录的信息数组是dp，则dp[index]的值含义是此数组index下的值的最长子串长度（即f(index)）