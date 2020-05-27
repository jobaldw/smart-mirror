import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

// app components
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NavigationBarComponent } from './navigation-bar/navigation-bar.component';

// entertainment components
import { RecentlyWatchedComponent } from './entainment/recently-watched/recently-watched.component';
import { RecentlyAddedComponent } from './entainment/recently-added/recently-added.component';

// home components
import { WelomeComponent } from './home/welome/welome.component';
import { NotificationsComponent } from './home/notifications/notifications.component';
import { CalendarComponent } from './home/calendar/calendar.component';
import { ClockComponent } from './home/clock/clock.component';
import { WeatherComponent } from './home/weather/weather.component';
import { GoogleHubComponent } from './home/google-hub/google-hub.component';

@NgModule({
  declarations: [
    AppComponent,
    RecentlyWatchedComponent,
    RecentlyAddedComponent,
    WelomeComponent,
    NotificationsComponent,
    CalendarComponent,
    ClockComponent,
    WeatherComponent,
    GoogleHubComponent,
    NavigationBarComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
