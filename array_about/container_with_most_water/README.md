### 收获

##### Go

##### 算法
- 算法解释:
    - water = min(leftVal, rightVal) * (rightIndex - leftIndex)
    - 首先无论谁移动、距离一定是变小的，故而想要变大就要提高min(leftVal, rightVal)，从而需要将更小的一边进行移动获取可能更大的此部分值
    - 由于移动小的部分是为了获得更高的高度、所以如果移动后的高度值比之前的还低，直接继续往下移动即可
