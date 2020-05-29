import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { Store, select } from '@ngrx/store';

import { WeatherService } from '../../../services/weather.service';

import { faCloud, faSnowflake, faSun, faCloudRain } from '@fortawesome/free-solid-svg-icons';
@Component({
  selector: 'app-forecast',
  templateUrl: './forecast.component.html',
  styleUrls: ['./forecast.component.css']
})
export class ForecastComponent implements OnInit {

  faSun = faSun;
  faCloud = faCloud;
  faCloudRain = faCloudRain;
  faSnowflake = faSnowflake;

  loc$: Observable<string>;
  loc: string;
  forecast: any = <any>{};
  msg: string;

  day0: string;
  day1: string;
  day2: string;
  day3: string;

  constructor(
    private store: Store<any>,
    private weatherService: WeatherService
  ) {
    this.loc$ = store.pipe(select('loc'));
    this.loc$.subscribe(loc => {
      this.loc = loc;
      this.searchForecast(loc);
    })
  }
  ngOnInit(): void {
  }

  searchForecast(loc: string) {
    this.weatherService.getForecast(loc)
      .subscribe(res => {
        this.forecast = res;
      }, err => {
        if (err.error && err.error.message) {
          this.msg = err.error.message;
          return;
        }
        alert('Failed to get 4 day forecast.');
      }, () => {

      })
  }
}
