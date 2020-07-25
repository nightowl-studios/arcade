import { EventBus } from "../../../src/eventBus.js";
import GameHandler from "../../../src/services/eventHandlers/gameHandler";
import { store } from "../../../src/store.js";

jest.mock("../../../src/eventBus");
jest.mock("../../../src/store.js", () => ({
    get store() {
        return {
            getters: {
                "application/getPlayerUuid": "123",
                "application/getPlayerWithUuid": (uuid) => {
                    return {
                        playerUuid: "123",
                    };
                },
            },
            commit: (a, b) => {},
        };
    },
}));

describe("gameHandler", () => {
    let gameHandler = new GameHandler();

    const wordSelectMock = {
        gameMasterAPI: "wordSelect",
        wordSelect: {
            chosenUUID: "123",
            choices: ["A", "B", "C"],
        },
    };

    describe("when wordSelect received", () => {
        describe("if player is chosen", () => {
            gameHandler.handle(wordSelectMock);
            it("commits", () => {
                expect(store.commit).toHaveBeenCalledWith;
                "scribble/setGameState", expect.anything();
                expect(EventBus.$emit).toBeCalledWith(Event.GAME_STARTED);
            });
        });
    });
});
