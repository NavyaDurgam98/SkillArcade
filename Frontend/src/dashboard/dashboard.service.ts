import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { Observable } from 'rxjs';

// Define the Category interface
interface Category {
  category_id: string;
  category_name: string;
  imgPath: string;
}

@Injectable({
  providedIn: 'root'
})
export class DashboardService {
  private url = `${environment.protectedApiUrl}/categories`; 

  constructor(private http: HttpClient) {}

  getCategories(searchText: string = ''): Observable<Category[]> {
    let params = new HttpParams();

    if (searchText) {
      params = params.set('searchText', searchText);
    }

    return this.http.get<Category[]>(this.url, { params });
  }
}
