
export default class ApiReceiverService {
    constructor() {
        this.listeners = [];
    }

    update(event, data) {
        this.listeners.forEach(listener => listener.update(event, data));
    }

    addListener(listener) {
        this.listeners.push(listener);
    }
}
