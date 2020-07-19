export function createChatMessage(message) {
  return createWebSocketMessage("chat", { message: message });
}

export function createDrawActionMessage(drawAction) {
  return createWebSocketMessage("draw", {
    action: drawAction,
    requestHistory: false,
  });
}

export function createChangeNicknameMessage(nickname) {
  return createWebSocketMessage("hub", { changeNameTo: nickname });
}

export function createWebSocketMessage(api, payload) {
  return {
    api: api,
    payload: payload,
  };
}
