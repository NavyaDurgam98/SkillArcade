import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, of } from 'rxjs';
import { catchError } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class LeaderboardService {
  private apiUrl = 'https://example.com/api/leaderboard'; // Replace with actual API URL

  constructor(private http: HttpClient) {}

  getRankings(userId: number): Observable<any> {
    return this.http.get<any>(`${this.apiUrl}?userId=${userId}`).pipe(
      catchError(error => {
        console.error('API error:', error);
        return of({ currentRank: 0, attempts: 0, rankings: [] });
      })
    );
  }
}
