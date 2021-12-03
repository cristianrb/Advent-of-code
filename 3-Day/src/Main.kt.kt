import java.io.File

fun main() {
    val input = File("src/input")
        .useLines { it.toList() }

    val matrix = convertListToMatrix(input)
    findPowerConsumption(matrix)
}

private fun convertListToMatrix(input: List<String>) : Array<IntArray> {
    var matrix = Array(input.size) { IntArray(input[0].length) }
    for( i in input.indices) {
        matrix[i] = input[i].map{ it.toString().toInt() }.toIntArray()
    }
    return matrix
}

private fun findPowerConsumption(matrix: Array<IntArray>) {
    var gammaRate = ""
    var epsilonRate = ""

    for (i in matrix.reversedArray().indices) {
        var ones = 0
        var zeros = 0
        for (j in matrix.indices) {
            if (matrix[j][i] == 0) zeros++
            else ones++
        }
        if (zeros > ones) {
            gammaRate += "0"
            epsilonRate += "1"
        } else {
            gammaRate += "1"
            epsilonRate += "0"
        }
    }
    val gammaRateDec = Integer.parseInt(gammaRate, 2)
    val epsilonRateDec = Integer.parseInt(epsilonRate, 2)
    var result = gammaRateDec * epsilonRateDec
    println(result)
}

private fun partOne(input: List<Int>) {
    var previousLine = 0
    var increases = 0
    input.forEach {
        if (it > previousLine && previousLine != 0) {
            increases++
        }
        previousLine = it
    }
    println(increases)
}