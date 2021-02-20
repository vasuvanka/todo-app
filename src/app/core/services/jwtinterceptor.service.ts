import { Injectable } from '@angular/core';
import { HttpRequest, HttpHandler, HttpEvent, HttpInterceptor } from '@angular/common/http';
import { Observable } from 'rxjs';
import { StorageService } from './storage.service';


@Injectable()
export class JwtInterceptor implements HttpInterceptor {
    constructor(private storageService: StorageService) { }

    intercept(request: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
        // add auth header with jwt if user is logged in and request is to api url
        const isApiUrl = request.url.includes('api');
        if (isApiUrl) {
            const token = this.storageService.getItem('token');
            if (token) {
                request = request.clone({
                    setHeaders: {
                        Authorization: `Bearer ${token}`
                    }
                });
            }
        }

        return next.handle(request);
    }
}