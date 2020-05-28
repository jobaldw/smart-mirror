import { Component} from '@angular/core';
import { faCloud, faSnowflake, faSun, faCloudRain, faTemperatureLow, faTemperatureHigh, faWater } from '@fortawesome/free-solid-svg-icons'

import { WeatherService } from '../../services/weather.service'


@Component({
  selector: 'app-weather',
  templateUrl: './weather.component.html',
  styleUrls: ['./weather.component.css']
})
export class WeatherComponent {

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

  constructor() {
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
}
