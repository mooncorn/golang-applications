
class Character {
    constructor(name, hp, ap) {
        this.name = name
        this.hp = hp
        this.ap = ap
    }

    attack(target) {
        target.hp -= this.ap
    }

    takeDamage(ap) {
        this.hp -= ap
    }
}

class Player extends Character {
    heal(hp) {
        this.hp += hp
    }
}

class Monster extends Character {
    constructor(name) {
        const hp = getRandomInt(100) + 1
        const ap = getRandomInt(100) + 1
        super(name, hp, ap)
    }
}

class Boss extends Monster {
    constructor(name) {
        const hp = getRandomInt(200) + 1
        const ap = getRandomInt(200) + 1
        super(name, hp, ap)
    }
}

const getRandomInt = (max) => {
    return Math.floor(Math.random() * max);
}