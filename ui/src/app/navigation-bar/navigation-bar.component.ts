import { Component, OnInit } from '@angular/core';
import { faHome, faTv, faFileInvoiceDollar, faNewspaper, faCalendar, faEllipsisH, faTools } from '@fortawesome/free-solid-svg-icons'

@Component({
  selector: 'app-navigation-bar',
  templateUrl: './navigation-bar.component.html',
  styleUrls: ['./navigation-bar.component.css']
})
export class NavigationBarComponent implements OnInit {
  faHome = faHome;
  faTv = faTv;
  faFileInvoiceDollar = faFileInvoiceDollar;
  faNewspaper = faNewspaper;
  faCalendar = faCalendar;
  faEllipsisH = faEllipsisH;
  faTools = faTools;

  constructor() { }

  ngOnInit(): void {
  }

}
