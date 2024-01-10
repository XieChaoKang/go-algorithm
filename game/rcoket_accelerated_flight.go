package game

import "math"

// 火箭多段加速飞行计算 秒级 支持小数位 精准到毫秒

var accelerationSecondRadiusList = [][]float64{{0, 10}, {10, 20}, {20, -1}}
var accelerationList = []float64{0.025, 0.04, 0.06}

// 计算已经走过的时间范围内走过的距离
func HaveGoneMulti(secondRadiusIndex int, sec float64) float64 {
	currentAcceleration := accelerationList[0]
	initialVelocity := float64(0)
	calcTime := sec
	for i := 0; i < secondRadiusIndex; i++ {
		if i >= len(accelerationSecondRadiusList) {
			break
		}
		second := accelerationSecondRadiusList[i]
		initialVelocity += (second[1] - second[0]) * accelerationList[i]
		calcTime -= second[1]
		currentAcceleration = accelerationList[i+1]
	}
	return initialVelocity*calcTime + (currentAcceleration * calcTime * calcTime * 0.5)
}

func CalculateMulti(sec float64) float64 {
	// 初始速度
	initialVelocity := float64(0)
	// 当前传递进来的时间参数在对应的时间范围的 实际参与计算的具体时间 需要减去之前已经走过的时间范围的时间
	calcTime := sec
	currentAcceleration := accelerationList[0]
	multi := float64(0)
	// 记录已经走过的时间范围下标
	var haveGoneSecondRadiusIndex []int
	// 优先考虑计算初始速度
	for index, second := range accelerationSecondRadiusList {
		// 当前时间小于时间范围的最大值 不用计算初始速度 直接跳出 -1: 无穷大
		if sec < second[1] || second[1] == -1 {
			break
		}
		calcTime = calcTime - (second[1] - second[0])
		initialVelocity += (second[1] - second[0]) * accelerationList[index]
		currentAcceleration = accelerationList[index+1]
		haveGoneSecondRadiusIndex = append(haveGoneSecondRadiusIndex, index)
	}
	// 已经走过了的时间范围需要计算一下走过距离
	for _, index := range haveGoneSecondRadiusIndex {
		multi += HaveGoneMulti(index, accelerationSecondRadiusList[index][1])
	}
	// 计算当前时间范围内走的距离
	multi += initialVelocity*float64(calcTime) + (currentAcceleration * float64(calcTime) * float64(calcTime) * 0.5)
	return multi
}

// 记录倍数对应的时间范围最大值
var multiSecondRadiusList = []int{125, 575}

func CalculateTimeByMulti(multi int) int {
	// 根据距离与时间范围快速找到对应的时间范围
	allSecond := float64(0)
	secondRadiusIndex := -1
	goneMulti := 0
	for index, temp := range multiSecondRadiusList {
		if multi < temp {
			break
		}
		secondRadiusIndex = index
		goneMulti = temp
		second := accelerationSecondRadiusList[index]
		allSecond = allSecond + (second[1] - second[0])
	}
	calcMulti := multi - goneMulti
	initialVelocity := float64(0)
	currentAcceleration := accelerationList[0]
	if secondRadiusIndex > -1 {
		for i := 0; i < secondRadiusIndex+1; i++ {
			if i >= len(accelerationSecondRadiusList) {
				break
			}
			second := accelerationSecondRadiusList[i]
			initialVelocity += float64(second[1]-second[0]) * accelerationList[i]
			currentAcceleration = accelerationList[i+1]
		}
	}
	t := solveQuadraticEquation(initialVelocity, currentAcceleration, float64(calcMulti)/100)
	return t + int(allSecond)
}

// 二元一次求解
func solveQuadraticEquation(initialVelocity, currentAcceleration, multi float64) int {
	sqrt := math.Sqrt(initialVelocity*initialVelocity + 2*currentAcceleration*multi)
	t := (-initialVelocity + sqrt) / currentAcceleration
	return int(math.Ceil(t))
}
