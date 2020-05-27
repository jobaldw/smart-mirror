import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { RecentlyWatchedComponent } from './recently-watched.component';

describe('RecentlyWatchedComponent', () => {
  let component: RecentlyWatchedComponent;
  let fixture: ComponentFixture<RecentlyWatchedComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ RecentlyWatchedComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(RecentlyWatchedComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
