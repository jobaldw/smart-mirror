import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class BudgetService {
  constructor(private http: HttpClient) { }

  getBudgets(id: string) {
    return this.http.get(`${environment.budgetAPIURL}/budget/${id}`)
  }

  updateBudget(budget: any) {
    return this.http.put(`${environment.budgetAPIURL}/budget`, budget)
  }
}
