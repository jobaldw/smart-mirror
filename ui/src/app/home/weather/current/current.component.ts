import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { Store, select } from '@ngrx/store';
import { WeatherService } from '../../../services/weather/weather.service';

import { faSun, faCloudSun, faCloud,  faCloudRain, faCloudSunRain, faCloudShowersHeavy, faSnowflake, faTemperatureLow, faTemperatureHigh, faWater, faQuestion, IconDefinition } from '@fortawesome/free-solid-svg-icons';

@Component({
  selector: 'app-current',
  templateUrl: './current.component.html',
  styleUrls: ['./current.component.css']
})
export class CurrentComponent implements OnInit {

  faTemperatureLow = faTemperatureLow;
  faTemperatureHigh = faTemperatureHigh;
  faWater = faWater
  faCloudRain = faCloudRain;

  loc$: Observable<string>;
  loc: string;
  currentWeather: any = <any>{};
  rain: any = <any>{};
  msg: string;
  icon: IconDefinition;
  dayIcon: string;

  constructor(
    private store: Store<any>,
    private weatherService: WeatherService
  ) {
    this.loc$ = store.pipe(select('loc'));
    this.loc$.subscribe(loc => {
      this.loc = loc;
      this.searchWeather(loc);
    })
  }

  ngOnInit(): void {
  }

  searchWeather(loc: string) {
    this.msg = '';
    this.currentWeather = {};
    this.weatherService.getCurrentWeather(loc)
      .subscribe(res => {
        this.currentWeather = res;
      }, err => {
        if (err.error && err.error.message) {
          this.msg = err.error.message;
          return;
        }
      }, () => {
        if (this.currentWeather.rain) {
          this.rain = (this.currentWeather.rain["1h"] * 100);
        } else {
          this.rain = 0;
        }

        this.dayIcon = this.currentWeather.weather[0].icon;
          if (this.dayIcon == "01n") {
            this.icon = faSun;
          } else if (this.dayIcon == "02n") {
            this.icon = faCloudSun;
          } else if (this.dayIcon == "03n" || this.dayIcon == "04n") {
            this.icon = faCloud;
          } else if (this.dayIcon == "09n") {
            this.icon = faCloudRain;
          } else if (this.dayIcon == "10n") {
            this.icon = faCloudSunRain;
          } else if (this.dayIcon == "11n") {
            this.icon = faCloudShowersHeavy;
          } else if (this.dayIcon == "13n") {
            this.icon = faSnowflake;
          } else {
            this.icon = faQuestion;
          }
      })
  }

  resultFound() {
    return Object.keys(this.currentWeather).length > 0;
  }
}
