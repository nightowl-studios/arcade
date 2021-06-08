import { Injectable } from '@angular/core';
import { webSocket, WebSocketSubject } from 'rxjs/webSocket';

@Injectable({
  providedIn: 'root',
})
export class WebsocketServiceService {
  private socket$: WebSocketSubject<any>;

  constructor(private hubID: string) {
    this.socket$ = this.getNewWebSocket();
  }
  //public connect(): void {

  //if (!this.socket$ || this.socket$.closed) {
  //this.socket$ = this.getNewWebSocket();
  //const messages = this.socket$.pipe(
  //tap({
  //error: error => console.log(error),
  //}), catchError(_ => EMPTY));
  //this.messagesSubject$.next(messages);
  //}
  //}

  private getNewWebSocket() {
    return webSocket('http://localhost:8081/hub/');
  }
}
