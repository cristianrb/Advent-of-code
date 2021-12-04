import java.io.File

fun main() {
    val input = File("src/input").useLines { it.toMutableList() }

    val bingoNumbers = input.removeAt(0)
    val matrixs = ArrayList<Array<IntArray>>(0)
    var currentMat = -1
    var row = 0
    var col = 0
    val regex = "\\s+".toRegex()
    input.forEach {
        val str = it.trim()
        if (str != "") {
            str.split(regex).forEach {
                matrixs[currentMat][row][col] = it.toInt()
                col++
            }
            row++
            col = 0
        } else {
            matrixs.add(Array(5) { IntArray(5) })
            currentMat++
            row = 0
            col = 0
        }
    }

    val resOne = playBingo(bingoNumbers.split(",").map { it.toInt() }, matrixs, true)
    val resTwo = playBingo(bingoNumbers.split(",").map { it.toInt() }, matrixs, false)
    println(resOne)
    println(resTwo)
}

fun checkBingo(matrix: Array<IntArray>): Boolean {
    var winH = true
    var winV = true
    for (i in matrix.indices) {
        winH = true
        winV = true
        for (j in matrix[0].indices) {
            if (matrix[i][j] != -1) {
                winH = false
            }
            if (matrix[j][i] != -1) {
                winV = false
            }
        }
        if (winH || winV) return true
    }
    return winH || winV
}

fun calculateRes(matrix: Array<IntArray>, num: Int): Int {
    var sum = 0
    for (i in matrix.indices) {
        for (j in matrix[0].indices) {
            if (matrix[i][j] != -1) {
                sum += matrix[i][j]
            }
        }
    }
    return sum * num
}

fun playBingo(bingoNumbers: List<Int>, matrixs: ArrayList<Array<IntArray>>, stopOnWin: Boolean): Int {
    var bingos = IntArray(matrixs.size)
    bingoNumbers.forEach { num ->
        matrixs.forEachIndexed { index, actualMat ->
            for (i in actualMat.indices) {
                for (j in actualMat[0].indices) {
                    if (actualMat[i][j] == num) {
                        actualMat[i][j] = -1
                        if (checkBingo(actualMat)) {
                            bingos[index]++
                            if (stopOnWin || bingos.none { it == 0 })
                                return calculateRes(actualMat, num)
                        }
                    }
                }
            }
        }
    }
    return -1
}
