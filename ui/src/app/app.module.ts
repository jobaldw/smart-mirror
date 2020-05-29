// anuglar modules
import { HttpClientModule } from '@angular/common/http';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { StoreModule } from '@ngrx/store';
import { NgModule } from '@angular/core';

// 3rd party modules
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';

// helper modules
import { AppRoutingModule } from './app-routing.module';
import { locationReducer } from './location-reducer';

// services
import { WeatherService } from './services/weather.service';

// app components
import { AppComponent } from './app.component';
import { NavigationBarComponent } from './navigation-bar/navigation-bar.component';

// entertainment components
import { RecentlyWatchedComponent } from './entainment/recently-watched/recently-watched.component';
import { RecentlyAddedComponent } from './entainment/recently-added/recently-added.component';

// home components
import { WelcomeComponent } from './home/welcome/welcome.component';
import { ClockComponent } from './home/clock/clock.component';
import { WeatherComponent } from './home/weather/weather.component';

import { CalendarComponent } from './home/calendar/calendar.component';
import { GoogleHubComponent } from './home/google-hub/google-hub.component';
import { NotificationsComponent } from './home/notifications/notifications.component';
import { CurrentComponent } from './home/weather/current/current.component';
import { SearchComponent } from './home/weather/current/search/search.component';

@NgModule({
  declarations: [
    AppComponent,
    RecentlyWatchedComponent,
    RecentlyAddedComponent,
    NotificationsComponent,
    CalendarComponent,
    ClockComponent,
    WeatherComponent,
    GoogleHubComponent,
    NavigationBarComponent,
    WelcomeComponent,
    CurrentComponent,
    SearchComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FontAwesomeModule,
    HttpClientModule,
    FormsModule,
    StoreModule.forRoot({
      loc: locationReducer
    }),
  ],
  providers: [
    WeatherService,
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
