import ChatHandler from "./chatHandler";
import GameHandler from "./gameHandler";
import HubHandler from "./hubHandler";

export default class EventHandlerFactory {
    constructor() {
        this.hubHandler = new HubHandler();
        this.gameHandler = new GameHandler();
        this.chatHandler = new ChatHandler();
    }

    getHandler(api) {
        if (api === "hub") {
            return this.hubHandler;
        } else if (api === "game") {
            return this.gameHandler;
        } else if (api === "chat") {
            return this.chatHandler;
        }
        return null;
    }
}
