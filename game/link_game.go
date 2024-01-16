package game

import (
	"math/rand"
	"time"
)

var nullPosition = 0

// 连连看游戏
type LinkGame struct {
	LinkMap [][]int // 0 代表没有东西 可以通行； 其他内容都代表有东西 无法通行
}

func (l *LinkGame) SetPositionNull(reel, line int) bool {
	if !l.IsValidPos(reel, line) {
		return false
	}
	l.LinkMap[reel][line] = nullPosition
	return true
}

func (l *LinkGame) FindPassablePath(startReel, startLine, endReel, endLine int) (bool, [][]int) {
	if startReel == endReel && startLine == endLine {
		return false, nil
	}
	if !l.IsValidPos(startReel, startLine) || !l.IsValidPos(endReel, endLine) || l.LinkMap[startReel][startLine] != l.LinkMap[endReel][endLine] {
		return false, nil
	}
	var suc bool
	var path [][]int
	if suc, path = l.FindTransverseConnectPath(startReel, startLine, endReel, endLine, false); suc {
		return suc, path
	}
	if suc, path = l.FindVerticalConnectPath(startReel, startLine, endReel, endLine, false); suc {
		return suc, path
	}
	if suc, path = l.FindLConnectPath(startReel, startLine, endReel, endLine); suc {
		return suc, path
	}
	if suc, path = l.FindZConnectPath(startReel, startLine, endReel, endLine); suc {
		return suc, path
	}
	if suc, path = l.FindHalfBorderPath(startReel, startLine, endReel, endLine); suc {
		return suc, path
	}
	return suc, path
}

// 边界判断
func (l *LinkGame) IsValidPos(reel, line int) bool {
	if len(l.LinkMap) == 0 {
		return false
	}
	return reel >= 0 && reel < len(l.LinkMap) && line >= 0 && line < len(l.LinkMap[0])
}

// 判断路标集合是否可以通行 二维固定长度为2 0：列 1：行
func (l *LinkGame) IsPassablePath(paths [][]int, checkTurningLine bool) bool {
	if len(paths) == 1 {
		return true
	}
	start := paths[0]
	end := paths[len(paths)-1]
	// 判断开头和结尾不一致的情况下 必须得有一个节点是空节点才代表可达
	if l.LinkMap[start[0]][start[1]] != l.LinkMap[end[0]][end[1]] && l.LinkMap[start[0]][start[1]] != nullPosition && l.LinkMap[end[0]][end[1]] != nullPosition {
		return false
	}
	// 如果是转折线 那开头和结尾一定要有一个是空节点
	if checkTurningLine && l.LinkMap[start[0]][start[1]] != nullPosition && l.LinkMap[end[0]][end[1]] != nullPosition {
		return false
	}
	// 去掉开头和结尾 判断路径中间是否已经被站位
	temp := paths[1 : len(paths)-1]
	for _, path := range temp {
		if l.LinkMap[path[0]][path[1]] != nullPosition {
			return false
		}
	}
	return true
}

// 横向直连 列号偏移 行号不变
// 是否可以通行 和 路径坐标数组（二维固定长度为2 0：列 1：行）
func (l *LinkGame) FindTransverseConnectPath(startReel, startLine, endReel, endLine int, checkTurningLine bool) (bool, [][]int) {
	if startLine != endLine {
		return false, nil
	}
	// 强行让开始节点一定在结束节点的左边 方便运算
	var tempStartReel, tempEndReel int
	if startReel <= endReel {
		tempStartReel = startReel
		tempEndReel = endReel
	} else {
		tempStartReel = endReel
		tempEndReel = startReel
	}
	var paths = [][]int{{tempStartReel, startLine}}
	for i := tempStartReel + 1; i <= tempEndReel; i++ {
		paths = append(paths, []int{i, endLine})
	}
	if l.IsPassablePath(paths, checkTurningLine) {
		return true, paths
	}
	return false, nil
}

