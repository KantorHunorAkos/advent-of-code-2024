package main

type OrderingRule struct {
	x int
	y int
}
type OrderingRules []OrderingRule

type Update []int
type Updates []Update

type Data struct {
	rules   OrderingRules
	updates Updates
}

type UpdateData struct {
	rules  OrderingRules
	update Update
}
