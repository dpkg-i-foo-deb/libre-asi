import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/internal/Observable';
import { environment } from 'src/environment/environment';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { ApiResponse } from '../models/api-response';

@Injectable({
  providedIn: 'root',
})
export class PingService {
  pingOptions = {
    headers: new HttpHeaders({
      'Content-Type': 'application/json',
    }),
    withCredentials: true,
  };

  constructor(protected http: HttpClient) {}

  ping(): Observable<ApiResponse> {
    return this.http.get<ApiResponse>(
      environment.apiUrl + environment.pingPath,
      this.pingOptions
    );
  }
}
