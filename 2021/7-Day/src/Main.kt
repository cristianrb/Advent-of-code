import java.io.File
import kotlin.math.abs
import kotlin.math.min

fun main() {
    val input = File("src/input")
        .readText()
        .split(Regex("\\D+"))
        .map { it.toLong() }

    calcMinFuel(input)
}

private fun sumOf1ToN(n: Long): Long {
    return n * (n+1) / 2
}

fun calcMinFuel(input: List<Long>) {
    var fuel = Long.MAX_VALUE
    var fuel2 = Long.MAX_VALUE

    for (i in input.indices) {
        fuel = min(input.sumOf { abs(it-i) }, fuel)
        fuel2 = min(input.sumOf { sumOf1ToN(abs(it-i))}, fuel2)
    }
    
    println(fuel)
    println(fuel2)
}
