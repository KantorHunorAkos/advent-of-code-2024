package main

type OrderingRule struct {
	x int
	y int
}

type OrderingRules []OrderingRule

type PageData struct {
	page        int
	rulesBefore []int
	rulesAfter  []int
}

type Update []PageData

type Updates []Update
