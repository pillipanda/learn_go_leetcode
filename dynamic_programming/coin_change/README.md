### 收获
##### Go

##### 算法
1. 假设有i1,i2,i3, ..., in总共n种面额的硬币，目标是凑amount，令最少的硬币数是x，即 f(amount) = x
2. 那么与子问题的关系就是 f(amount) = min(f(amount - i1), f(amount - i2), f(amount - i3), ..., f(amount - in)) + 1
3. 故令从子问题开始记录的信息数组是dp，则dp[index]的值含义是当amount == index时所需的最少硬币数（1 <= index <= amount）