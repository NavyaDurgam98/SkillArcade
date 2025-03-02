import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class TakequizService {

  private jsonUrl = 'assets/quiz_data.json';

  constructor(private http: HttpClient) { }

  // Fetch quiz data
  getQuizData(): Observable<any> {
    return this.http.get<any>(this.jsonUrl);
  }
}
