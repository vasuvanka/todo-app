import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { JwtInterceptor } from './services/jwtinterceptor.service';
import { HTTP_INTERCEPTORS } from '@angular/common/http';



@NgModule({
  declarations: [],
  imports: [
    CommonModule
  ],
  providers: [
    { provide: HTTP_INTERCEPTORS, useClass: JwtInterceptor, multi: true },
  ]
})
export class CoreModule { }
