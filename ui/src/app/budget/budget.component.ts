import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-budget',
  templateUrl: './budget.component.html',
  styleUrls: ['./budget.component.css']
})
export class BudgetComponent implements OnInit {

  overview: string

  constructor() { 
    this.overview = "none"
  }

  ngOnInit(): void {
  }

  budgetOverview(person: string) {
    this.overview = person
  }

}
