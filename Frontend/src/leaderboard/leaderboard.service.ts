import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, of } from 'rxjs';
import { catchError } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class LeaderboardService {
  private apiUrl = `${environment.protectedApiUrl}/leaderboard`

  constructor(private http: HttpClient) {}

  // Fetch the leaderboard data
  getLeaderboard(): Observable<any[]> {
    return this.http.get<any[]>(this.apiUrl).pipe(
      catchError(error => {
        console.error('API error:', error);
        return of([]); 
      })
    );
  }

  // Fetch the current user's rank data
  getUserRank(userId: string|null): Observable<any> {
    return this.http.get<any>(`${this.apiUrl}?user_id=${userId}`).pipe(
      catchError(error => {
        console.error('API error:', error);
        return of({ rank: 0, quizzes_taken: 0 }); 
      })
    );
  }
}
