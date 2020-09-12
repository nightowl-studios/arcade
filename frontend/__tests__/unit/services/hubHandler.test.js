// import { EventBus } from "../../../src/eventBus.js";
// import { Event } from "../../../src/events.js";
// import Player from "../../../src/modules/common/entities/player.js";
// import HubHandler from "../../../src/services/eventHandlers/hubHandler";
// import { store } from "../../../src/store.js";

// jest.mock("../../../src/eventBus");
// jest.mock("../../../src/store.js");

// const player1 = new Player("ABC", "player1", 0);
// const player2 = new Player("DEF", "player2", 1);

// afterEach(() => {
//     jest.clearAllMocks();
// });

// describe("hubHandler", () => {
//     let hubHandler;
//     beforeEach(() => (hubHandler = new HubHandler()));
//     describe("handlePlayersChanged", () => {
//         test("if player joined", () => {
//             const playerMock = [player1, player2];
//             store.getters = {
//                 "application/getPlayers": [player1],
//             };

//             hubHandler.handlePlayersChanged(playerMock);
//             expect(EventBus.$emit).toBeCalledTimes(1);
//             expect(EventBus.$emit).toBeCalledWith(Event.PLAYER_JOIN, player2);
//         });
//         test("if player left", () => {
//             const playerMock = [];
//             store.getters = {
//                 "application/getPlayers": [player1],
//             };
//             hubHandler.handlePlayersChanged(playerMock);
//             expect(EventBus.$emit).toBeCalledTimes(1);
//             expect(EventBus.$emit).toBeCalledWith(Event.PLAYER_LEFT, player1);
//         });
//     });
// });

describe("todo", () => {
    it("todo", () => {
    });
});
