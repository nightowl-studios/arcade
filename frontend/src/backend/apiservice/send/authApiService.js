export default class AuthApiService {
    constructor(webSocketService, cookieService) {
        this.webSocketService = webSocketService;
        this.cookieService = cookieService;
    }

    authenticate() {
        // const arcadeSession = this.cookieService.getArcadeCookie();
        // if (arcadeSession != null && arcadeSession.ContainsToken != false) {
        //     this.send(arcadeSession);
        // } else {
        const noToken = {
            api: "auth",
            payload: {
                ContainsToken: false,
            },
        };

        this.send(noToken);
    }

    send(data) {
        this.webSocketService.send(data);
    }
}