// 竖向连接 行号偏移 列号不变
// 是否可以通行 和 路径坐标数组（二维固定长度为2 0：列 1：行）
func (l *LinkGame) FindVerticalConnectPath(startReel, startLine, endReel, endLine int, checkTurningLine bool) (bool, [][]int) {
	if startReel != endReel {
		return false, nil
	}
	// 强行让开始节点一定在结束节点的左边 方便运算
	var tempStartLine, tempEndLine int
	if startLine <= endLine {
		tempStartLine = startLine
		tempEndLine = endLine
	} else {
		tempStartLine = endLine
		tempEndLine = startLine
	}
	var paths = [][]int{{startReel, tempStartLine}}
	for i := tempStartLine + 1; i <= tempEndLine; i++ {
		paths = append(paths, []int{startReel, i})
	}
	if l.IsPassablePath(paths, checkTurningLine) {
		return true, paths
	}
	return false, nil
}

// L向连接
// 是否可以通行 和 路径坐标数组（二维固定长度为2 0：列 1：行）
func (l *LinkGame) FindLConnectPath(startReel, startLine, endReel, endLine int) (bool, [][]int) {
	// 强行让开始节点一定在结束节点的左边 方便运算
	var tempStartReel, tempStartLine, tempEndReel, tempEndLine int
	if startReel <= endReel {
		tempStartReel = startReel
		tempStartLine = startLine
		tempEndReel = endReel
		tempEndLine = endLine
	} else {
		tempStartReel = endReel
		tempStartLine = endLine
		tempEndReel = startReel
		tempEndLine = startLine
	}
	// L 有两种 分别查找
	var turningPointReel = tempStartReel
	var turningPointLine = tempEndLine
	has, transverseConnectPath := l.FindTransverseConnectPath(turningPointReel, turningPointLine, tempEndReel, tempEndLine, true)
	find, verticalConnectPath := l.FindVerticalConnectPath(tempStartReel, tempStartLine, turningPointReel, turningPointLine, true)
	if has && find {
		return true, append(verticalConnectPath[:len(verticalConnectPath)-1], transverseConnectPath...)
	}
	// 第二种
	turningPointReel = tempEndReel
	turningPointLine = tempStartLine
	has, transverseConnectPath = l.FindTransverseConnectPath(tempStartReel, tempStartLine, turningPointReel, turningPointLine, true)
	find, verticalConnectPath = l.FindVerticalConnectPath(turningPointReel, turningPointLine, tempEndReel, tempEndLine, true)
	if has && find {
		return true, append(transverseConnectPath[:len(transverseConnectPath)-1], verticalConnectPath...)
	}
	return false, nil
}

