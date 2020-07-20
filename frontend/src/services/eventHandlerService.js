import { GlobalStore } from "@/modules/common/store/globalstore/index";
import { EventBus } from "../eventBus.js";

export default class EventHandlerService {
  constructor() {}

  handle(api, payload) {
    if (api === "hub") {
      this.handleHubEvent(payload);
    }

    EventBus.$emit(api, payload);
  }

  handleHubEvent(payload) {
    if (payload.connectedClients != null) {
      GlobalStore.commit("setPlayers", payload.connectedClients);
    }
  }
}
