import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ActiveComponentService {
  private searchTextSubject = new BehaviorSubject<string>(''); 

  constructor() {}
  setSearchText(searchText: string) {
    this.searchTextSubject.next(searchText);
  }

  getSearchText(): Observable<string> {
    return this.searchTextSubject.asObservable(); 
  }

}
