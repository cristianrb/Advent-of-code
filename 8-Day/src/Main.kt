import java.io.File

fun main() {
    val input = File("src/input")
        .readLines()

    solve(input)
}

fun solve(input: List<String>) {
    var sum1 = 0
    var sum2 = 0
    val part1Nums = listOf(1, 4, 7, 8)

    input.forEach {
        val ins = it.split("|")[0].trim().split(" ")
        val outs = it.split("|")[1].trim().split(" ")

        val one = ins.filter { it.length == 2 }[0]
        val four = ins.filter { it.length == 4 }[0]
        val seven = ins.filter { it.length == 3 }[0]
        val eight = ins.filter { it.length == 7 }[0]
        val nine = ins.filter { it.length == 6 && it.containsAll(four, false) }[0]
        val zero = ins.filter { it.length == 6 && it.containsAll(seven, false) && it != nine }[0]
        val six = ins.filter { it.length == 6 && it != nine && it != zero }[0]
        val three = ins.filter { it.length == 5 && it.containsAll(seven, false) }[0]
        val five = ins.filter { it.length == 5 && it.addNonExistingChars(one).containsAll(nine, false) }[0]
        val two = ins.filter { it.length == 5 && it != three && it != five}[0]

        val segments = listOf(zero, one, two, three, four, five, six, seven, eight, nine)

        var decodedNum = ""
        outs.forEach {
            val index = segments.indexOfUnordered(it)
            decodedNum += index
            if (index in part1Nums) {
                sum1++
            }
        }
        sum2 += decodedNum.toInt()
    }
    println(sum1)
    println(sum2)
}

private fun List<String>.indexOfUnordered(str: String) : Int {
    this.forEachIndexed { index, s ->
        if (s.containsAll(str, true)) {
            return index
        }
    }
    return -1
}

private fun String.containsAll(str: String, matchSize: Boolean): Boolean {
    if (this.length != str.length && matchSize) return false
    var containsAll = true
    str.forEach {
        val containingChar = this.contains(it.toString())
        if (!containingChar) containsAll = false
    }
    return containsAll
}

private fun String.addNonExistingChars(str: String): String {
    var newStr = this
    str.forEach {
        if (!this.contains(it.toString())) newStr += it.toString()
    }
    return newStr
}