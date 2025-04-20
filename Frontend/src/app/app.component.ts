import { Component } from '@angular/core';
import { RouterModule } from '@angular/router';
import { NavbarComponent } from '../navbar/navbar.component';
import { Router, NavigationEnd } from '@angular/router';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-root',
  // standalone: true,
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
  imports: [RouterModule,NavbarComponent, CommonModule],

})
export class AppComponent {
  hideNavbar: boolean = false;  
  constructor(private router: Router) {}
  ngOnInit() {
    // Listen for route changes
    this.router.events.subscribe((event: any) => { 
      if (event instanceof NavigationEnd) {
        // Hide navbar for login and signup pages
        this.hideNavbar = this.router.url === '/login' || this.router.url === '/signup';
      }
    });
  }
  
}
