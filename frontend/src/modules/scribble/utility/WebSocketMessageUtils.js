export function createChatMessage(message) {
  return createWebSocketMessage("chat", { message: message });
}

export function createWebSocketMessage(api, payload) {
  return {
    api: api,
    payload: payload,
  };
}
