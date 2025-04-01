import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { CommonModule } from '@angular/common';
import { Router, RouterModule } from '@angular/router';

export interface UserProfile {
  username: string;
  firstname: string;
  lastname: string;
  email: string;
  quizzes_taken: number|string;
  rank: number|string;
}

export interface LeaderboardStats {
  username: string;
  total_score: number;
  quizzes_taken: number;
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
  userProfile: UserProfile | null = null;
  defaultUserPhoto = '../assets/user-icon.png';
  readonly profileApiUrl = 'http://localhost:8080/userprofile';
  readonly leaderboardApiUrl = 'http://localhost:8080/leaderboard';
  isLoading = true;
  profileLoadError = false;
  statsLoadError = false;

  constructor(private http: HttpClient, private router: Router) {
    this.initializeProfile();
    this.loadUserProfile();
  }

  initializeProfile(): void {
    this.userProfile = {
      username: ' Not Found',
      firstname: ' Not Found',
      lastname: ' Not Found',
      email: ' Not Found',
      quizzes_taken: ' Not Found',
      rank: ' Not Found'
    };
  }

  loadUserProfile(): void {
    const userId = localStorage.getItem('userId');
    if (!userId) {
      console.error('No user ID found in local storage');
      this.isLoading = false;
      this.profileLoadError = true;
      return;
    }

    // Fetch basic profile data
    const profileUrl = `${this.profileApiUrl}?user_id=${userId}`;
    this.http.get<UserProfile>(profileUrl).subscribe({
      next: (profileData) => {
        console.log('Profile API response:', profileData);
        if (profileData) {
          this.userProfile = {
            ...profileData,
            quizzes_taken: this.userProfile?.quizzes_taken || 'Not Found',
            rank: this.userProfile?.rank || 'Not Found'
          };
        }
        this.loadLeaderboardStats(userId);
      },
      error: (error) => {
        console.error('Error fetching profile', error);
        this.profileLoadError = true;
        this.isLoading = false;
      }
    });
  }

  loadLeaderboardStats(userId: string): void {
    const leaderboardUrl = `${this.leaderboardApiUrl}?user_id=${userId}`;
    this.http.get<LeaderboardStats>(leaderboardUrl).subscribe({
      next: (statsData) => {
        console.log('Leaderboard API response:', statsData);
        if (this.userProfile && statsData) {
          this.userProfile = {
            ...this.userProfile,
            quizzes_taken: statsData.quizzes_taken ?? 'Not Found',
            rank: statsData.rank ?? 'Not Found'
          };
        }
        this.isLoading = false;
      },
      error: (error) => {
        console.error('Error fetching leaderboard stats', error);
        this.statsLoadError = true;
        this.isLoading = false;
      }
    });
  }
}