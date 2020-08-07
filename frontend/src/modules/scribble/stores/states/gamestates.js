class GameStateFlags {
    constructor() {
        this.lockCanvas = false;
        this.showWordChoices = false;
        this.showPlayerChoosing = false;
        this.showWordToGuess = false;
        this.showLobby = false;
    }

    forState(state) {
        this.state = state;
        return this;
    }

    lockedCanvas() {
        this.lockCanvas = true;
        return this;
    }

    showingWordChoices() {
        this.showWordChoices = true;
        return this;
    }

    showingPlayerChoosing() {
        this.showPlayerChoosing = true;
        return this;
    }

    showingWordToGuess() {
        this.showWordToGuess = true;
        return this;
    }

    showingLobby() {
        this.showLobby = true;
        return this;
    }
}

class GameState {
    constructor(gameStateFlags) {
        if (gameStateFlags.state == null) {
            throw new Error("GameState requires a state parameter");
        }
        this.state = gameStateFlags.state;

        this.lockCanvas = gameStateFlags.lockCanvas;
        this.showWordChoices = gameStateFlags.showWordChoices;
        this.showPlayerChoosing = gameStateFlags.showPlayerChoosing;
        this.showWordToGuess = gameStateFlags.showWordToGuess;
        this.showLobby = gameStateFlags.showLobby;
    }
}

export class WaitingInLobby extends GameState {
    static STATE = "WaitingInLobby";
    constructor() {
        super(
            new GameStateFlags()
                .forState(WaitingInLobby.STATE)
                .lockedCanvas()
                .showingLobby()
        );
    }
}

export class ChoosingWord extends GameState {
    static STATE = "ChoosingWord";
    constructor(player, words, durationSec) {
        super(
            new GameStateFlags()
                .forState(ChoosingWord.STATE)
                .lockedCanvas()
                .showingWordChoices()
                .showingPlayerChoosing()
        );
        this.player = player;
        this.words = words;
        this.durationSec = durationSec;
    }
}

export class WaitingForPlayerToChooseWord extends GameState {
    static STATE = "WaitingForPlayerToChooseWord";
    constructor(player, durationSec) {
        super(
            new GameStateFlags()
                .forState(WaitingForPlayerToChooseWord.STATE)
                .lockedCanvas()
                .showingPlayerChoosing()
        );
        this.player = player;
        this.durationSec = durationSec;
    }
}

export class Drawing extends GameState {
    static STATE = "Drawing";
    constructor(word, durationSec) {
        super(
            new GameStateFlags().forState(Drawing.STATE).showingWordToGuess()
        );
        this.word = word;
        this.durationSec = durationSec;
    }
}

export class Guessing extends GameState {
    static STATE = "Guessing";
    constructor(word, durationSec) {
        super(
            new GameStateFlags()
                .forState(Guessing.STATE)
                .lockedCanvas()
                .showingWordToGuess()
        );
        this.word = word;
        this.durationSec = durationSec;
    }
}