// Z 向连接
// 是否可以通行 和 路径坐标数组（二维固定长度为2 0：列 1：行）
func (l *LinkGame) FindZConnectPath(startReel, startLine, endReel, endLine int) (bool, [][]int) {
	// 强行让开始节点一定在结束节点的左边 方便运算
	var tempStartReel, tempStartLine, tempEndReel, tempEndLine int
	if startReel <= endReel {
		tempStartReel = startReel
		tempStartLine = startLine
		tempEndReel = endReel
		tempEndLine = endLine
	} else {
		tempStartReel = endReel
		tempStartLine = endLine
		tempEndReel = startReel
		tempEndLine = startLine
	}
	// 横向查找
	for i := tempStartReel + 1; i <= tempEndReel; i++ {
		// 横向移动过不去 直接结束
		if l.LinkMap[i][tempStartLine] != nullPosition {
			break
		}
		// 判断当前列是否可以通行到结束节点的那一行
		suc, path := l.FindVerticalConnectPath(i, tempStartLine, i, tempEndLine, true)
		if !suc {
			continue
		}
		// 判断和结束节点是同一个节点 是则代表已经找到了可行通路 直接返回
		if suc && i == tempEndReel {
			return true, path
		}
		// 判断结束节点是否可以抵达当前列
		find, transverseConnectPath := l.FindTransverseConnectPath(i, tempEndLine, tempEndReel, tempEndLine, true)
		if find {
			res := append(path, transverseConnectPath[1:]...)
			// 拼接开始坐标到当前 i 的坐标
			var starPath [][]int
			for j := tempStartReel; j < i; j++ {
				starPath = append(starPath, []int{j, tempStartLine})
			}
			res = append(starPath, res...)
			return true, res
		}
	}
	// 纵向查找
	var i = 0
	if tempStartLine > tempEndLine {
		i = tempStartLine - 1
	} else {
		i = tempStartLine + 1
	}
	for (i <= tempEndLine || i >= tempEndLine) && i >= 0 && i <= len(l.LinkMap[0])-1 {
		// 纵向移动过不去 直接结束
		if l.LinkMap[tempStartReel][i] != nullPosition {
			break
		}
		// 判断当前节点是否可以通行到结束节点的那一列
		suc, path := l.FindTransverseConnectPath(tempStartReel, i, tempEndReel, i, true)
		if !suc {
			if tempStartLine > tempEndLine {
				i--
			} else {
				i++
			}
			continue
		}
		// 判断和结束节点是同一个节点 是则代表已经找到了可行通路 直接返回
		if suc && i == tempEndLine {
			return true, path
		}
		// 判断结束节点是否可以抵达当前节点
		find, transverseConnectPath := l.FindVerticalConnectPath(tempEndReel, i, tempEndReel, tempEndLine, true)
		if find {
			res := append(path, transverseConnectPath[1:]...)
			// 拼接开始坐标到当前 i 的坐标
			var starPath [][]int
			for j := tempStartReel; j < i; j++ {
				starPath = append(starPath, []int{j, tempStartLine})
			}
			res = append(starPath, res...)
			return true, res
		}
		if tempStartLine > tempEndLine {
			i--
		} else {
			i++
		}
	}
	return false, nil
}

// [ ] 连法判断
func (l *LinkGame) FindHalfBorderPath(startReel, startLine, endReel, endLine int) (bool, [][]int) {
	// 优先判断是否都在同一条边界上
	// 竖向边界
	if startReel == endReel && (startReel == 0 || startReel == len(l.LinkMap)-1) {
		return true, [][]int{{startReel, startLine}, {endReel, endLine}}
	}
	// 横向边界
	if startLine == endLine && (startLine == 0 || startLine == len(l.LinkMap[0])-1) {
		return true, [][]int{{startReel, startLine}, {endReel, endLine}}
	}
	// 判断是否有一列处于边界列 边界列的话则只需要判断另外一列是否能抵达边界列的同一行的节点即可
	if startReel == 0 || startReel == len(l.LinkMap)-1 {
		suc, path := l.FindTransverseConnectPath(startReel, endLine, endReel, endLine, true)
		if suc {
			return suc, path
		}
	}
	if endReel == 0 || endReel == len(l.LinkMap)-1 {
		suc, path := l.FindTransverseConnectPath(startReel, startLine, endReel, startLine, true)
		if suc {
			return suc, path
		}
	}
	// 判断是否有一行处于边界行 边界行的话则只需要判断另外一行是否能抵达边界行的同一列的节点即可
	if startLine == 0 || startLine == len(l.LinkMap[0])-1 {
		suc, path := l.FindVerticalConnectPath(endReel, startLine, endReel, endLine, true)
		if suc {
			return suc, path
		}
	}
	if endLine == 0 || endLine == len(l.LinkMap[0])-1 {
		suc, path := l.FindVerticalConnectPath(startReel, endLine, startReel, startLine, true)
		if suc {
			return suc, path
		}
	}
	// 横竖两个方向分别判断是否可达
	suc, path := l.findVerticalHalfBorderPath(startReel, startLine, endReel, endLine)
	if suc {
		return suc, path
	}
	suc, path = l.findTransverseHalfBorderPath(startReel, startLine, endReel, endLine)
	if suc {
		return suc, path
	}
	return false, nil
}

