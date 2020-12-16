import { Component, OnInit, Input } from '@angular/core';

@Component({
  selector: 'app-savings',
  templateUrl: './savings.component.html',
  styleUrls: ['./savings.component.css']
})
export class SavingsComponent implements OnInit {

  current: number;
  progress: number;
  total: number;
  color: string;
  title: string;

  constructor() { 
    this.title = "First Home Down Payment Goal"
    this.current = 7000
    this.progress = this.current
    this.total = 20000
  }

  ngOnInit(): void {
    //if we don't have progress, set it to 0.
    if (!this.progress) {
      this.progress = 0;
    }
    //if we don't have a total aka no requirement, it's 100%.
    if (this.total === 0) {
      this.total = this.progress;
    } else if (!this.total) {
      this.total = 100;
    }
    //if the progress is greater than the total, it's also 100%.
    if (this.progress > this.total) {
      this.progress = 100;
      this.total = 100;
    }
    this.progress = ((this.progress / this.total) * 100);
  }
}
