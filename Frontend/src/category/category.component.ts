import { Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { CategoryService } from './category.service';
import { CommonModule } from '@angular/common';
import { ActiveComponentService } from '../app/active.component.service'
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-category',
  templateUrl: './category.component.html',
  styleUrls: ['./category.component.css'],
  standalone: true, 
  imports: [CommonModule] 
})
export class CategoryComponent implements OnInit,OnDestroy {
  categoryName: string = ''; 
  subCategories: any[] = []; 
  searchText: string = ''; 
  searchSubscription: Subscription | null = null; 
  constructor(private route: ActivatedRoute, private router: Router, private categoryService: CategoryService,private activeComponentService:ActiveComponentService) {}

  ngOnInit() {
    this.categoryName = this.route.snapshot.paramMap.get('category') || '';

    this.searchSubscription = this.activeComponentService.getSearchText().subscribe(searchText => {
      this.searchText = searchText;
      this.loadCategories();  
    });

    this.loadCategories();
  }
  
  ngOnDestroy() {
    if (this.searchSubscription) {
      this.searchSubscription.unsubscribe();
    }
  }
  loadCategories() {
    if (this.searchText.trim()) {
      this.categoryService.searchCategories(this.categoryName, this.searchText).subscribe(data => {
        this.subCategories = data;
        console.log("Filtered subcategories of", this.categoryName, ":", this.subCategories);
      });
    } else {
      this.categoryService.getCategories(this.categoryName).subscribe(data => {
        this.subCategories = data;
        console.log("All subcategories of", this.categoryName, ":", this.subCategories);
      });
    }
  }
  goToQuizTopic(subCategory: string) {
    console.log(`Navigating to: /${this.categoryName}/${subCategory}`);
    this.router.navigate([`/${this.categoryName}/${subCategory}`]); 
  }
  
}
