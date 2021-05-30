import { TestBed } from '@angular/core/testing';

import { CreateLobbyService } from './create-lobby.service';

describe('CreateLobbyService', () => {
  let service: CreateLobbyService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(CreateLobbyService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
