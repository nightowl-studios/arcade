import GameHandler from "./gameHandler";
import HubHandler from "./hubHandler";

export default class EventHandlerFactory {
    constructor() {}

    getHandler(api) {
        if (api === "hub") {
            return new HubHandler();
        } else if (api === "game") {
            return new GameHandler();
        }
        return null;
    }
}
