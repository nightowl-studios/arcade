import { EventBus } from "@/eventBus";
import { Event } from "@/events";
import {
    ChoosingWord,
    Drawing,
    Guessing,
    WaitingForPlayerToChooseWord,
} from "@/modules/scribble/stores/states/gamestates";
import { store } from "@/store";

const NANOSECOND_TO_SECONDS_FACTOR = 1000000000;

// Event handler for Game API
export default class GameHandler {
    constructor() {
        this.setGameStateKey = "scribble/setGameState";
    }

    _convertToSeconds(durationNS) {
        return durationNS / NANOSECOND_TO_SECONDS_FACTOR;
    }

    handle(payload) {
        if (payload.gameMasterAPI === "waitForStart") {
            store.commit(
                "application/setPlayerReadyState",
                payload.waitForStart
            );
        } else if (payload.gameMasterAPI === "wordSelect") {
            const playerUuid = store.getters["application/getPlayerUuid"];
            if (playerUuid === payload.wordSelect.chosenUUID) {
                const player = store.getters["application/getPlayerWithUuid"](
                    playerUuid
                );
                const state = new ChoosingWord(
                    player,
                    payload.wordSelect.choices,
                    this._convertToSeconds(payload.wordSelect.duration)
                );
                store.commit(this.setGameStateKey, state);
            } else {
                const player = store.getters["application/getPlayerWithUuid"](
                    playerUuid
                );
                const state = new WaitingForPlayerToChooseWord(
                    player,
                    this._convertToSeconds(payload.wordSelect.duration)
                );
                store.commit(this.setGameStateKey, state);
            }
        } else if (payload.gameMasterAPI === "playTime") {
            const currentState = store.getters["scribble/getGameState"];
            if (currentState.state === ChoosingWord.STATE) {
                const selectedWord = store.getters["scribble/getWordSelected"];
                const state = new Drawing(
                    selectedWord,
                    this._convertToSeconds(payload.playTimeSend.duration)
                );
                store.commit(this.setGameStateKey, state);
            } else if (
                currentState.state === WaitingForPlayerToChooseWord.STATE
            ) {
                const state = new Guessing(
                    payload.playTimeSend.hint,
                    this._convertToSeconds(payload.playTimeSend.duration)
                );
                store.commit(this.setGameStateKey, state);
            } else {
                const totalScores = payload.playTimeSend.totalScore;
                Object.keys(totalScores).forEach((key) => {
                    const playerScore = {
                        uuid: key,
                        score: totalScores[key],
                    };
                    store.commit("application/setPlayerScore", playerScore);
                });

                let correctClientUuid = payload.playTimeSend.correctClient.UUID;

                const player = store.getters["application/getPlayerWithUuid"](
                    correctClientUuid
                );
                EventBus.$emit(Event.CORRECT_GUESS, player);
            }
        }
    }
}
