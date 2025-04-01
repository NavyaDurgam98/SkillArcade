import { Component, OnInit, OnDestroy } from '@angular/core';
import { Router } from '@angular/router';
import { DashboardService } from './dashboard.service';
import { CommonModule } from '@angular/common';
import { ActiveComponentService } from '../app/active.component.service';
import { Subscription } from 'rxjs';

declare var bootstrap: any;   

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css'],
  standalone: true, 
  imports: [CommonModule]
})
export class DashboardComponent implements OnInit, OnDestroy {
  categories: any[] = [];
  searchText: string = ''; 
  searchSubscription: Subscription | null = null; 

  constructor(
    private dashboardService: DashboardService, 
    private router: Router,
    private activeComponentService: ActiveComponentService
  ) {}

  ngOnInit() {
    this.searchSubscription = this.activeComponentService.getSearchText().subscribe((searchText: string) => {
      this.searchText = searchText; 
      this.loadCategories(); 
    });

    this.loadCategories();
  }

  loadCategories() {
    this.dashboardService.getCategories(this.searchText).subscribe({
      next: (data) => {
        this.categories = data.map(category => ({
          ...category,
          imgPath: category.imgPath 
        }));
        console.log("Categories with images:", this.categories);
      },
      error: (error) => {
        console.error('Error fetching categories:', error);
      },
      complete: () => {
        console.log('Categories fetch completed');
      }
    });
  }

  goToCategory(category: string) {
    this.router.navigate([`/${category}`]);
  }

  openAbout() {
    let modal = new bootstrap.Modal(document.getElementById('aboutModal'));
    modal.show();
  }

  logout() {
    this.router.navigate(['/login']);
  }
  
  ngOnDestroy() {
    if (this.searchSubscription) {
      this.searchSubscription.unsubscribe();
    }
  }
}
