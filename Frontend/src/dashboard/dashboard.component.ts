

import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { DashboardService } from './dashboard.service';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css'],
  standalone: true, 
  imports: [CommonModule] // Import CommonModule here
})
export class DashboardComponent implements OnInit {
  categories: any[] = [];

  constructor(private dashboardService: DashboardService, private router: Router) {}

  ngOnInit() {
    this.dashboardService.getCategories().subscribe(data => {
      this.categories = data.map(category => ({
        ...category,
        imgPath: category.imgPath 
      }));
      console.log("Categories with images:", this.categories);
    });
  }


  goToCategory(category: string) {
    this.router.navigate([`/${category}`]);
  }
}