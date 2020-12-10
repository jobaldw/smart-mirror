import { Component, OnInit } from '@angular/core';
import { Store } from '@ngrx/store';
import { NgForm } from '@angular/forms';
import { faPlus, faMinus, faTimes, IconDefinition } from '@fortawesome/free-solid-svg-icons';

import { EntertainmentService } from '../../services/entertainment/entertainment.service';

@Component({
  selector: 'app-watch-list',
  templateUrl: './watch-list.component.html',
  styleUrls: ['./watch-list.component.css']
})
export class WatchListComponent implements OnInit {
  resp: any = <any>{};
  respOMBD: any = <any>{};
  result: any = <any>{};  
  
  addedTitle: string

  title: string
  ombdTitles: string
  genre: string
  platform: string; streamer: string
  movie: boolean
  series: boolean
  
  add: IconDefinition;
  remove: IconDefinition;
  stop: IconDefinition;
  
  adding: boolean
  settingStreamer: boolean

  msg: string;

  constructor(
    private store: Store<any>,
    private EntertainmentService: EntertainmentService
  ) {  
    this.getTitles()

    this.movie = true
    this.series = true
    this.genre = ""
    this.platform = ""

    this.add = faPlus
    this.remove = faMinus
    this.stop = faTimes

    this.adding = false
    this.settingStreamer = false
  }

  ngOnInit(): void {
  }

  deleteTitle(int: number) {
    this.EntertainmentService.removeTitle(this.resp.elements[int]._id)
    .subscribe(res => {
      this.respOMBD = res;
    }, err => {
      if (err.error && err.error.message) {
        this.msg = err.error.message;
        return;
      }
      this.msg = "Failed to remove OMBD titles.";
    }, () => {
      location.reload();
    })
  }

  addTitle(title: string) {
    this.EntertainmentService.addTitles(title, this.streamer)
    .subscribe(res => {
      this.respOMBD = res;
    }, err => {
      if (err.error && err.error.message) {
        this.msg = err.error.message;
        return;
      }
      this.msg = "Failed to get OMBD titles.";
    }, () => {
      location.reload();
    })
  }

  searchOMBD(searchForm: NgForm) {
    this.EntertainmentService.searchTitles(this.ombdTitles)
    .subscribe(res => {
      this.respOMBD = res;
    }, err => {
      if (err.error && err.error.message) {
        this.msg = err.error.message;
        return;
      }
      this.msg = "Failed to find OMBD titles.";
    }, () => {
    })
  }

  queryWatchlist(searchForm: NgForm) {
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
      this.msg = "Failed to find titles.";
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

  toggleOverlay() {
    if (this.adding == false) {
      this.adding = true
    } else {
      this.adding = false
    }
  }
  
  toggleStreamer(int: number) {
    this.addedTitle = this.respOMBD.Search[int].Title
    if (this.settingStreamer == false) {
      this.settingStreamer = true
    } else {
      this.settingStreamer = false
    }
  }

  setStreamer(value: string) {
    this.streamer = value
    this.addTitle(this.addedTitle)
  }
}
