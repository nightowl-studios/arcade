import axios from "axios";

// Service for making REST API requests to HubAPI
export default class HubApiService {
    constructor(httpUrl) {
        this.apiUrl = `${httpUrl}/hub`;
    }

    async createLobby() {
        const url = this.apiUrl;
        const response = await axios.get(url);
        return response.data.hubID;
    }

    async checkLobbyExists(lobbyId) {
        const url = `${this.apiUrl}/${lobbyId}`;
        const response = await axios.get(url);
        return response.data.exists;
    }
}
