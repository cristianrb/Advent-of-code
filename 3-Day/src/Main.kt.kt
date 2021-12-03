import java.io.File

fun main() {
    val input = File("src/input")
        .useLines { it.toList() }

    partOne(input)
    partTwo(input)
}

private fun partOne(input : List<String>) {
    val matrix = convertListToMatrix(input)
    val res = findPowerConsumption(matrix)
    println(res)
}

private fun convertListToMatrix(input: List<String>) : Array<IntArray> {
    val matrix = Array(input.size) { IntArray(input[0].length) }
    for( i in input.indices) {
        matrix[i] = input[i].map{ it.toString().toInt() }.toIntArray()
    }
    return matrix
}

private fun findPowerConsumption(matrix: Array<IntArray>): Int {
    var gammaRate = ""
    var epsilonRate = ""

    for (i in matrix[0].indices) {
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
    return gammaRateDec * epsilonRateDec
}

private fun partTwo(input: List<String>) {
    val compFunc1 = { a: List<String>, b: List<String> -> a.size >= b.size }
    val compFunc2 = { a: List<String>, b: List<String> -> b.size > a.size }
    val oxygenRating = findRating(input, 0, compFunc1)
    val co2Rating = findRating(input, 0, compFunc2)
    println(oxygenRating * co2Rating)
}

private fun findRating(input: List<String>, pos: Int, comparison: (List<String>, List<String>) -> Boolean) : Int{
    if (input.size == 1) return Integer.parseInt(input[0], 2)

    val ones = input.filter { it[pos].toString().toInt() == 1 }
    val zeros =input.filter { it[pos].toString().toInt() == 0 }

    val newPos = pos + 1

    if (comparison(ones, zeros)) return findRating(ones, newPos, comparison)
    return findRating(zeros, newPos, comparison)
}