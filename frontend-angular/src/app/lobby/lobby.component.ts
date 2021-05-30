import { Component, OnInit } from '@angular/core';
import { CreateLobbyService } from '../create-lobby.service';

@Component({
  selector: 'app-lobby',
  templateUrl: './lobby.component.html',
  styleUrls: ['./lobby.component.scss'],
})
export class LobbyComponent implements OnInit {
  constructor(private createLobbyService: CreateLobbyService) {}

  lobbyID = '';

  ngOnInit(): void {
    this.lobbyID = this.createLobbyService.createLobby();
    console.log('the lobby id is: ', this.lobbyID);
  }
}
