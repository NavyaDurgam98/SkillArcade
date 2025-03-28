import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class QuizService {
  private baseUrl = 'http://localhost:8080/categories/';

  constructor(private http: HttpClient) {}

  getQuizTopics(categoryName: string, subCategoryName: string): Observable<any> {
    const formattedCategory = encodeURIComponent(categoryName.trim());
    const formattedSubCategory = encodeURIComponent(subCategoryName.trim());
    console.log('Formatted category URL:', formattedCategory);
    const url = `${this.baseUrl}${formattedCategory}/subcategories/${formattedSubCategory}/quiz_topics`;
    console.log('Fetching data from:', url);
    return this.http.get<any>(url);
  }
}
