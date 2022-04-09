import java.io.File

fun main() {
    val input = File("src/input")
        .readLines()
        .map { it.map(Character::toString) }

    solve(input)
}

fun solve(input: List<List<String>>) {
    val openers = listOf("(", "{", "[", "<")
    val closers = listOf(")", "}", "]", ">")
    var score = 0
    val closingScores = arrayListOf<Long>()
    for (list in input) {
        val stack = ArrayDeque<String>()
        var invalid = false
        for (str in list) {
            if (openers.contains(str)) stack.add(str)
            else if (closers.contains(str)) {
                val removedItem = stack.removeLast()
                if (openers.indexOf(removedItem) != closers.indexOf(str)) {
                    score += calcPoints(str)
                    invalid = true
                }
            }
        }
        if (!invalid) {
            closingScores.add(findClosingScore(stack))
        }
    }

    closingScores.sort()

    println(score)
    println( closingScores[(closingScores.size - 1) / 2] )
}

fun findClosingScore(stack: ArrayDeque<String>): Long {
    var score = 0L
    val scoreList = mapOf("(" to 1, "[" to 2, "{" to 3, "<" to 4)
    while (stack.isNotEmpty()) {
        score = 5 * score + scoreList[stack.removeLast()]!!
    }
    return score
}

fun calcPoints(removedItem: String): Int {
    val scoreList = mapOf(")" to 3, "]" to 57, "}" to 1197, ">" to 25137)
    return scoreList[removedItem]!!
}
