import { Component, OnInit } from '@angular/core';
import { Store } from '@ngrx/store';

import { CalendarService } from '../../services/calendar/calendar.service';

import { faMapMarker, IconDefinition } from '@fortawesome/free-solid-svg-icons';

@Component({
  selector: 'app-calendar',
  templateUrl: './calendar.component.html',
  styleUrls: ['./calendar.component.css']
})

export class CalendarComponent implements OnInit {
  calendar: any = <any>{};

  location: IconDefinition;

  msg: string;

  constructor(
    private store: Store<any>,
    private CalendarSerice: CalendarService
  ) {
    this.getEvents()
    this.location = faMapMarker
  }

  ngOnInit(): void {
  }

  getEvents() {
    this.CalendarSerice.getEvents()
    .subscribe(res => {
      this.calendar = res;
    }, err => {
      if (err.error && err.error.message) {
        this.msg = err.error.message;
        return;
      }
      this.msg = "Failed to get calendar events.";
    }, () => {})
  }

  hasLocation(int) {
    return (this.calendar.events[int].location != "");
  }

  isToday(date: string) {
    var today = new Date();
    var input = new Date(date);
    return (input.setHours(0,0,0,0) == today.setHours(0,0,0,0));
  }
}
