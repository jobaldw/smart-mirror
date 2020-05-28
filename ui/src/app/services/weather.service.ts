import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { environment } from 'src/environments/environment';

const apiKey: string = environment.apiKey;
@Injectable({
  providedIn: 'root'
})
export class WeatherService {
  constructor(private http: HttpClient) { }

  getCurrentWeather() {
    return this.http.get(`${environment.apiUrl}/weather?q=Saint Louis,usa&appid=${apiKey}`)
  }
  getForecast() {
    return this.http.get(`${environment.apiUrl}/forecast?q=Saint Louis,usa&appid=${apiKey}`)
  }
}