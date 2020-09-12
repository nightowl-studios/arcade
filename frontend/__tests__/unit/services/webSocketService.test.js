// import { EventBus } from "../../../src/eventBus.js";
// import { Event } from "../../../src/events.js";
// import WebSocketService from "../../../src/services/webSocketService";

// jest.mock("../../../src/eventBus");

// describe("webSocketService", () => {
//     let cookieService = {
//         getArcadeCookie: jest.fn(),
//         setArcadeCookie: jest.fn(),
//     };
//     let eventHandlerService = {
//         handle: jest.fn(),
//     };

//     const originalError = console.error;

//     let service = new WebSocketService(
//         "abc",
//         cookieService,
//         eventHandlerService
//     );

//     let websocket = {
//         send: jest.fn(),
//         readyState: WebSocket.CLOSED,
//         close: jest.fn(),
//     };

//     afterEach(() => {
//         console.error = originalError;
//         websocket.readyState = WebSocket.CLOSED;
//         jest.clearAllMocks();
//     });

//     describe("after initing websocket", () => {
//         service.webSocket = websocket;

//         describe("if the session contains a token on open", () => {
//             cookieService.getArcadeCookie.mockReturnValueOnce({
//                 ContainsToken: true,
//             });
//             it("sends json data if connected", () => {
//                 websocket.readyState = WebSocket.OPEN;

//                 service.initWebSocket(websocket, "123");
//                 websocket.onopen();

//                 expect(EventBus.$emit).toBeCalledWith(
//                     Event.WEBSOCKET_CONNECTED,
//                     "123"
//                 );
//                 expect(websocket.send).toBeCalledWith('{"ContainsToken":true}');
//             });
//             it("logs an error if not connected", () => {
//                 let consoleOutput = [];
//                 console.error = (output) => consoleOutput.push(output);

//                 service.initWebSocket(websocket, "123");
//                 websocket.onopen();

//                 expect(EventBus.$emit).toBeCalledWith(
//                     Event.WEBSOCKET_CONNECTED,
//                     "123"
//                 );
//                 expect(consoleOutput).toEqual(["NOT CONNECTED"]);
//             });
//         });
//         describe("if the session does not contain a token on open", () => {
//             cookieService.getArcadeCookie.mockReturnValueOnce({
//                 ContainsToken: false,
//             });
//             it("sends no token json data if connected", () => {
//                 websocket.readyState = WebSocket.OPEN;

//                 service.initWebSocket(websocket, "123");
//                 websocket.onopen();

//                 expect(EventBus.$emit).toBeCalledWith(
//                     Event.WEBSOCKET_CONNECTED,
//                     "123"
//                 );
//                 expect(websocket.send).toBeCalledWith(
//                     '{"api":"auth","payload":{"ContainsToken":false}}'
//                 );
//             });
//             it("logs an error if not connected", () => {
//                 let consoleOutput = [];
//                 console.error = (output) => consoleOutput.push(output);

//                 service.initWebSocket(websocket, "123");
//                 websocket.onopen();

//                 expect(EventBus.$emit).toBeCalledWith(
//                     Event.WEBSOCKET_CONNECTED,
//                     "123"
//                 );
//                 expect(consoleOutput).toEqual(["NOT CONNECTED"]);
//             });
//         });
//     });
//     describe("on disconnect", () => {
//         it("disconnects if connected", () => {
//             websocket.readyState = WebSocket.OPEN;

//             service.disconnect();

//             expect(websocket.close).toBeCalledTimes(1);
//         });
//         it("does nothing if not connected", () => {
//             service.disconnect();
//             expect(websocket.close).toBeCalledTimes(0);
//         });
//     });
// });

describe("todo", () => {
    it("does nothing if not connected", () => {
    });
});
