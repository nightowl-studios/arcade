class GameState {
    constructor(state) {
        this.state = state;
    }
}

export class ChoosingWord extends GameState {
    static STATE = 'ChoosingWord';
    constructor(player, words) {
        super(ChoosingWord.STATE);
        this.player = player;
        this.words = words;
    }

}

export class WaitingForPlayerToChooseWord extends GameState {
    static STATE = 'WaitingForPlayerToChooseWord';
    constructor(player) {
        super(WaitingForPlayerToChooseWord.STATE);
        this.player = player;
    }

}
