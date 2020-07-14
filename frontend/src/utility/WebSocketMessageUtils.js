export function createChatMessage(message) {
  return createWebSocketMessage("chat", { message: message });
}

export function createBrushStrokeMessage(brushstroke) {
  return createWebSocketMessage("brushstroke", { brushstroke: brushstroke });
}

export function createWebSocketMessage(api, payload) {
  return {
    api: api,
    payload: payload,
  };
}
