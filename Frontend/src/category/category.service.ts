import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class CategoryService {
  private jsonUrl = 'http://localhost:8080/categories/'; 

  constructor(private http: HttpClient) {}

  getCategories(categoryName:string): Observable<any[]> {
    return this.http.get<any[]>(this.jsonUrl+categoryName);
  }
}
