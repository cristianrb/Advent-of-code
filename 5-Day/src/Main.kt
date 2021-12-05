import java.io.File
import kotlin.math.abs
import kotlin.math.max
import kotlin.math.sign

private val onlyDigits = Regex("\\D+")

data class Line(val coords: List<Pair<Int, Int>>, val notDiagonal: Boolean)

fun main() {
    val input = File("src/input")
        .useLines { it.toList() }
        .map { createLine(it) }

    println(partOne(input))
    println(partTwo(input))
}

private fun createLine(str: String) : Line {
    return str.trim()
        .split(onlyDigits)
        .map { it.toInt() }
        .let { (x0, y0, x1, y1) ->
            val xLen = abs(x1 - x0)
            val yLen = abs(y1 - y0)
            val xSign = (x1-x0).sign
            val ySign = (y1 - y0).sign

            val points = 0..max(xLen, yLen)
            val coords = points.map { n -> Pair(x0 + n * xSign, y0 + n * ySign) }
            val notDiagonal = xLen == 0 || yLen == 0

            Line(coords, notDiagonal)
        }
}

fun partOne(input: List<Line>): Int {
    return calcOverlapping(input.filter { it.notDiagonal })
}

fun partTwo(input: List<Line>): Int {
    return calcOverlapping(input)
}

private fun calcOverlapping(input: List<Line>): Int {
    return input
        .flatMap { it.coords }
        .groupBy { it }
        .filter { it.value.size > 1}
        .size
}
