import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-welcome',
  templateUrl: './welcome.component.html',
  styleUrls: ['./welcome.component.css']
})
export class WelcomeComponent implements OnInit {

  public hours;
  public msg;

  constructor() {
    this.decide()
  }

  ngOnInit(): void {
  }

  decide() {                                                                 
     this.hours = new Date().getHours();   

    if (this.hours < 10) {         // < 10:00am                                                       
      this.msg = 'Good Morning';
    } else if (this.hours < 16) {  // < 4:00pm                                                 
      this.msg = 'Good Afternoon';
    } else if (this.hours < 19) {  // < 7:00pm                                                
      this.msg = 'Good Evening';
    } else if (this.hours < 24) {  // < 12:00am                                            
      this.msg = 'Good Night';
    }
  }

}
