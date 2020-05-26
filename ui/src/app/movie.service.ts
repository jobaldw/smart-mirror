import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class MovieService {

  movieUrl = '../assets/movie.json';

  constructor() { }

  getAlbum(id: string) {
  }
}
