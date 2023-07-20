package game

// a 星算法 也是图的一个遍历找路径方法
// 吸取了Dijkstra 算法中的cost_so_far，为每个边长设置权值，不停的计算每个顶点到起始顶点的距离，以获得最短路线
// 同时也汲取贪婪最佳优先搜索算法中不断向目标前进优势，并持续计算每个顶点到目标顶点的距离，以引导搜索队列不断想目标逼近，从而搜索更少的顶点，保持寻路的高效

// 核心思想：计算每一个节点的综合优先级 选择下一个要遍历的节点时 永远选取综合优先级最好的节点 直到找到终点 或者 没有找到可通行路径
// 三大变量：
// 1，g(n)：节点n 距离起点的距离
// 2，h(n)：节点n 距离终点的预计距离 也就是a*的启发函数 这里预计是因为是预估的一个理想距离 中途有可能会碰到障碍物之类的
// 3，f(n)：节点n 的综合优先级 g(n) + h(n)
// 总结：在遍历的过程中，通过计算节点与起点的距离 和 节点与终点的理想距离 得出一个分数，永远取分数最低的，也就是代价最小的那个节点往下走，直到结束
// 注：启发函数的话需要根据当前可走的方向来选择，只允许朝上下左右四个方向移动 -> 使用曼哈顿距离 ；允许朝八个方向移动 -> 对角距离；任意方向 -> 欧几里得距离

// 步骤如下：
// 1，初始化open_set和close_set；
// 2，如果open_set不为空，则从open_set中选取优先级最高（f 分数最低）的节点n：
// 2.1，该节点如果是终点，则从终点开始追踪 parent 节点，一直到起点，形成路径，结束
// 2.2，遍历邻居节点
// 2.2.1，如果该邻居节点在 close_set 中，跳过
// 2.2.2，该节点不在 open_set 中，设置parent为节点n，计算优先级，放入open_set 中
// 2.3.3，该节点已经存在 open_set 中，计算从当前节点n 移动到邻居的成本，也就是从起点经过当前节点到邻居节点的成本，如果比邻居节点之前的G小，则更新 G 值和父亲节点为当前节点n

// 参考论文：https://www.gamedev.net/reference/articles/article2003.asp

// 网格中的 坐标类型
type Vector struct {
	X int
	Y int
}

// ManhattanDistance 计算两个节点之中的曼哈顿距离 因为本次的想法实现固定是四个方向 所以使用曼哈顿距离 也就是直接累加两个节点的 x y 轴的距离
func (v Vector) ManhattanDistance(next Vector) int {
	return abs(v.X-next.X) + abs(v.Y-next.Y)
}

func (v Vector) Add(other Vector) Vector {
	return Vector{v.X + other.X, v.Y + other.Y}
}

func (v Vector) Equal(other Vector) bool {
	return v.X == other.X && v.Y == other.Y
}

// 取绝对值
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//上 左 下 右 四个方向
var dirs = [4]Vector{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

type Grid [][]int

// Neighbors 根据设定的方向 计算出邻居节点
func (g Grid) Neighbors(pos Vector) []Vector {
	var neighbors []Vector
	for _, dir := range dirs {
		neighbor := pos.Add(dir)
		if g.IsValidPos(neighbor) && g[neighbor.Y][neighbor.X] != 1 {
			neighbors = append(neighbors, neighbor)
		}
	}
	return neighbors
}

// IsValidPos 判断是否越界
func (g Grid) IsValidPos(pos Vector) bool {
	return pos.X >= 0 && pos.X < len(g[0]) && pos.Y >= 0 && pos.Y < len(g)
}

type Node struct {
	// 坐标点
	Position Vector
	// 与起点的距离
	G int
	// 与终点的理想距离
	H int
	// 优先级 从起始节点经过该节点到达结束节点的总成本
	F int
	// 父亲节点
	Parent *Node
}

type AStar struct {
	Grid     Grid
	Start    *Node
	End      *Node
	OpenSet  map[Vector]*Node
	CloseSet map[Vector]*Node
}

func NewAStar(grid Grid, start, end Vector) *AStar {
	startNode := &Node{
		Position: start,
		G:        0,
		H:        start.ManhattanDistance(end),
		F:        start.ManhattanDistance(end),
	}
	endNode := &Node{
		Position: end,
	}
	return &AStar{
		Grid:     grid,
		Start:    startNode,
		End:      endNode,
		OpenSet:  map[Vector]*Node{start: startNode},
		CloseSet: map[Vector]*Node{},
	}
}

func (a *AStar) GetLowestFScoreNode() *Node {
	var lowest *Node
	for _, node := range a.OpenSet {
		if lowest == nil || node.F < lowest.F {
			lowest = node
		}
	}
	return lowest
}

func (a *AStar) Search() []Vector {
	for len(a.OpenSet) > 0 {
		// 选取优先级最高的那个节点 也就是 F 值最低的那个
		lowestFScoreNode := a.GetLowestFScoreNode()

		// 判断是否是终点
		if lowestFScoreNode.Position.Equal(a.End.Position) {
			return a.ReconstructPath(lowestFScoreNode)
		}

		// 从 open_set 删除 并加入close_set
		delete(a.OpenSet, lowestFScoreNode.Position)
		a.CloseSet[lowestFScoreNode.Position] = lowestFScoreNode

		// 获取所有邻居节点
		neighbors := a.Grid.Neighbors(lowestFScoreNode.Position)
		for _, neighbor := range neighbors {
			// 在 close_set 中 跳过
			if _, ok := a.CloseSet[neighbor]; ok {
				continue
			}
			// 计算从当前节点移动到邻居的成本 这一步是为了计算从起点经过当前节点 到 该邻居节点的成本
			cost := lowestFScoreNode.G + lowestFScoreNode.Position.ManhattanDistance(neighbor)

			v, ok := a.OpenSet[neighbor]
			if !ok || cost < v.G {
				node := &Node{
					Position: neighbor,
					G:        cost,
					H:        neighbor.ManhattanDistance(a.End.Position),
					F:        cost + neighbor.ManhattanDistance(a.End.Position),
					Parent:   lowestFScoreNode,
				}
				a.OpenSet[neighbor] = node
			}
		}
	}
	return nil
}

// ReconstructPath 回溯路径
func (a *AStar) ReconstructPath(end *Node) []Vector {
	path := []Vector{end.Position}
	current := end
	for current.Parent != nil {
		path = append(path, current.Parent.Position)
		current = current.Parent
	}
	// 路径反转 因为是从终点开始回溯的 所以需要反转一下
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}
