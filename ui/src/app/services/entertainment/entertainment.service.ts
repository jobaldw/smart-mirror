import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class EntertainmentService {
  constructor(private http: HttpClient) { }

  getTitles(key: string, value: string) {
    return this.http.get(`${environment.entertainmentAPIURL}/v1/watchlist?${key}=${value}`)
  }
}
