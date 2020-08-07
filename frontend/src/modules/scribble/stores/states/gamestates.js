const NANOSECOND_TO_SECONDS_FACTOR = 1000000000;

class GameState {
    constructor(
        state,
        lockCanvas,
        showWordChoices,
        showPlayerChoosing,
        showWordToGuess,
        showLobby,
        showResults
    ) {
        this.state = state;

        this.lockCanvas = lockCanvas;
        this.showWordChoices = showWordChoices;
        this.showPlayerChoosing = showPlayerChoosing;
        this.showWordToGuess = showWordToGuess;
        this.showLobby = showLobby;
        this.showResults = showResults;
    }
}

export class WaitingInLobby extends GameState {
    static STATE = "WaitingInLobby";
    constructor() {
        super(WaitingInLobby.STATE, true, false, false, false, true, false);
    }
}

export class ChoosingWord extends GameState {
    static STATE = "ChoosingWord";
    constructor(player, words, duration) {
        super(ChoosingWord.STATE, true, true, true, false, false, false);
        this.player = player;
        this.words = words;
        this.duration = duration / NANOSECOND_TO_SECONDS_FACTOR;
    }
}

export class WaitingForPlayerToChooseWord extends GameState {
    static STATE = "WaitingForPlayerToChooseWord";
    constructor(player, duration) {
        super(
            WaitingForPlayerToChooseWord.STATE,
            true,
            false,
            true,
            false,
            false,
            false
        );
        this.player = player;
        this.duration = duration / NANOSECOND_TO_SECONDS_FACTOR;
    }
}

export class Drawing extends GameState {
    static STATE = "Drawing";
    constructor(word, duration) {
        super(Drawing.STATE, false, false, false, true, false, false);
        this.word = word;
        this.duration = duration / NANOSECOND_TO_SECONDS_FACTOR;
    }
}

export class Guessing extends GameState {
    static STATE = "Guessing";
    constructor(word, duration) {
        super(Guessing.STATE, true, false, false, true, false, false);
        this.word = word;
        this.duration = duration / NANOSECOND_TO_SECONDS_FACTOR;
    }
}

export class GameOver extends GameState {
    static STATE = "GameOver";
    constructor() {
        super(GameOver.STATE, true, false, false, false, false, true);
    }
}
