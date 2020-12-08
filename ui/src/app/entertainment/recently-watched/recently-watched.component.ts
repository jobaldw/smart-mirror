import { Component, OnInit } from '@angular/core';
import { Store } from '@ngrx/store';
import { NgForm } from '@angular/forms';

import { EntertainmentService } from '../../services/entertainment/entertainment.service';

@Component({
  selector: 'app-recently-watched',
  templateUrl: './recently-watched.component.html',
  styleUrls: ['./recently-watched.component.css']
})
export class RecentlyWatchedComponent implements OnInit {
  resp: any = <any>{};
  
  title: string
  genre: string
  platform: string  
  movie: boolean
  series: boolean
  
  msg: string;


  constructor(
    private store: Store<any>,
    private EntertainmentService: EntertainmentService
  ) {  
    this.getTitles()
  }

  ngOnInit(): void {
  }

  search(searchForm: NgForm) {
    let type = ""
    if (this.movie) {
      type = "movie"
      if (this.series) {
        type = ""
      }
    } else if (this.series) {
      type = "series"
      if (this.movie) {
        type = ""
      }
    } else {
      type = ""
    }

    this.EntertainmentService.getTitles(this.title, this.genre, this.platform, type)
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

  getTitles() {
    this.EntertainmentService.getTitles("", "", "", "")
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

  getPlatform(int: number): string {
    if (this.resp.elements[int].platform == "Disney+") {
      return "disney"
    }
    if (this.resp.elements[int].platform == "HBO Max") {
      return "hbo"
    }
    if (this.resp.elements[int].platform == "Hulu") {
      return "hulu"
    }
    if (this.resp.elements[int].platform == "Netflix") {
      return "netflix"
    }
    if (this.resp.elements[int].platform == "Plex") {
      return "plex"
    }
    if (this.resp.elements[int].platform == "Prime Video") {
        return "prime"
    }
    if (this.resp.elements[int].platform == "In Theaters") {
        return "theaters"
    }
    return ""
  }
}
