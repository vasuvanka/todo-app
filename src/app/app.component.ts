import { AfterViewInit, Component, OnDestroy, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Subject, Subscription } from 'rxjs';
import { JwtPayload } from './core/models/user';
import { AuthService } from './core/services/auth.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit, OnDestroy, AfterViewInit {
  user: JwtPayload;
  userSub: Subscription = Subscription.EMPTY;
  constructor(private authService: AuthService, private router: Router) { }
  ngOnInit() {
    this.userSub = this.authService.getUser().subscribe(u => {
      setTimeout(() => {
        this.user = u
      })
    })

  }

  ngAfterViewInit() {
    if (this.authService.isLoggedIn()) {
      setTimeout(() => {
        this.user = this.authService.getJwtPayload() as JwtPayload
      })
    }
  }

  ngOnDestroy() {
    this.userSub.unsubscribe()
  }

  logout() {
    this.user = null
    this.authService.logout()
    this.router.navigate([''])
  }
}
