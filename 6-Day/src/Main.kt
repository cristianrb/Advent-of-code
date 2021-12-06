import java.io.File

fun main() {
    val input = File("src/input")
        .readText()
        .split(Regex("\\D+"))
        .map { it.toInt() }
        .groupBy { it }
        .map { (k, v) -> k to v.size.toLong() }
        .toMap()

    val res = calculateFishes(input.toMutableMap(), 80)
    println(res)
    val res2 = calculateFishes(input.toMutableMap(),256)
    println(res2)
}

fun calculateFishes(input: MutableMap<Int, Long>, maxDays: Int): Long {
    repeat(maxDays) {
        val newHashMap = mutableMapOf<Int, Long>()
        input.forEach { (key, value) ->
            if (key == 0) {
                newHashMap[8] = newHashMap[8]?.plus(value) ?: value
                newHashMap[6] = newHashMap[6]?.plus(value) ?: value
            } else {
                newHashMap[key - 1] = newHashMap[key - 1]?.plus(value) ?: value
            }
        }
        input.clear()
        input.putAll(newHashMap)
    }

    return input.values.sum()
}
