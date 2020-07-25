class GameState {
    constructor(state) {
        this.state = state;
    }
}

export class ChoosingWord extends GameState {
    static STATE = "ChoosingWord";
    constructor(player, words) {
        super(ChoosingWord.STATE);
        this.player = player;
        this.words = words;
    }
}

export class WaitingForPlayerToChooseWord extends GameState {
    static STATE = "WaitingForPlayerToChooseWord";
    constructor(player) {
        super(WaitingForPlayerToChooseWord.STATE);
        this.player = player;
    }
}

export class Drawing extends GameState {
    static STATE = "Drawing";
    constructor() {
        super(Drawing.STATE);
    }
}

export class Guessing extends GameState {
    static STATE = "Guessing";
    constructor() {
        super(Guessing.STATE);
    }
}