func (l *LinkGame) findVerticalHalfBorderPath(startReel int, startLine int, endReel int, endLine int) (bool, [][]int) {
	lines := make([]int, 0)
	// 上下分别遍历 可直达的line集合
	for i := startLine; i >= 0; i-- {
		if l.LinkMap[startReel][i] != nullPosition && i != startLine {
			break
		}
		lines = append(lines, i)
	}
	for i := startLine + 1; i <= len(l.LinkMap[0])-1; i++ {
		if l.LinkMap[startReel][i] != nullPosition {
			break
		}
		lines = append(lines, i)
	}
	// 遍历lines 寻找end position 可达到的节点
	for _, line := range lines {
		if l.LinkMap[endReel][line] != nullPosition {
			continue
		}
		// 边界特殊处理
		suc, path := l.FindTransverseConnectPath(startReel, line, endReel, line, true)
		// 边界特殊处理 到达边界行 则只需要end节点能竖向直达边界行的同一列节点即可
		if !suc && line != 0 && line != len(l.LinkMap[0])-1 {
			continue
		}
		find, verticalPath := l.FindVerticalConnectPath(endReel, line, endReel, endLine, true)
		if find {
			return true, append(path, verticalPath...)
		}
	}
	return false, nil
}

func (l *LinkGame) findTransverseHalfBorderPath(startReel int, startLine int, endReel int, endLine int) (bool, [][]int) {
	reels := make([]int, 0)
	// 左右分别遍历寻找 可直达的reel集合
	for i := startReel; i >= 0; i-- {
		if l.LinkMap[i][startLine] != nullPosition && i != startReel {
			break
		}
		reels = append(reels, i)
	}
	for i := startReel + 1; i <= len(l.LinkMap)-1; i++ {
		if l.LinkMap[i][startLine] != nullPosition {
			break
		}
		reels = append(reels, i)
	}
	// 遍历reels 寻找end position 可达到的节点
	for _, reel := range reels {
		if l.LinkMap[reel][endLine] != nullPosition {
			continue
		}
		suc, path := l.FindVerticalConnectPath(reel, startLine, reel, endLine, true)
		// 边界特殊处理 到达边界列 则只需要end节点横向能直达边界列的同一行节点即可
		if !suc && reel != 0 && reel != len(l.LinkMap)-1 {
			continue
		}
		find, path2 := l.FindTransverseConnectPath(reel, endLine, endReel, endLine, true)
		if find {
			return true, append(path, path2...)
		}
	}
	return false, nil
}

func (l *LinkGame) ShuffleMap() {
	if len(l.LinkMap) < 2 {
		return
	}
	// 记录当前地图有icon的下标和icon列表
	iconIndex := make([][]int, 0)
	icons := make([]int, 0)
	for reel, lineIcons := range l.LinkMap {
		for line, icon := range lineIcons {
			if icon == nullPosition {
				continue
			}
			iconIndex = append(iconIndex, []int{reel, line})
			icons = append(icons, icon)
		}
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(icons), func(i, j int) {
		icons[i], icons[j] = icons[j], icons[i]
	})
	// 写回地图
	for index, position := range iconIndex {
		l.LinkMap[position[0]][position[1]] = icons[index]
	}
}

func (l *LinkGame) ShuffleMapWithPassablePath() {
	if len(l.LinkMap) < 2 {
		return
	}
	// 记录当前地图有icon的下标和icon列表
	iconIndex := make([][]int, 0)
	icons := make([]int, 0)
	// 随机找一个两个及以上的icon
	findFlag := false
	iconNum := make(map[int][][]int, 0)
	var sameIconPosition [][]int
	for reel, lineIcons := range l.LinkMap {
		for line, icon := range lineIcons {
			if icon == nullPosition {
				continue
			}
			iconIndex = append(iconIndex, []int{reel, line})
			icons = append(icons, icon)
			if !findFlag {
				iconNum[icon] = append(iconNum[icon], []int{reel, line})
				length := len(iconNum[icon])
				if length == 2 {
					findFlag = true
					sameIconPosition = iconNum[icon]
				}
			}
		}
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(icons), func(i, j int) {
		icons[i], icons[j] = icons[j], icons[i]
	})
	// 写回地图
	for index, position := range iconIndex {
		l.LinkMap[position[0]][position[1]] = icons[index]
	}
	if len(sameIconPosition) == 0 {
		return
	}
	// 查找可行路径 进行icon替换
	reel, line := l.tryPassablePath(sameIconPosition)
	endReel := sameIconPosition[1][0]
	endLine := sameIconPosition[1][1]
	if reel != -1 && line != -1 {
		oldIcon := l.LinkMap[reel][line]
		l.LinkMap[reel][line] = l.LinkMap[endReel][endLine]
		l.LinkMap[endReel][endLine] = oldIcon
	}
}

