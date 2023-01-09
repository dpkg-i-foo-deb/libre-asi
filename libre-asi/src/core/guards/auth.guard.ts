import { Injectable } from '@angular/core';
import { CanActivate, Router } from '@angular/router';
import { Observable, of, pipe } from 'rxjs';
import { PingService } from 'src/shared/api/ping.service';
import { ApiResponse } from 'src/shared/models/api-response';

@Injectable({
  providedIn: 'root',
})
export class AuthGuard implements CanActivate {
  constructor(private router: Router, private pingService: PingService) {}

  canActivate(): Observable<boolean> {
    const response: Observable<ApiResponse> = this.pingService.ping();

    response.subscribe({
      error: (err) => {
        console.error(err);
        return of(false);
      },
      complete: () => {
        return of(false);
      },
    });

    return of(false);
  }
}
