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
  

  constructor(private store: Store<any>) {
    this.loc$ = store.pipe(select('loc'));
    this.loc$.subscribe(loc => {
      this.loc = loc;
    })
   }

  ngOnInit() {
  }
}
