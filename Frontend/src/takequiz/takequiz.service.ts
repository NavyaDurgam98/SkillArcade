import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class TakequizService {
  private baseUrl = 'http://localhost:8080';  // Base URL for your API

  constructor(private http: HttpClient) { }

  // Fetch quiz topics for a specific category and subcategory
  getQuizTopics(category: string, subcategory: string): Observable<string[]> {
    const url = `${this.baseUrl}/categories/${category}/subcategories/${subcategory}/quiz_topics`;
    return this.http.get<string[]>(url);
  }

  // Fetch quiz data for a specific topic
  getQuizData(quizTopic: string): Observable<any> {
    return this.http.get<any>(`${this.baseUrl}/quiz/${quizTopic}`);
  }
}
