import { Component, OnInit } from '@angular/core';
import { LeaderboardService } from './leaderboard.service';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';

@Component({
  selector: 'app-leaderboard',
  templateUrl: './leaderboard.component.html',
  styleUrls: ['./leaderboard.component.css'],
  standalone : true,
  imports : [ CommonModule, RouterModule]
})
export class LeaderboardComponent implements OnInit {

  currentUserRank: number = 0;
  attemptsMade: number = 0;
  leaderboardData: any[] = [];
  sortColumn: string = 'rank';
  sortDirection: boolean = true;
  isLoading: boolean = true;
  currentUser: any;

  constructor(private leaderboardService: LeaderboardService) {}

  ngOnInit(): void {
    this.loadLeaderboard();
  }

  loadLeaderboard(): void {
    const userId :string|null= localStorage.getItem('userId');
    this.isLoading = true;  // Set loading state to true before fetching data

    // Fetch full leaderboard first
    this.leaderboardService.getLeaderboard().subscribe({
      next: (data) => {
        console.log("Leaderboard API Response:", data);
        this.leaderboardData = data; // Store leaderboard data

        // Fetch current user's data
        this.leaderboardService.getUserRank(userId).subscribe({
          next: (userData) => {
            console.log("User Data:", userData);
            this.currentUserRank = userData.rank ?? 0;
            this.attemptsMade = userData.quizzes_taken ?? 0;
            this.currentUser = userData; // Store current user's data
          },
          error: (error) => {
            console.error("Error fetching user rank:", error);
            this.currentUserRank = 0;
            this.attemptsMade = 0;
          },
          complete: () => {
            this.isLoading = false;
            console.log("User data fetched successfully.");
          }
        });
      },
      error: (error) => {
        console.error("Error fetching leaderboard data:", error);
        this.isLoading = false;
      },
      complete: () => {
        console.log("Leaderboard data fetched successfully.");
      }
    });
  }

  toggleSort(column: string): void {
    if (this.sortColumn === column) {
      this.sortDirection = !this.sortDirection;
    } else {
      this.sortColumn = column;
      this.sortDirection = true;
    }
    this.sortLeaderboard();
  }

  sortLeaderboard(): void {
    this.leaderboardData.sort((a, b) => {
      const aValue = a[this.sortColumn];
      const bValue = b[this.sortColumn];
      if (aValue < bValue) {
        return this.sortDirection ? -1 : 1;
      } else if (aValue > bValue) {
        return this.sortDirection ? 1 : -1;
      }
      return 0;
    });
  }
}
