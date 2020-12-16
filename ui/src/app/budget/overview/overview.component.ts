import { Component, OnInit } from '@angular/core';

import { faEye, faEyeSlash, IconDefinition} from '@fortawesome/free-solid-svg-icons';

@Component({
  selector: 'app-overview',
  templateUrl: './overview.component.html',
  styleUrls: ['./overview.component.css']
})
export class OverviewComponent implements OnInit {
  open: IconDefinition;
  close: IconDefinition;
  minimize: boolean

  constructor() { 
  }

  ngOnInit(): void {
    this.open = faEye
    this.close = faEyeSlash
    this.minimize = true
  }

  isMinimized(boolean) {
    this.minimize = boolean
  }

}
