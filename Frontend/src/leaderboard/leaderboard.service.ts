import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, of } from 'rxjs';
import { catchError } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class LeaderboardService {
  private apiUrl = 'http://localhost:8080/leaderboard'; // Your actual API base URL

  constructor(private http: HttpClient) {}

  // Fetch the leaderboard data
  getLeaderboard(): Observable<any[]> {
    return this.http.get<any[]>(this.apiUrl).pipe(
      catchError(error => {
        console.error('API error:', error);
        return of([]); // Return an empty array on error
      })
    );
  }

  // Fetch the current user's rank data
  getUserRank(userId: string): Observable<any> {
    return this.http.get<any>(`${this.apiUrl}?user_id=${userId}`).pipe(
      catchError(error => {
        console.error('API error:', error);
        return of({ rank: 0, quizzes_taken: 0 }); // Return a default value on error
      })
    );
  }
}
