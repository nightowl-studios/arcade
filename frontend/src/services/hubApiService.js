import axios from 'axios'

export default class HubApiService {
    constructor(httpUrl) {
        this.apiUrl = httpUrl + '/hub'
    }

    async createLobby() {
        console.log('Creating room...')
        let url = this.apiUrl
        let response = await axios.get(url)
        return response.data.hubID
    }

    async checkLobbyExists(lobbyId) {
        let url = this.apiUrl + '/' + lobbyId
        let response = await axios.get(url)
        if (response.data.exists) {
            return true
        } else {
            return false
        }
    }
}
