import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { HttpParams } from "@angular/common/http";

import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class EntertainmentService {
  constructor(private http: HttpClient) { }

  getTitles(title: string, genre: string, platform: string, type: string) {
    let params = new HttpParams();
    if (title != undefined && title != "") {
      params = params.append('Title', title);
    }
    if (genre != undefined && genre != "" ) {
      params = params.append('Genre', genre);
    }
    if (platform != undefined && platform != "") {
      params = params.append('platform', platform);
    }
    if (type != undefined) {
      params = params.append('Type', type);
    }

    return this.http.get(`${environment.entertainmentAPIURL}/v1/watchlist`, {params: params})
  }

  searchTitles(title: string) {
    return this.http.get(`${environment.entertainmentAPIURL}/v1/search/${title}`)
  }

  addTitles(title: string, platform: string) {
    return this.http.get(`${environment.entertainmentAPIURL}/v1/retrieve/${title}?platform=${platform}`)
  }

  removeTitle(id: string) {
    return this.http.delete(`${environment.entertainmentAPIURL}/v1/watchlist/${id}`)
  }
}
