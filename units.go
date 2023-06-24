package main

var zombieKind = unitKind{
	name:           "Zombie",
	hitPoints:      100,
	perception:     10000,
	attackRange:    1,
	attackAccuracy: 99,
	attackRate:     0,
	attackDamage:   10,
	attackArea:     0,
	moveSpeed:      1,
	slotCount:      1,
}

var infantryKind = unitKind{
	name:           "Infantry",
	hitPoints:      150,
	perception:     30,
	attackRange:    20,
	attackAccuracy: 50,
	attackRate:     1,
	attackDamage:   100,
	attackArea:     0,
	moveSpeed:      2,
	slotCount:      1,
}

var operatorKind = unitKind{
	name:           "Operator",
	hitPoints:      200,
	perception:     30,
	attackRange:    30,
	attackAccuracy: 90,
	attackRate:     0,
	attackDamage:   10,
	attackArea:     0,
	moveSpeed:      3,
	slotCount:      1,
}

var tankKind = unitKind{
	name:           "Tank",
	hitPoints:      5000,
	perception:     600,
	attackRange:    500,
	attackAccuracy: 90,
	attackRate:     3,
	attackDamage:   2500,
	attackArea:     2,
	moveSpeed:      6,
	slotCount:      40,
}

var howitzerKind = unitKind{
	name:           "Howitzer",
	hitPoints:      1000,
	perception:     1000,
	attackRange:    1000,
	attackAccuracy: 99,
	attackRate:     3,
	attackDamage:   3000,
	attackArea:     3,
	moveSpeed:      0,
	slotCount:      10,
}
