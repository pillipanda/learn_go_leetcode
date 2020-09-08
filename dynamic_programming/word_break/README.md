### 收获
##### Go

##### 算法
1. 令输入为s，定义dp[index]表示s[:index]是否是能够凑出来的子串、若是则值为1、否则为0
2. 同时令输入的word字典为word_dict
3. 故解决index为x时候的思路就是遍历当i = 1; i < index; i++;时，若dp[i]值为1，则检验s[index+1:x]是否在word_dict中，若是，设置dp[x]为1；