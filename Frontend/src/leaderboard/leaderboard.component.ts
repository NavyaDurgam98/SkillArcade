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
//export class LeaderboardComponent implements OnInit {
  export class LeaderboardComponent {

  currentUserRank: number = 0;
  attemptsMade: number = 0;
  leaderboardData: any[] = [];
  sortColumn: string = 'rank';
  sortDirection: boolean = true;
  isLoading: boolean = true;

  constructor(private leaderboardService: LeaderboardService) {}

  // ngOnInit(): void {
  //   this.loadLeaderboard();
  // }

  loadLeaderboard(): void {
    const userId = 123; // Replace with actual logged-in user ID
    this.leaderboardService.getRankings(userId).subscribe({
      next: (data) => {
        console.log("API Response:", data);
        this.currentUserRank = data.currentRank ?? 0;
        this.attemptsMade = data.attempts ?? 0;
        this.leaderboardData = data.rankings ?? [];
      },
      error: (error) => {
        console.error("Error fetching leaderboard data:", error);
        this.currentUserRank = 0;
        this.attemptsMade = 0;
        this.leaderboardData = [];
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
