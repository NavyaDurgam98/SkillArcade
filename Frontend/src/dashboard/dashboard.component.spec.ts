import { ComponentFixture, TestBed } from '@angular/core/testing';
import { DashboardComponent } from './dashboard.component';
import { DashboardService } from './dashboard.service';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { of } from 'rxjs';
import { By } from '@angular/platform-browser';
import { RouterTestingModule } from '@angular/router/testing';

describe('DashboardComponent', () => {
  let component: DashboardComponent;
  let fixture: ComponentFixture<DashboardComponent>;
  let dashboardService: DashboardService;

  // Modified mock data to match what the TEMPLATE expects (using 'category' instead of 'category_name')
  const mockCategories = [
    { 
      category_id: "1", 
      category: 'Web Development',  // Changed from category_name to category
      imgPath: '1.jpg',
      category_name: 'Web Development' // Keep original if service needs it
    },
    { 
      category_id: "2", 
      category: 'Cybersecurity',
      imgPath: '2.jpg',
      category_name: 'Cybersecurity'
    },
    { 
      category_id: "3", 
      category: 'AI & Machine Learning',
      imgPath: '3.jpg',
      category_name: 'AI & Machine Learning'
    }
  ];

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [DashboardComponent, HttpClientTestingModule, RouterTestingModule],
      providers: [
        {
          provide: DashboardService,
          useValue: {
            getCategories: () => of(mockCategories.map(c => ({
              ...c,
              category_name: c.category // Ensure service gets what it expects
            })))
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(DashboardComponent);
    component = fixture.componentInstance;
    dashboardService = TestBed.inject(DashboardService);
    
    // Initialize with transformed data that matches template
    component.categories = mockCategories;
    fixture.detectChanges();
  });

  it('should display categories correctly', () => {
    const titles = fixture.debugElement.queryAll(By.css('.card-title'));
    expect(titles[0].nativeElement.textContent.trim()).toBe('Web Development');
    expect(titles[1].nativeElement.textContent.trim()).toBe('Cybersecurity');
    expect(titles[2].nativeElement.textContent.trim()).toBe('AI & Machine Learning');
  });

  it('should call goToCategory with correct parameter', () => {
    spyOn(component, 'goToCategory');
    const buttons = fixture.debugElement.queryAll(By.css('.btn-primary'));
    
    buttons[0].nativeElement.click();
    expect(component.goToCategory).toHaveBeenCalledWith('Web Development');
  });
});