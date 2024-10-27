package graph

type Graph interface {
	GetNoEdgeValue() int
	SetNoEdgeValue(int)
	GetVertexCount() int
	GetEdgeCount() int
	GetAllEdges() []Edge
	GetEdgesFromVertex(startVertex int) []Edge
	GetEdgesToVertex(endVertex int) []Edge
	GetEdge(startVertex, endVertex int) Edge
	AddEdge(startVertex, endVertex, weight int)
	RemoveEdge(startVertex, endVertex int)
	IsAdjacent(startVertex, endVertex int) bool
	CalculatePathWeight(path []int) int
	ToString() string
}
