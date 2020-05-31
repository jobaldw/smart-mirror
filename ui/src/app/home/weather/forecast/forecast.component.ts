import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { Store, select } from '@ngrx/store';

import { WeatherService } from '../../../services/weather.service';

import { faSun, faCloudSun, faCloud,  faCloudRain, faCloudSunRain, faCloudShowersHeavy, faSnowflake, faQuestion, IconDefinition } from '@fortawesome/free-solid-svg-icons';

@Component({
  selector: 'app-forecast',
  templateUrl: './forecast.component.html',
  styleUrls: ['./forecast.component.css']
})

export class ForecastComponent implements OnInit {
  loc$: Observable<string>;
  loc: string;
  forecast: any = <any>{};
  currentWeather: any = <any>{};
  
  day0: Date;
  day1: Date;
  day2: Date;
  day3: Date;
  
  dayicon0: IconDefinition;
  dayicon1: IconDefinition;
  dayicon2: IconDefinition;
  dayicon3: IconDefinition;

  msg: string;
  
  constructor(
    private store: Store<any>,
    private weatherService: WeatherService
  ) {
    this.loc$ = store.pipe(select('loc'));
    this.loc$.subscribe(loc => {
      this.loc = loc;

      this.getLatitiudeLongitude(loc);
    })
  }
  ngOnInit(): void {
  }

  searchForecast(lat: string, lon: string) {
    this.weatherService.getForecast(lat, lon)
      .subscribe(res => {
        this.forecast = res;
      }, err => {
        if (err.error && err.error.message) {
          this.msg = err.error.message;
          return;
        }
        this.msg = "Failed to get 4 day forecast.";
      }, () => {

        var weekday0 = new Date();
        var weekday1 = new Date();
        var weekday2 = new Date();
        var weekday3 = new Date();

        var dayIcon: string;
        var icon: IconDefinition;

        for (let i = 0; i < 4; i++) {
          dayIcon = this.forecast.daily[i].weather[0].icon;
          if (dayIcon == "01d") {
            icon = faSun;
          } else if (dayIcon == "02d") {
            icon = faCloudSun;
          } else if (dayIcon == "03d" || dayIcon == "04d") {
            icon = faCloud;
          } else if (dayIcon == "09d") {
            icon = faCloudRain;
          } else if (dayIcon == "10d") {
            icon = faCloudSunRain;
          } else if (dayIcon == "11d") {
            icon = faCloudShowersHeavy;
          } else if (dayIcon == "13d") {
            icon = faSnowflake;
          } else {
            icon = faQuestion;
          }

          if (i == 0) {
            weekday0.setDate(weekday0.getDate() + 1);
            this.day0 = weekday0;
            this.dayicon0 = icon;
          }
          if (i == 1) {
            weekday1.setDate(weekday1.getDate() + 2);
            this.day1 = weekday1;
            this.dayicon1 = icon;
          }
          if (i == 2) {
            weekday2.setDate(weekday2.getDate() + 3);
            this.day2 = weekday2;
            this.dayicon2 = icon;
          }
          if (i == 3) {
            weekday3.setDate(weekday3.getDate() + 4);
            this.day3 = weekday3;
            this.dayicon3 = icon;
          }
        }
      })
  }

  getLatitiudeLongitude(loc: string) {
    this.weatherService.getCurrentWeather(loc)
    .subscribe(res => {
      this.currentWeather = res;
    }, err => {
      if (err.error && err.error.message) {
        this.msg = err.error.message;
        return;
      }
      this.msg = "Failed to get current weather.";
    }, () => {
      this.searchForecast(this.currentWeather.coord.lat, this.currentWeather.coord.lon);
    })
  }



  resultFound() {
    return Object.keys(this.currentWeather).length > 0;
  }
}
