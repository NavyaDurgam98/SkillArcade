import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ActiveComponentService} from '../app/active.component.service';

declare var bootstrap: any;

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {
  searchText: string = ''; 
  activeComponent: string = '';

  constructor(private router: Router,private activeComponentService:ActiveComponentService) {}

  ngOnInit(): void {
   
  }

  openAbout() {
    let modal = new bootstrap.Modal(document.getElementById('aboutModal'));
    modal.show();
  }

  logout() {
    localStorage.removeItem('authToken');
    localStorage.removeItem('userId');
    this.router.navigate(['/login']);
  }

 onSearch(event: any) {
  const searchText = event.target.value;
    if (searchText.length >= 3 || searchText.length === 0) {
    this.activeComponentService.setSearchText(searchText);
  }
}
 navigateToProfile() {
    this.router.navigate(['/profile']);
  }
navigateToLeaderBoard(){
    this.router.navigate(['/leaderboard']);
  }
navigateToDashboard(){
  this.router.navigate(['/dashboard']);
}
}
