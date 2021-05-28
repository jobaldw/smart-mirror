import { Component, OnInit } from '@angular/core';
import { Store } from '@ngrx/store';
import { NgForm } from '@angular/forms';

import { BudgetService } from '../../services/budget/budget.service';
import { faCog, faTimes, faTrashAlt, faPlus, faPencilAlt, IconDefinition } from '@fortawesome/free-solid-svg-icons';

@Component({
  selector: 'app-overview',
  templateUrl: './overview.component.html',
  styleUrls: ['./overview.component.css']
})
export class OverviewComponent implements OnInit {
  budgetId: string;
  msg: string;

  add: IconDefinition;
  edit: IconDefinition;
  stop: IconDefinition;
  delete: IconDefinition;
  settings: IconDefinition;

  editing: boolean;
  editingElement: number;
  elementDesc: string;
  elementCost: number;
  object: any = <any>{};
  desc: string;
  cost: number;

  constructor(
    private store: Store<any>,
    private BudgetService: BudgetService
  ) {
    this.add = faPlus
    this.edit = faPencilAlt
    this.stop = faTimes
    this.delete = faTrashAlt
    this.settings = faCog

    this.editing = false

    this.budgetId = "5fdca740fa4d94ae49bd5e68"
  }

  ngOnInit(): void {
    this.getBudgets(this.budgetId)

  }

  getBudgets(id: string) {
    let obj: any
    this.BudgetService.getBudgets(id)
      .subscribe(res => {
        obj = res;
      }, err => {
        if (err.error && err.error.message) {
          this.msg = err.error.message;
          return;
        }
        this.msg = "Failed to get Budget.";
      }, () => {
        this.calculate(obj)
      })
  }

  test: any = <any>{};

  deleteElement(item:  any = <any>{}) {
    this.test = item
  }


  updatelement(i: number) {
    this.editingElement = i
    this.elementDesc = this.object.records[i].desc
    this.elementCost = this.object.records[i].cost
  }

  updateBudget(searchForm: NgForm) {
    let obj: any

    // this.BudgetService.updateBudget(budget)
    // .subscribe(res => {
    //   obj = res;
    // }, err => {
    //   if (err.error && err.error.message) {
    //     this.msg = err.error.message;
    //     return;
    //   }
    //   this.msg = "Failed to get Budget.";
    // }, () => {
    //   this.calculate(obj)
    // })
    // location.reload
  }


  selectBudget(id: string) {
    this.budgetId = id
    this.getBudgets(this.budgetId)
  }

  salary: number;
  check: number;

  // Income
  benefits: any = <any>{};
  bTotal: number;

  taxes: any = <any>{};
  tTotal: number;

  invest: any = <any>{};
  iTotal: number;
  jTotal: number;

  // Expense
  expenses: any = <any>{};
  eTotals: number[] = [];
  nTotal: number;
  uTotal: number;

  // Savings
  savings: any = <any>{};
  sTotals: number[] = [];
  sTotal: number;
  gTotal: number;

  calculate(obj: any) {
    let total = 0
    this.salary = obj.budget.income.salary
    this.benefits = obj.budget.income.benefits
    this.taxes = obj.budget.income.taxes
    this.invest = obj.budget.income.invests
    this.expenses = obj.budget.expense
    this.savings = obj.budget.savings

    this.tTotal = 0
    for (let index = 0; index < this.taxes.records.length; index++) {
      this.tTotal += (this.taxes.records[index].percentage/100)*(this.salary/26);
    }

    this.bTotal = 0
    for (let index = 0; index < this.benefits.records.length; index++) {
      this.bTotal += this.benefits.records[index].cost;
    }

    this.iTotal = 0
    this.jTotal = 0
    for (let i = 0; i < this.invest.length; i++) {
      for (let j = 0; j < this.invest[i].records.length; j++) {
        if (this.invest[i].records[j].given != true) {
          this.iTotal += (this.invest[i].records[j].percentage/100)*(this.salary/26);
        } else {
          this.jTotal += (this.invest[i].records[j].percentage/100)*(this.salary/26);
        }
      }
    }

    this.check = (this.salary / 26) - (this.tTotal + this.bTotal + this.iTotal)

    this.nTotal = 0
    this.uTotal = 0
    for (let i = 0; i < this.expenses.length; i++) {
      total = 0
      for (let j = 0; j < this.expenses[i].records.length; j++) {
        total += this.expenses[i].records[j].cost
        if (this.expenses[i].need == true) {
          this.nTotal += this.expenses[i].records[j].cost;
        } else {
          this.uTotal += this.expenses[i].records[j].cost;
        }
      }
      this.eTotals.push(total)
    }

    this.sTotal = 0
    this.gTotal = 0
    for (let i = 0; i < this.savings.length; i++) {
      total = 0
      for (let j = 0; j < this.savings[i].records.length; j++) {
        total += this.savings[i].records[j].cost
        if (this.savings[i].records[j].given == true) {
          this.gTotal += this.savings[i].records[j].cost;
        } else {
          this.sTotal += this.savings[i].records[j].cost;
        }
      }
      this.sTotals.push(total)
    }
  }

  toggleOverlay(obj: any) {
    if (this.editing == false) {
      this.editing = true
      this.object = obj
    } else {
      this.editing = false
    }
  }
}