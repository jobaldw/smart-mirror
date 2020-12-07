import { Component, OnInit } from '@angular/core';
import { Store } from '@ngrx/store';

import { EntertainmentService } from '../../services/entertainment/entertainment.service';

@Component({
  selector: 'app-recently-watched',
  templateUrl: './recently-watched.component.html',
  styleUrls: ['./recently-watched.component.css']
})
export class RecentlyWatchedComponent implements OnInit {
  resp: any = <any>{};
  msg: string;

  constructor(
    private store: Store<any>,
    private EntertainmentService: EntertainmentService
  ) {  
    this.getTitles()
  }

  ngOnInit(): void {
  }

  getTitles() {
    this.EntertainmentService.getTitles("", "")
    .subscribe(res => {
      this.resp = res;
    }, err => {
      if (err.error && err.error.message) {
        this.msg = err.error.message;
        return;
      }
      this.msg = "Failed to get titles.";
    }, () => {
    })
  }

  getPlatform(int): string {
    if (this.resp.elements[int].platform == "Netflix") {
        return "netflix"
    }
    if (this.resp.elements[int].platform == "Hulu") {
        return "hulu"
    }
    if (this.resp.elements[int].platform == "HBO Max") {
        return "hbo"
    }
    if (this.resp.elements[int].platform == "Disney+") {
        return "disney"
    }
    if (this.resp.elements[int].platform == "Plex") {
        return "plex"
    }
    if (this.resp.elements[int].platform == "In Theaters") {
        return "theaters"
    }
    return ""
  }
}
