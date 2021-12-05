import java.io.File
import kotlin.math.abs

fun main() {
    val input = File("src/input")
        .useLines { it.toList() }

    val diagram1 = createDiagramFromListPartOne(input)
    val diagram2 = createDiagramFromListPartTwo(input)
    val res = calculateOverlappingLines(diagram1)
    val res2 = calculateOverlappingLines(diagram2)
    println(res)
    println(res2)
}

fun calculateOverlappingLines(diagram: Array<IntArray>): Int {
    var count = 0
    for (i in diagram.indices) {
        for (j in diagram[0].indices) {
            if (diagram[i][j] > 1) count++
        }
    }
    return count
}

fun createDiagramFromListPartOne(input: List<String>): Array<IntArray> {
    val mat = Array(findMaxCoordFromAxis(input, 0)) { IntArray(findMaxCoordFromAxis(input, 1))}
    input.forEach {
        val coords = it.split("->")
        val pair0 = coords[0].split(",")
        val pair1 = coords[1].split(",")
        var x0 = pair0[0].toInt()
        var y0 = pair0[1].trim().toInt()
        val x1 = pair1[0].trim().toInt()
        val y1 = pair1[1].toInt()

        if (x0 == x1) {
            while ( (abs(y1-y0) >= 0) && (y0 != -1 && y1 != -1)  ) {
                mat[y0][x0]++
                y0 = increaseOrDecreaseCoord(y0, y1)
            }
        } else if (y0 == y1) {
            while ( (abs(x1-x0) >= 0) && (x0 != -1 && x1 != -1)  ) {
                mat[y0][x0]++
                x0 = increaseOrDecreaseCoord(x0, x1)
            }
        }
    }

    return mat
}

fun createDiagramFromListPartTwo(input: List<String>): Array<IntArray> {
    val mat = Array(findMaxCoordFromAxis(input, 0)) { IntArray(findMaxCoordFromAxis(input, 1))}
    input.forEach {
        val coords = it.split("->")
        val pair0 = coords[0].split(",")
        val pair1 = coords[1].split(",")
        var x0 = pair0[0].toInt()
        var y0 = pair0[1].trim().toInt()
        val x1 = pair1[0].trim().toInt()
        val y1 = pair1[1].toInt()

        var end = false
        while ( (abs(x1-x0) >= 0 || abs(y1-y0) >= 0) && (!end)  ) {
            mat[y0][x0]++
            if ((x0 == x1 && x0 == y0 && x0 == y1) || (x0 == x1 && y0 == y1)) {
                end = true
            } else {
                if (x0 != x1) x0 = increaseOrDecreaseCoord(x0, x1)
                if (y0 != y1) y0 = increaseOrDecreaseCoord(y0, y1)
            }
        }
    }

    return mat
}

fun findMaxCoordFromAxis(input: List<String>, axis: Int): Int {
    var max = 0
    input.forEach {
        val coords = it.split("->")
        val pair0 = coords[0].split(",")
        val pair1 = coords[1].split(",")
        val n0 = pair0[axis].trim().toInt()
        val n1 = pair1[axis].trim().toInt()
        if (n0 > max) max = n0
        if (n1 > max) max = n1
    }
    return max + 1
}

private fun increaseOrDecreaseCoord(n0: Int, n1: Int): Int {
    var tmp = n0
    when {
        tmp > n1 -> tmp--
        tmp < n1 -> tmp++
        else -> tmp = -1
    }
    return tmp
}
