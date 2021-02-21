import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { DefaultDirectoryComponent } from './default-directory.component';

describe('DefaultDirectoryComponent', () => {
  let component: DefaultDirectoryComponent;
  let fixture: ComponentFixture<DefaultDirectoryComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ DefaultDirectoryComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(DefaultDirectoryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
