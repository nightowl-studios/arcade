export default class Player {
    constructor(uuid, nickname, joinOrder) {
        this.uuid = uuid;
        this.nickname = nickname;
        this.joinOrder = joinOrder;
        this.isReady = false;
    }
}
