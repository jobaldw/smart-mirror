import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { RecentComponent } from './recent/recent.component';
import { StatsComponent } from './stats/stats.component';

import { MovieService } from './movie.service';
import { MovieComponent } from './movie/movie.component';


@NgModule({
  declarations: [
    AppComponent,
    RecentComponent,
    StatsComponent,
    MovieComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule
  ],
  providers: [MovieService],
  bootstrap: [AppComponent]
})
export class AppModule { }
