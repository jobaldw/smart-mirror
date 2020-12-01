import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class CalendarService {
  constructor(private http: HttpClient) { }

  getEvents() {
    return this.http.get(`${environment.calendarAPIURL}/events`)
  }
}
