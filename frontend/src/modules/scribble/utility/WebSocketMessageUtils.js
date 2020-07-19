export function createChatMessage(message) {
  return createWebSocketMessage("chat", { message: message });
}

export function createDrawMessage(payload) {
  return createWebSocketMessage("draw", payload);
}

export function createWebSocketMessage(api, payload) {
  return {
    api: api,
    payload: payload,
  };
}
