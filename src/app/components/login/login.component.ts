import { Component, OnInit } from '@angular/core';
import { AbstractControl, FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { GeneralResponse, JwtPayload, LoginSuccess } from 'src/app/core/models/user';
import { AuthService } from 'src/app/core/services/auth.service';
import { NotificationService } from 'src/app/core/services/notification.service';
import { StorageService } from 'src/app/core/services/storage.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {

  loginForm = new FormGroup({
    email: new FormControl('', [Validators.required, Validators.email]),
    password: new FormControl('', [Validators.required, Validators.minLength(4)])
  })

  constructor(private authService: AuthService,
    private storageService: StorageService,
    private notificationService: NotificationService,
    private router: Router) { }

  async onLogin() {
    try {
      const { email, password } = this.loginForm.value
      const response = await this.authService.login(email, password)
      if ((response as LoginSuccess).token) {
        this.storageService.setItem('token', (response as LoginSuccess).token)
        this.authService.updateUser(this.authService.decodeToken((response as LoginSuccess).token) as JwtPayload)
        this.router.navigateByUrl('/home/dir')
        return
      }
      throw new Error((response as GeneralResponse).message)
    } catch (err) {
      console.log(err);

      this.notificationService.notify((err.error || {}).message || 'Something went wrong!')
    }
  }

  get email(): AbstractControl {
    return this.loginForm.get('email')!
  }
  get password(): AbstractControl {
    return this.loginForm.get('password')!
  }

  getEmailErrorMessage() {
    if (this.email.hasError('required')) {
      return 'You must enter a value';
    }

    return this.email.hasError('email') ? 'Not a valid email' : '';
  }

  getPasswordErrorMessage() {
    if (this.password.hasError('required')) {
      return 'You must enter a value';
    }
    return 'Password should be greater than 4 char in length';
  }
}
