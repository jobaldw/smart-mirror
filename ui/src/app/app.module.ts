// anuglar modules
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { StoreModule } from '@ngrx/store';

// 3rd party modules
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';

// helper modules
import { AppRoutingModule } from './app-routing.module';
import { locationReducer } from './location-reducer';

// services
import { BudgetService } from './services/budget/budget.service';
import { CalendarService } from './services/calendar/calendar.service';
import { EntertainmentService } from './services/entertainment/entertainment.service';
import { NotificationsService } from './services/notifications/notifications.service';
import { WeatherService } from './services/weather/weather.service';

// app components
import { AppComponent } from './app.component';
import { NavigationBarComponent } from './navigation-bar/navigation-bar.component';

// entertainment components
import { EntertainmentComponent } from './entertainment/entertainment.component';
import { WatchListComponent } from './entertainment/watch-list/watch-list.component';

// home components
import { HomeComponent } from './home/home.component';
import { CalendarComponent } from './home/calendar/calendar.component';
import { ClockComponent } from './home/clock/clock.component';
import { CurrentComponent } from './home/weather/current/current.component';
import { ForecastComponent } from './home/weather/forecast/forecast.component';
import { NotificationsComponent } from './home/notifications/notifications.component';
import { SearchComponent } from './home/weather/search/search.component';
import { WeatherComponent } from './home/weather/weather.component';

// page not found component
import { PageNotFoundComponent } from './page-not-found/page-not-found.component';

import { GoogleHubComponent } from './home/google-hub/google-hub.component';
import { BudgetComponent } from './budget/budget.component';
import { SavingsComponent } from './budget/savings/savings.component';
import { OverviewComponent } from './budget/overview/overview.component';

@NgModule({
  declarations: [
    AppComponent,
    CalendarComponent,
    ClockComponent,
    CurrentComponent,
    EntertainmentComponent,
    ForecastComponent,
    GoogleHubComponent,
    HomeComponent,
    NavigationBarComponent,
    NotificationsComponent,
    PageNotFoundComponent,
    WatchListComponent,
    WatchListComponent,
    SearchComponent,
    WeatherComponent,
    BudgetComponent,
    SavingsComponent,
    OverviewComponent,
  ],
  imports: [
    AppRoutingModule,
    BrowserModule,
    FontAwesomeModule,
    FormsModule,
    HttpClientModule,
    StoreModule.forRoot({
      loc: locationReducer
    }),
  ],
  providers: [
    BudgetService,
    CalendarService,
    EntertainmentService,
    NotificationsService,
    WeatherService,
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
