class GameState {
    constructor(
        state,
        lockCanvas,
        showWordChoices,
        showPlayerChoosing,
        showWordToGuess
    ) {
        this.state = state;

        this.lockCanvas = lockCanvas;
        this.showWordChoices = showWordChoices;
        this.showPlayerChoosing = showPlayerChoosing;
        this.showWordToGuess = showWordToGuess;
    }
}

export class ChoosingWord extends GameState {
    static STATE = "ChoosingWord";
    constructor(player, words, duration) {
        super(ChoosingWord.STATE, true, true, true, false);
        this.player = player;
        this.words = words;
        this.duration = duration;
    }
}

export class WaitingForPlayerToChooseWord extends GameState {
    static STATE = "WaitingForPlayerToChooseWord";
    constructor(player) {
        super(WaitingForPlayerToChooseWord.STATE, true, false, true, false);
        this.player = player;
    }
}

export class Drawing extends GameState {
    static STATE = "Drawing";
    constructor(word) {
        super(Drawing.STATE, false, false, false, true);
        this.word = word;
    }
}

export class Guessing extends GameState {
    static STATE = "Guessing";
    constructor(word) {
        super(Guessing.STATE, true, false, false, true);
        this.word = word;
    }
}
