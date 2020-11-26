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

  newsTotal: number;
  sportsTotal: number;
  techTotal: number;

  newsCount: number;
  sportsCount: number;
  techCount: number;

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
      this.newsTotal = this.getTotal(this.news.payload.articles)
      this.newsCount = 0
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
      this.sportsTotal = this.getTotal(this.sports.payload.articles)
      this.sportsCount = 0
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
      this.techTotal = this.getTotal(this.tech.payload.articles)
      this.techCount = 0
    })
  }

  getTotal (obj: any) {
    var count = 0;

    for (var property in obj) {
        if (Object.prototype.hasOwnProperty.call(obj, property)) {
            count++;
        }
    }

    return count;
  }

  nextNews(){
      this.newsCount+=1;
      if (this.newsCount == this.newsTotal) {
        this.newsCount = 0
      }
  }

  nextSports(){
      this.sportsCount+=1;
      if (this.sportsCount == this.sportsTotal) {
        this.sportsCount = 0
      }
  }

  nextTech(){
      this.techCount+=1;
      if (this.techCount == this.techTotal) {
        this.techCount = 0
      }
  }
}

