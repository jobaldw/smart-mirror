import { Component, OnInit } from '@angular/core';
import { Store } from '@ngrx/store';
import { SET_LOCATION } from '../../../location-reducer';
import { NgForm } from '@angular/forms';

@Component({
  selector: 'app-search',
  templateUrl: './search.component.html',
  styleUrls: ['./search.component.css']
})
export class SearchComponent implements OnInit {

  loc: string;

  constructor(private store: Store<any>) { }

  ngOnInit() {
  }

  search(searchForm: NgForm) {
    if (searchForm.invalid) {
      return;
    }
    this.store.dispatch({ type: SET_LOCATION, payload: this.loc });
  }

}
