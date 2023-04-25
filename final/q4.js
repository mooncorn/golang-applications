
const vowels = ["a", "e", "i", "o", "u"]

const vowelCount = (input) => {
    let count = 0
    for (let i = 0; i < input.length; i++) {
        if (vowels.includes(input[i].toLowerCase())) {
            count++
        }
    }
    return count
}

console.log(vowelCount("aekkkkffffdsdfiou"))