const readline = require("readline")

const letters = "abcdefghijklmnopqrstuvwxyz1234567890!@#$%^&*()_+~`{}|:\"<>?-=[]\\;',./"

const read = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

const promtLength = (cb) => {
    read.question("Enter the length of password: ", (ans) => {
        const num = Number(ans)

        if (!isNaN(num)) {
            cb(num)
            read.close()
        } else {
            console.log("Invalid input. Try again.")
            promtLength(cb)
        }
    })
}

const getRandomInt = (max) => {
    return Math.floor(Math.random() * max);
}

const getRandomPassword = (length) => {
    let password = ""

    for (let i = 0; i < length; i++) {
        const pos = getRandomInt(letters.length)
        password += letters[pos]
    }

    return password
} 

const main = () => {
    promtLength((num) => {
        const password = getRandomPassword(num)
        console.log(`Generated password:\n${password}`)
    })
}

main()