func (l *LinkGame) tryPassablePath(sameIconPosition [][]int) (reel, line int) {
	reel, line = -1, -1
	// 优先尝试竖向查找 如果没有 代表整列都为空的 那就整列的横线都遍历一遍
	startReel := sameIconPosition[0][0]
	startLine := sameIconPosition[0][1]
	reel, line = l.tryVerticalPath(startReel, startLine)
	if reel != -1 && line != -1 {
		return
	}
	for i := 0; i < len(l.LinkMap); i++ {
		reel, line = l.tryTransversePath(i, startLine)
		if reel != -1 && line != -1 {
			return
		}
	}
	return
}

// 横向查找可替换的icon的下标
func (l *LinkGame) tryTransversePath(startIconReel, startIconLine int) (reel, line int) {
	// 向左偏移 尝试查找是否有可达的存在icon的位置
	for i := startIconReel - 1; i >= 0; i-- {
		if l.LinkMap[i][startIconLine] != nullPosition {
			return i, startIconLine
		}
	}
	// 向右偏移
	for i := startIconReel + 1; i < len(l.LinkMap); i++ {
		if l.LinkMap[i][startIconLine] != nullPosition {
			return i, startIconLine
		}
	}
	return -1, -1
}

// 竖向查找
func (l *LinkGame) tryVerticalPath(startIconReel, startIconLine int) (reel, line int) {
	// 向上偏移 尝试查找是否有可达的存在icon的位置
	for i := startIconLine - 1; i >= 0; i-- {
		if l.LinkMap[startIconReel][i] != nullPosition {
			return startIconReel, i
		}
	}
	// 向下偏移
	for i := startIconLine + 1; i < len(l.LinkMap[0]); i++ {
		if l.LinkMap[startIconReel][i] != nullPosition {
			return startIconReel, i
		}
	}
	return -1, -1
}

// 查找是否有解
func (l *LinkGame) CheckHasPassablePath() bool {
	// 找出相同icon的下标集合
	// 过程中如果不是处于边界的icon可以尝试判断上下左右四个方向是否存在icon,都存在且都是该icon，那代表这个位置已经被堵死了，不需要考虑通路了；存在且是同一个icon 证明找到了解 直接return
	sameIconIndex := make(map[int][][]int, 8)
	for reel, lineIcons := range l.LinkMap {
		for line, icon := range lineIcons {
			if reel == 0 || line == 0 || reel == len(l.LinkMap) || line == len(l.LinkMap[0]) {
				sameIconIndex[icon] = append(sameIconIndex[icon], []int{reel, line})
				continue
			}
			if l.LinkMap[reel-1][line] == icon || l.LinkMap[reel+1][line] == icon || l.LinkMap[reel][line-1] == icon || l.LinkMap[reel][line+1] == icon {
				return true
			}
			if l.LinkMap[reel-1][line] == nullPosition || l.LinkMap[reel+1][line] == nullPosition || l.LinkMap[reel][line-1] == nullPosition || l.LinkMap[reel][line+1] == nullPosition {
				sameIconIndex[icon] = append(sameIconIndex[icon], []int{reel, line})
			}
		}
	}
	for _, indexes := range sameIconIndex {
		if len(indexes) < 2 {
			continue
		}
		for i, index := range indexes {
			for j := i + 1; j < len(indexes); j++ {
				find, _ := l.FindPassablePath(index[0], index[1], indexes[j][0], indexes[j][1])
				if find {
					return true
				}
			}
		}
	}
	return false
}
