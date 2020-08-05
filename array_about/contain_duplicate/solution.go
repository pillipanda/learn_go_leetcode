package contain_duplicate

func ContainDuplicate(iptSlice []int)  bool{
	iptMapExist := map[int]bool{}

	for _, ipt := range iptSlice {
		if iptMapExist[ipt] {
			return true
		}
		iptMapExist[ipt] = true
	}
	return false
}
