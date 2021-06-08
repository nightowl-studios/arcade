import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { catchError, map, tap } from 'rxjs/operators';
import { hub } from '../app/interface/hub';
import { webSocket } from 'rxjs/webSocket';
//import { webSocket, WebSocketSubject } from 'rxjs/webSocket';

@Injectable({
  providedIn: 'root',
})
export class CreateLobbyService {
  constructor(private http: HttpClient) {}

  private createLobbyReqUrl = 'http://localhost:8081/hub';

  // createLobby will return a lobbyID if successful
  createLobby(): string {
    var someObject = this.http.get<hub>(this.createLobbyReqUrl);
    const createLobbyObserver = {
      next: (x: hub) => {
        console.log('Observer got a next value: ' + x.hubID);
        //this.socket$ =
        var subject = webSocket('ws://localhost:4200/' + x.hubID);
      },
      error: (err: Error) => console.error('Observer got an error: ' + err),
      complete: () => console.log('Observer got a complete notification'),
    };
    someObject.subscribe(createLobbyObserver);
    return 'hello';
  }
}
