import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { ActiveComponentService} from '../app/active.component.service';

declare var bootstrap: any;

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent {

  constructor(private router: Router,private activeComponentService:ActiveComponentService) {}

  // To open the 'About' modal
  openAbout() {
    let modal = new bootstrap.Modal(document.getElementById('aboutModal'));
    modal.show();
  }
  
  // To handle logout and navigate to the login page
  logout() {
    this.router.navigate(['/login']);
  }
  onSearch(event: any) {
    const searchText = event.target.value;
    this.activeComponentService.setSearchText(searchText);
  }
  
  navigateToProfile() {
    this.router.navigate(['/profile']);
  }
  navigateToLeaderBoard(){
    this.router.navigate(['/leaderboard']);
  }
}
