import { Component, OnInit } from '@angular/core';
import { Store, select } from '@ngrx/store';
import { Observable } from 'rxjs';

import { faCloud, faSnowflake, faSun, faCloudRain, faTemperatureLow, faTemperatureHigh, faWater } from '@fortawesome/free-solid-svg-icons'

@Component({
  selector: 'app-weather',
  templateUrl: './weather.component.html',
  styleUrls: ['./weather.component.css']
})
export class WeatherComponent implements OnInit{

  loc$: Observable<string>;
  loc: string;

  faTemperatureLow = faTemperatureLow;
  faTemperatureHigh = faTemperatureHigh;
  faWater = faWater
  
  faSun = faSun;
  faCloud = faCloud;
  faCloudRain = faCloudRain;
  faSnowflake = faSnowflake;
  

  temperature;

  feelsLike;
  rain;
  humidity;
  high;
  low;

  forecast;
  icon;
  description;

  constructor(private store: Store<any>) {
    this.loc$ = store.pipe(select('loc'));
    this.loc$.subscribe(loc => {
      this.loc = loc;
    })

    this.temperature = 77;

    this.feelsLike = 79;
    this.rain = 40;
    this.humidity = 83;
    this.high = 81;
    this.low = 68;

    this.forecast = 'sunny';
    this.description = 'mostly sunny, with a few clouds';
   }

  // constructor(private weatherService:WeatherService) {
  //   this.currently = weatherService.getCurrentWeather();
  //   this.forecast = weatherService.getForecast()
  //  }

  ngOnInit() {
  }
}
