import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { catchError, map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class TakequizService {
  private baseUrl = 'http://localhost:8080';  // Base URL for your API

  constructor(private http: HttpClient) { }

  // Fetch quiz topics for a specific category and subcategory
  getQuizTopics(category: string, subcategory: string): Observable<string[]> {
    const url = `${this.baseUrl}/categories/${category}/subcategories/${subcategory}/quiz_topics`;
    return this.http.get<{ quiz_topics: any[] }>(url).pipe(
      map((res: { quiz_topics: any; }) => res.quiz_topics),
      catchError((error) => {
        console.error('Error fetching quiz topics:', error);
        throw error;
      })
    );
  }

  // Fetch quiz data for a specific topic
  getQuizData(quizTopic: string): Observable<any> {
    const url = `${this.baseUrl}/quiz/${quizTopic}`;
    return this.http.get<any>(url).pipe(
      catchError((error) => {
        console.error('Error fetching quiz data:', error);
        throw error;  // Rethrow error to handle it elsewhere (e.g., in component)
      })
    );
  }

  submitQuizResults(payload: {
    user_id: string;
    quiz_topic_id: string | null;
    quiz_topic_name: string;
    Score: number;
  }): Observable<any> {
    return this.http.post(`${this.baseUrl}/submitquiz`, payload);
  }
}
