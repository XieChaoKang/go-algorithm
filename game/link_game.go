package game

var nullPosition = 0

// 连连看游戏
type LinkGame struct {
	LinkMap [][]int // 0 代表没有东西 可以通行； 其他内容都代表有东西 无法通行
}

func (l *LinkGame) FindPassablePath(startReel, startLine, endReel, endLine int) (bool, [][]int) {
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
		if l.LinkMap[path[0]][path[1]] != 0 {
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
		res := append(path, transverseConnectPath[1:]...)
		// 拼接开始坐标到当前 i 的坐标
		var starPath [][]int
		for j := tempStartReel; j < i; j++ {
			starPath = append(starPath, []int{j, tempStartLine})
		}
		res = append(starPath, res...)
		if find {
			return true, res
		}
	}
	return false, nil
}
