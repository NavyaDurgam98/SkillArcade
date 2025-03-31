import { ComponentFixture, TestBed } from '@angular/core/testing';
import { CategoryComponent } from './category.component';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { of } from 'rxjs';
import { By } from '@angular/platform-browser';
import { ActivatedRoute } from '@angular/router';

describe('CategoryComponent', () => {
  let component: CategoryComponent;
  let fixture: ComponentFixture<CategoryComponent>;

  const mockActivatedRoute = {
    snapshot: {
      paramMap: {
        get: (key: string) => 'Computer Science'
      }
    }
  };

  // Updated mock data to match component's expected structure
  const mockSubCategories = [
    { subCategory: 'Programming Languages', subImgPath: 'programming.jpg' },
    { subCategory: 'Data Structures', subImgPath: 'datastructures.jpg' },
    { subCategory: 'Algorithms', subImgPath: 'algorithms.jpg' }
  ];

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [CategoryComponent, HttpClientTestingModule],
      providers: [
        { provide: ActivatedRoute, useValue: mockActivatedRoute }
      ]
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(CategoryComponent);
    component = fixture.componentInstance;
    component.subCategories = mockSubCategories;
    fixture.detectChanges();
  });

  it('should create the category component', () => {
    expect(component).toBeTruthy();
  });

  it('should display the correct number of subcategory cards', () => {
    const subcategoryElements = fixture.debugElement.queryAll(By.css('.card'));
    expect(subcategoryElements.length).toBe(mockSubCategories.length);
  });

  it('should display correct subcategory names', () => {
    const subcategoryTitles = fixture.debugElement.queryAll(By.css('.card-title'));
    
    subcategoryTitles.forEach((element, index) => {
      expect(element.nativeElement.textContent.trim())
        .toBe(mockSubCategories[index].subCategory);
    });
  });

  it('should call goToQuizTopic() with correct subcategory when button is clicked', () => {
    spyOn(component, 'goToQuizTopic');
    
    const buttons = fixture.debugElement.queryAll(By.css('.btn-primary'));
    const firstButton = buttons[0];
    
    firstButton.nativeElement.click();
    
    expect(component.goToQuizTopic)
      .toHaveBeenCalledWith(mockSubCategories[0].subCategory);
  });

  it('should display no subcategories if the API returns an empty array', () => {
    component.subCategories = [];
    fixture.detectChanges();

    const subcategoryElements = fixture.debugElement.queryAll(By.css('.card'));
    expect(subcategoryElements.length).toBe(0);
  });
});