import { ComponentFixture, TestBed } from '@angular/core/testing';
import { NavbarComponent } from './navbar.component';
import { Router } from '@angular/router';
import { NO_ERRORS_SCHEMA } from '@angular/core';

describe('NavbarComponent', () => {
  let component: NavbarComponent;
  let fixture: ComponentFixture<NavbarComponent>;
  let routerSpy = jasmine.createSpyObj('Router', ['navigate']);
  let modalInstanceSpy: jasmine.SpyObj<{ show: () => void }>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [NavbarComponent], // Standalone component
      providers: [{ provide: Router, useValue: routerSpy }],
      schemas: [NO_ERRORS_SCHEMA] // Ignore unknown elements and attributes
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(NavbarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();

    // Correctly mock Bootstrap Modal constructor
    modalInstanceSpy = jasmine.createSpyObj('ModalInstance', ['show']);
    (window as any).bootstrap = {
      Modal: jasmine.createSpy().and.returnValue(modalInstanceSpy) // Return the mocked instance
    };
  });

  it('should create the NavbarComponent', () => {
    expect(component).toBeTruthy();
  });

  it('should call openAbout() and show the modal', () => {
    // Mock the modal element
    const modalElement = document.createElement('div');
    modalElement.id = 'aboutModal';
    document.body.appendChild(modalElement);

    component.openAbout();

    expect((window as any).bootstrap.Modal).toHaveBeenCalledWith(modalElement);
    expect(modalInstanceSpy.show).toHaveBeenCalled();

    // Cleanup
    document.body.removeChild(modalElement);
  });

  it('should call logout() and navigate to /login', () => {
    component.logout();
    expect(routerSpy.navigate).toHaveBeenCalledWith(['/login']);
  });
});