import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-clock',
  templateUrl: './clock.component.html',
  styleUrls: ['./clock.component.css']
})

export class ClockComponent implements OnInit {

  public today;

  constructor() { 
    setInterval(() => {
      this.today = new Date();
    }, 1000);
  }

  ngOnInit(): void {
  }
}
