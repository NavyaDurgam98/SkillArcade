import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { DashboardService } from './dashboard.service';
import { CommonModule } from '@angular/common';

declare var bootstrap: any; 

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css'],
  standalone: true,
  imports: [CommonModule]
})
export class DashboardComponent implements OnInit {
  categories: any[] = [];

  constructor(private dashboardService: DashboardService, private router: Router) {}

  ngOnInit() {
    this.dashboardService.getCategories().subscribe(data => {
      this.categories = data.map((category, index) => ({
        ...category,
        image: `assets/${index + 1}.jpg`
      }));
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
}
