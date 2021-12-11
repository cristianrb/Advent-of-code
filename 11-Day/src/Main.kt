import java.io.File

class Octopus(val point: Pair<Int, Int>, var energy: Int) {

    var flashedTimes = 0
    var flashedStep = 0

    fun addEnergy(actualStep: Int) {
        if (actualStep != flashedStep)
            energy++
    }

    fun flash(actualStep: Int) {
        flashedTimes++
        flashedStep = actualStep
        energy = 0
    }

    fun ready() = energy > 9

}

fun main() {
    val input = File("src/input")
        .readLines()
        .flatMapIndexed { i, line ->
            line.mapIndexed { j, energy ->
                Pair(i, j) to Octopus(Pair(i, j), energy.toString().toInt())
            }
        }.toMap()

    solve(input)
}

fun solve(input: Map<Pair<Int, Int>, Octopus>) {
    repeat(101) { step ->
        doOneStep(input, step)
    }

    println("Part 1: " + input.values.sumOf { it.flashedTimes })

    var firstStepAllFlash = 0
    var step = 101
    while (firstStepAllFlash == 0) {
        doOneStep(input, step)
        if (input.values.all { it.energy == 0 }) {
            firstStepAllFlash = step
        }
        step++
    }
    println("Part 2: $firstStepAllFlash")
}

private fun doOneStep(input: Map<Pair<Int, Int>, Octopus>, step: Int) {
    input.values.map {
        it.addEnergy(step)
    }

    var octopusToFlash = input.values.filter { it.ready() }
    while (octopusToFlash.isNotEmpty()) {
        octopusToFlash.map {
            it.flash(step)
            it.point.adjacentPoints().map {
                input[it]?.addEnergy(step)
            }
        }
        octopusToFlash = input.values.filter { it.ready() }
    }
}

fun Pair<Int, Int>.adjacentPoints(): List<Pair<Int, Int>> {
    return listOf(
        Pair(this.first + 1, this.second),
        Pair(this.first - 1, this.second),
        Pair(this.first, this.second + 1),
        Pair(this.first, this.second - 1),
        Pair(this.first + 1, this.second + 1),
        Pair(this.first + 1, this.second - 1),
        Pair(this.first - 1, this.second + 1),
        Pair(this.first - 1, this.second - 1),
    )
}
