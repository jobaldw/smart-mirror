import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { GoogleHubComponent } from './google-hub.component';

describe('GoogleHubComponent', () => {
  let component: GoogleHubComponent;
  let fixture: ComponentFixture<GoogleHubComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ GoogleHubComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(GoogleHubComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
