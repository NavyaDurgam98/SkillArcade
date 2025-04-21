import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class QuizService {
  private baseUrl = `${environment.protectedApiUrl}/categories/`;

  constructor(private http: HttpClient) {}

  getQuizTopics(categoryName: string, subCategoryName: string, searchText: string = ''): Observable<any> {
    const formattedCategory = encodeURIComponent(categoryName.trim());
    const formattedSubCategory = encodeURIComponent(subCategoryName.trim());
    const params = new HttpParams().set('searchText', searchText); 
    const url = `${this.baseUrl}${formattedCategory}/subcategories/${formattedSubCategory}/quiz_topics`;
    return this.http.get<any>(url, { params });
  }
}

