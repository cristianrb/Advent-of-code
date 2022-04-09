import java.io.File

fun main() {
    val points = arrayListOf<Pair<Int, Int>>()
    val folds = arrayListOf<String>()

    File("src/input")
        .readLines()
        .forEach {
            if (it.contains("fold")) folds.add(it)
            else if (it.isNotEmpty()) {
                val point = it.split(",")
                points.add(Pair(point[0].toInt(), point[1].toInt()))
            }
        }

    val matrix = createMat(points)
    println("Part 1: " + solve(matrix, folds, true))
    println("Part 2:")
    solve(matrix, folds, false)
}

private fun createMat(points: ArrayList<Pair<Int, Int>>): Array<CharArray> {
    val maxX = points.maxByOrNull { it.second }!!.second + 1
    val maxY = points.maxByOrNull { it.first }!!.first + 1

    val matrix = Array(maxX) { CharArray(maxY) { '.' } }
    points.forEach { matrix[it.second][it.first] = '#' }
    return matrix
}

fun solve(matrix: Array<CharArray>, folds: ArrayList<String>, stopEarly: Boolean): Int {
    var rowSize = matrix.size
    var colSize = matrix[0].size
    for (fold in folds) {
        if (fold.contains("fold along y=")) {
            rowSize = foldY(fold, matrix)
        } else if (fold.contains("fold along x=")) {
            colSize = foldX(fold, matrix)
        }
        if (stopEarly)
            return matrix.countChar('#')
    }

    printMat(matrix, rowSize, colSize)
    return matrix.countChar('#')
}

private fun foldX(fold: String, matrix: Array<CharArray>): Int {
    val maxX = fold.split("fold along x=")[1].toInt()
    for (i in matrix.indices) {
        for (j in matrix[0].indices) {
            if (j < maxX) {
                continue
            }
            val newJ = maxX - (j - maxX)
            if (matrix[i][j] != '.') matrix[i][newJ] = matrix[i][j]
            matrix[i][j] = '.'
        }
    }
    return maxX
}

private fun foldY(fold: String, matrix: Array<CharArray>): Int {
    val maxY = fold.split("fold along y=")[1].toInt()
    for (i in matrix.indices) {
        if (i < maxY) {
            continue
        }
        for (j in matrix[0].indices) {
            val newI = maxY - (i - maxY)
            if (matrix[i][j] != '.') matrix[newI][j] = matrix[i][j]
            matrix[i][j] = '.'
        }
    }
    return maxY
}

fun Array<CharArray>.countChar(character: Char): Int {
    var counter = 0
    map {
        it.map {
            if (it == character) counter++
        }
    }
    return counter
}

private fun printMat(matrix: Array<CharArray>, rowSize: Int, colSize: Int) {
    for (row in 0 until rowSize) {
        for (col in 0 until colSize) {
            print("${matrix[row][col]} ")
        }
        println()
    }
}