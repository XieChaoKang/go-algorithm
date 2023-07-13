package game

import "fmt"

// 广度版本的迷宫根据出口寻路算法

// 广度算法 BFS 是一种盲目搜寻法，目的是系统地展开并检查图中的所有节点，以找寻结果。换句话说，它并不考虑结果的可能位址，彻底地搜索整张图，直到找到结果为止，步骤如下：
// 1，根节点放入队列中
// 2，循环 取出队列中的节点，判断是否为目标节点，找到目标则结束并回传结果；否则将所有还没检验过的直接字节点（四周的邻节点）加入队列中
// 3，直到队列为空

// 走迷宫比普通的广度多了一些东西，需要记录每一步后所消耗的步数 形成一个步数地图，最终从终点开始回溯到起点，形成一条可行路径

//定义坐标类型，不过坐标类型需要依照二维数组下标的规则
type point struct {
	i, j int
}

//上 左 下 右 四个方向
var directions = [4]point{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

//计算点计算后的结果
func (p point) addPoint(dir point) point {
	//故意不设置为引用传递，这样不需要重新创建一个point返回
	p.i += dir.i
	p.j += dir.j
	return p
}

//判断该点是否越界以及返回该点的值为多少
func (p point) normalGetValue(myMap [][]int) (int, bool) {
	//先比行
	if p.i < 0 || p.i >= len(myMap) {
		return -1, false
	}

	//再比列
	if p.j < 0 || p.j >= len(myMap[0]) {
		return -1, false
	}

	//返回具体的值
	return myMap[p.i][p.j], true
}

//打印地图
func printMap(myMap [][]int) {
	fmt.Println("开始打印地图:")
	for _, value := range myMap {
		for _, val := range value {
			fmt.Printf("%d\t", val)
		}
		fmt.Println()
	}
}

// TryGoEndPoint 开始尝试走到终点
// @params: maze：迷宫地图 ；start：起点；end：终点
// @return: step：步数地图 记录走向迷宫的每一点所需步数
func TryGoEndPoint(maze [][]int, start point, end point) [][]int {
	steps := make([][]int, len(maze))
	// 构建步数地图 记录每一个点需要消耗的步数
	for i, ints := range maze {
		temp := make([]int, len(ints))
		for index, _ := range ints {
			temp[index] = 0
		}
		steps[i] = temp
	}
	// 创建队列 将起点也就是根节点放入队列中
	queue := []point{start}
	for {
		if len(queue) == 0 {
			//如果队列取完了，证明已经结束了，退出程序
			break
		}
		first := queue[0]
		queue = queue[1:]

		findEnd := false
		// 从预设好的可行方向开始遍历直接子节点（邻节点）
		for _, direction := range directions {
			next := first.addPoint(direction)
			// 判断是否是终点
			if next == end {
				steps[next.i][next.j] = steps[first.i][first.j] + 1
				findEnd = true
				break
			}
			//判断该点是否符合要求
			//前提都需要该点没有越界
			//1.判断在maze地图中是否为1，1是墙就跳过
			res, ok := next.normalGetValue(maze)
			if !ok || res == 1 {
				continue
			}

			//2.判断在steps地图中是否已经走过了，已经走过证明不通或者别的点会走，直接跳过
			res, ok = next.normalGetValue(steps)
			if !ok || res != 0 {
				continue
			}
			//3.判断是不是起点，如果是起点就跳过
			if next == start {
				continue
			}
			//如果全部通过，证明可以走
			//把这个点加入队列，可以继续探测
			//并且将步数进行记录
			queue = append(queue, next)
			steps[next.i][next.j] = steps[first.i][first.j] + 1
		}
		if findEnd {
			break
		}
	}
	return steps
}

// LookBackPath 开始尝试回溯路径
// @params: steps：步数地图 ；start：起点；end：终点
// @return: 路径的点数组 注意：这里是从终点开始回溯 所以路径的顺序也是从终点倒退回去的
func LookBackPath(steps [][]int, start point, end point) ([]point, bool) {
	var queue []point
	endValue := steps[end.i][end.j]
	if endValue == 0 {
		//证明没有到达终点
		return queue, false
	}
	//如果找到了,先将终点存入路径
	queue = append(queue, end)
	for {
		if endValue < 0 {
			break
		}
		// endValue 代表我们从上一个节点回退到这个节点的对应步数值
		endValue--
		// 一样 遍历直接字节点寻找应该回退到的节点
		for _, direction := range directions {
			next := end.addPoint(direction)
			// 判断该点是否越界和是否为我们所寻找的点
			value, flag := next.normalGetValue(steps)
			if !flag || value != endValue {
				continue
			}
			// 步数为0的情况下需要判断该点是否为起点 确保回到起点的倒数一步没有问题
			if value == 0 && (next.i != start.i || next.j != start.j) {
				continue
			}
			// 该点符合要求 记录进路径中 同时更新回溯的起点
			queue = append(queue, next)
			end = next
			break
		}
	}
	return queue, true
}

// @description 查找起点到终点的可达路径 二维数组的迷宫
// @param maze 二维迷宫地图 0 代表可以通过 1 代表不可以通过
// @param startRow 起点行坐标
// @param startCol 起点列坐标
// @param endRow 终点行坐标
// @param endCol 终点列坐标
// @return []point 路径 bool 是否可达
func GetPathInMaze(maze [][]int, startRow, startCol, endRow, endCol int) ([]point, bool) {
	if len(maze) == 0 {
		return nil, false
	}
	startPoint := point{
		i: startRow,
		j: startCol,
	}
	endPoint := point{
		i: endRow,
		j: endCol,
	}
	steps := TryGoEndPoint(maze, startPoint, endPoint)
	path, b := LookBackPath(steps, startPoint, endPoint)

	return path, b
}

// @description 查找起点到终点的可达路径 一维数组的迷宫
// @param maze 一维迷宫地图 0 代表可以通过 1 代表不可以通过
// @param row, col 行列数量 startIndex 起点下标 endIndex 终点下标
// @return []int 路径坐标 bool 是否可达
func GetPathInOneDimensionalMaze(maze []int, row, col int, startIndex int, endIndex int) ([]int, bool) {
	mazeLength := len(maze)
	if mazeLength == 0 || mazeLength%row != 0 || mazeLength%col != 0 {
		return nil, false
	}
	newMaze := make([][]int, row)
	for i := 0; i < row; i++ {
		temp := make([]int, col)
		for j := 0; j < col; j++ {
			index := i*col + j
			if index >= len(maze) {
				return nil, false
			}
			temp[j] = maze[index]
		}
		newMaze[i] = temp
	}
	startRow := startIndex / col
	startCol := startIndex % col
	endRow := endIndex / col
	endCol := endIndex % col
	points, b := GetPathInMaze(newMaze, startRow, startCol, endRow, endCol)
	if !b {
		return nil, false
	}
	var res []int
	for _, p := range points {
		// 剔除起点
		if p.i == startRow && p.j == startCol {
			continue
		}
		res = append(res, p.i*col+p.j)
	}
	// 通过回溯找的路径 所以需要倒过来才是起点到终点
	var result []int
	for i := len(res) - 1; i >= 0; i-- {
		result = append(result, res[i])
	}
	return result, b
}
