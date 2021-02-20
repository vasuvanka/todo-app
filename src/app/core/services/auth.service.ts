import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { StorageService } from './storage.service';
import { environment } from "../../../environments/environment";
import { GeneralResponse, JwtPayload, LoginSuccess, UserSignUp } from '../models/user';
import { Observable, Subject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  private _user: Subject<JwtPayload> = new Subject<JwtPayload>()
  constructor(private http: HttpClient,
    private storageService: StorageService) { }

  login(email: string, password: string): Promise<LoginSuccess | GeneralResponse> {
    return this.http.post<LoginSuccess | GeneralResponse>(`${environment.api}/login`, { email, password }).toPromise()
  }

  signup(user: UserSignUp): Promise<GeneralResponse> {
    return this.http.post<GeneralResponse>(`${environment.api}/signup`, user).toPromise()
  }

  getUser() {
    return this._user
  }

  updateUser(payload: JwtPayload) {
    this._user.next(payload)
  }

  logout() {
    this.storageService.clear()
  }

  isLoggedIn(): boolean {
    const token = this.storageService.getItem('token')
    if (token) {
      this.updateUser(token)
    }
    return !!token;
  }

  decodeToken(token: string): string | JwtPayload {
    try {
      const tokenList = token.split('.')
      if (tokenList.length < 2) {
        return "invalid token"
      }
      const payload = atob(tokenList[1])
      return JSON.parse(payload.toString())
    } catch (err) {
      return (err || {}).message || 'invalid token'
    }
  }

  getJwtPayload(): string | JwtPayload {
    return this.decodeToken(this.storageService.getItem('token') || '') as JwtPayload
  }
}
