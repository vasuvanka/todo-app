import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment';
import { User } from '../models/user';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor(private http: HttpClient) { }

  getProfile(id: string): Promise<User> {
    return this.http.get<User>(`${environment.api}/users/${id}`).toPromise()
  }
}
