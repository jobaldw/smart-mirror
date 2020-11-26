import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment';

const newsAPIURL: string = environment.newsAPIURL;

@Injectable({
  providedIn: 'root'
})

export class NotificationsService {
  constructor(private http: HttpClient) { }

  // TODO: may need to configure: AP, ESPN, TechCrunch
  getSources() {
    return this.http.get(`${environment.newsAPIURL}/source`)
  }

  getArticles(source: string) { 
    return this.http.get(`${environment.newsAPIURL}/source/${source}/article`)
  }
}
