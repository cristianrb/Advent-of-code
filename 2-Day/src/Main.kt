import java.io.File

fun main() {
    val input = File("src/input")
        .useLines { it.toList() }

    partOne(input)
    partTwo(input)
}

private fun partOne(input: List<String>) {
    var hPosition = 0
    var depth = 0
    input.map {
        val command = it.split(" ")
        when (command[0]) {
            "forward" -> hPosition += command[1].toInt()
            "up" -> depth -= command[1].toInt()
            "down" -> depth += command[1].toInt()
        }
    }
    println(hPosition * depth)
}

private fun partTwo(input: List<String>) {
    var hPosition = 0
    var depth = 0
    var aim = 0
    input.map {
        val command = it.split(" ")
        when (command[0]) {
            "forward" -> {
                hPosition += command[1].toInt()
                depth += command[1].toInt() * aim
            }
            "up" -> aim -= command[1].toInt()
            "down" -> aim += command[1].toInt()
        }
    }
    println(hPosition * depth)
}