import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';

// app components
import { AppRoutingModule } from './app-routing.module';
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
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FontAwesomeModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
