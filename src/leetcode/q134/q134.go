package q134

func canCompleteCircuit(gas []int, cost []int) int {
	totalGas := 0
	totalCost := 0
	left := 0
	start := 0
	for i:=0;i < len(gas);i ++ {
		totalGas += gas[i]
		totalCost += cost[i]
		left += gas[i] - cost[i]
		if left < 0 {
			start = i + 1
			left = 0
		}
	}
	if totalGas >= totalCost {
		 return start
	}
	return -1
}
