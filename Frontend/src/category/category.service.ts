import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class CategoryService {
  private url = `${environment.protectedApiUrl}/categories/`; 

  constructor(private http: HttpClient) {}

  getCategories(categoryName:string): Observable<any[]> {
    return this.http.get<any[]>(this.url+categoryName);
  }
   searchCategories(categoryName: string, searchText: string): Observable<any[]> {
    let params = new HttpParams().set('searchText', searchText);
    return this.http.get<any[]>(`${this.url}${categoryName}`, { params });
  }
}
