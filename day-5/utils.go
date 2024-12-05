package main

func findApplingRules(rules OrderingRules, page int) OrderingRules {
	filteredRules := OrderingRules{}

	for _, rule := range rules {
		if page == rule.x {
			filteredRules = append(filteredRules, rule)
		}
	}

	return filteredRules
}


func (data UpdateData) Len() int { return len(data.update) }
func (data UpdateData) Swap(i, j int) {
	data.update[i], data.update[j] = data.update[j], data.update[i]
}
func (data UpdateData) Less(i, j int) bool {
	return needSwap(data.rules, data.update[i], data.update[j])
}

func needSwap(rules OrderingRules, p0 int, p1 int) bool {
	rulesApply := findApplingRules(rules, p0)
	for _, rule := range rulesApply {
		if rule.y == p1 {
			return true
		}
	}
	return false
}
