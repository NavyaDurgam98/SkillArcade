import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { CommonModule } from '@angular/common';
import { Router, RouterModule } from '@angular/router';

interface UserProfile {
  userName: string;
  firstName: string;
  lastName: string;
  email: string;
  quizzesTaken: number;
  rank: number;
}

@Component({
  selector: 'userProfile',
  templateUrl: './userProfile.component.html',
  styleUrls: ['./userProfile.component.css'],
  standalone: true,
  imports: [CommonModule, RouterModule]
})
export class UserProfileComponent {
  userProfile: UserProfile | null =null;
  defaultUserPhoto = '../assets/user-icon.png';
  private apiUrl = 'http://localhost:8080/userProfile';

  constructor(private http: HttpClient, private router: Router) {
    console.log('Constructor called'); // Check if constructor is executed
    this.loadUserProfile('testUser');
  }

  loadUserProfile(username: string): void {
    const url = `${this.apiUrl}?username=${username}`;
    this.http.get<UserProfile>(url).subscribe({
      next: (data) => {
        console.log('API response:', data);
        this.userProfile = data;
      },
      error: (error) => {
        console.error('Error fetching profile', error);
      },
      complete: () => {
        console.log('User profile fetch complete');
      }
    });
  }
}
