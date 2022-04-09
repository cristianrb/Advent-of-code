import java.io.File

fun main() {
    val input = File("src/input")
        .readLines()
        .map{ it.map(Character::getNumericValue) }

    solve(input)
}

fun solve(matrix: List<List<Int>>) {
    val lowPoints = arrayListOf<Int>()
    val basins = arrayListOf<Int>()
    for (i in matrix.indices) {
        for (j in matrix[0].indices) {
            if (isLowPoint(i, j, matrix)) {
                lowPoints.add(matrix[i][j])
                basins.add(findBasins(matrix, Pair(i, j)))
            }
        }
    }
    println(lowPoints.sumOf { it+1 })
    basins.sortDescending()
    println(basins[0] * basins[1] * basins[2])
}

private fun isLowPoint(i: Int, j: Int, matrix: List<List<Int>>): Boolean {
    var left = Int.MAX_VALUE
    var right = Int.MAX_VALUE
    var top = Int.MAX_VALUE
    var bottom = Int.MAX_VALUE
    if (i - 1 >= 0) left = matrix[i - 1][j]
    if (i + 1 < matrix.size) right = matrix[i + 1][j]
    if (j - 1 >= 0) top = matrix[i][j - 1]
    if (j + 1 < matrix[0].size) bottom = matrix[i][j + 1]
    val actual = matrix[i][j]
    if (actual < left && actual < right && actual < top && actual < bottom) {
        return true
    }
    return false
}

fun findBasins(matrix: List<List<Int>>, pair: Pair<Int, Int>) : Int {
    val basins = hashSetOf<Pair<Int, Int>>()
    val candidates = arrayListOf<Pair<Int, Int>>()
    candidates.add(pair)
    while (candidates.isNotEmpty()) {
        val pair = candidates.removeFirst()
        val i = pair.first
        val j = pair.second
        val actual = matrix[i][j]
        basins.add(pair)
        if (i-1 >= 0 && matrix[i-1][j] != 9) {
            if (actual<matrix[i-1][j]) candidates.add(Pair(i-1, j))
        }
        if (i+1 < matrix.size && matrix[i+1][j] != 9) {
            if (actual<matrix[i+1][j]) candidates.add(Pair(i+1, j))
        }
        if (j-1 >= 0 && matrix[i][j-1] != 9) {
            if (actual<matrix[i][j-1]) candidates.add(Pair(i, j-1))
        }
        if (j+1 < matrix[0].size && matrix[i][j+1] != 9) {
            if (actual<matrix[i][j+1]) candidates.add(Pair(i, j+1))
        }
    }
    return basins.size
}