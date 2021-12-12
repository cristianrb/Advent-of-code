import java.io.File
import java.util.*

fun main() {
    val input = File("src/sample")
        .readLines()
        .map { it.split("-") }
        .flatMap { (begin, end) -> listOf(begin to end, end to begin) }
        .groupBy({ it.first }, { it.second })

    val graph = Graph(input)
    graph.dfs("start", "end")
}

class Graph(private val adjacencyMap: Map<String, List<String>>) {
    private var pathsCount = 0

    fun dfs(start: String, end: String) {
        val visited = hashMapOf<String, Boolean>()

        val path = arrayListOf<String>()
        dfsRec(visited, start, end, path, null)
        println("Part 1: $pathsCount")

        visited.clear()
        path.clear()
        pathsCount = 0
        dfsRec( visited, start, end, path, false)
        println("Part 2: $pathsCount")
    }

    private fun dfsRec(visited: HashMap<String, Boolean>, node: String, end: String, path: ArrayList<String>, smallVisitedTwice: Boolean?) {
        visited[node] = true
        path.add(node)

        if (node == end) {
            pathsCount++
            return
        }

        for (nextNode in adjacencyMap[node]!!) {
            if (nextNode == "start") continue

            val visitedCave = visited.getOrDefault(nextNode, false)
            if (nextNode.isLowerCase() && visitedCave) {
                if (smallVisitedTwice != null && !smallVisitedTwice) {
                    dfsRec(visited, nextNode, end, path, true)
                }
                continue
            }
            dfsRec(visited, nextNode, end, path, smallVisitedTwice)
            path.removeLast()
            visited[nextNode] = false
        }
    }
}

fun String.isLowerCase(): Boolean {
    this.forEach {
        if (Character.isUpperCase(it)) {
            return false
        }
    }
    return true
}