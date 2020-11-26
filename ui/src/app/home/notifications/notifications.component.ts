import { Component, OnInit } from '@angular/core';
import { Store } from '@ngrx/store';

import { NotificationsService } from '../../services/notifications/notifications.service';

@Component({
  selector: 'app-notifications',
  templateUrl: './notifications.component.html',
  styleUrls: ['./notifications.component.css']
})

export class NotificationsComponent implements OnInit {
  sources: any = <any>{};
  news: any = <any>{};
  sports: any = <any>{};
  tech: any = <any>{};

  newsOutlet: string = "AP";
  sportsOutlet: string = "ESPN";
  techOutlet: string = "TechCrunch";

  msg: string;

  constructor(
    private store: Store<any>,
    private NotificationsService: NotificationsService
  ) { 
    this.geNewsArticles()
    this.geSportsArticles()
    this.getTechArticles()
  }

  ngOnInit(): void {
  }

  getSources() {
    this.NotificationsService.getSources()
    .subscribe(res => {
      this.sources = res;
    }, err => {
      if (err.error && err.error.message) {
        this.msg = err.error.message;
        return;
      }
      this.msg = "Failed to get news sources.";
    }, () => {
      console.log(this.sources)
    })
  }

  geNewsArticles() {
    this.NotificationsService.getArticles("associated-press")
    .subscribe(res => {
      this.news = res;
    }, err => {
      if (err.error && err.error.message) {
        this.msg = err.error.message;
        return;
      }
      this.msg = "Failed to get news articles.";
    }, () => {
      console.log(this.sources)
    })
  }

  geSportsArticles() {
    this.NotificationsService.getArticles("espn")
    .subscribe(res => {
      this.sports = res;
    }, err => {
      if (err.error && err.error.message) {
        this.msg = err.error.message;
        return;
      }
      this.msg = "Failed to get news articles.";
    }, () => {
      console.log(this.sources)
    })
  }

  getTechArticles() {
    this.NotificationsService.getArticles("techCrunch")
    .subscribe(res => {
      this.tech = res;
    }, err => {
      if (err.error && err.error.message) {
        this.msg = err.error.message;
        return;
      }
      this.msg = "Failed to get news articles.";
    }, () => {
      console.log(this.sources)
    })
  }
}
