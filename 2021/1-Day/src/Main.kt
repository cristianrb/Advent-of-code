import java.io.File

fun main() {
    val input = File("src/input")
        .useLines { it.toList() }
        .map { it.toInt() }

    partOne(input)
    partTwo(input)
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

private fun partTwo(input: List<Int>) {
    var convertedInput = arrayListOf<Int>()

    for (i in input.indices) {
        if ((input.size - i) > 2) {
            convertedInput.add(input[i] + input[i + 1] + input[i + 2])
        }
    }
    partOne(convertedInput)
}