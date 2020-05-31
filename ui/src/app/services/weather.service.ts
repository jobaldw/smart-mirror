import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment';

const apiKey: string = environment.apiKey;
@Injectable({
  providedIn: 'root'
})
export class WeatherService {
  constructor(private http: HttpClient) { }

  getCurrentWeather(loc: string) {
    return this.http.get(`${environment.apiUrl}/weather?units=imperial&q=${loc}&appid=${apiKey}`)
  }

  getForecast(lat: string, lon: string) {
    return this.http.get(`${environment.apiUrl}/onecall?lat=${lat}&lon=${lon}&units=imperial&appid=${apiKey}`)
  }
}