import { Component, OnInit } from '@angular/core';
import { AbstractControl, FormControl, FormGroup, Validators } from '@angular/forms';
import { AuthService } from 'src/app/core/services/auth.service';
import { NotificationService } from 'src/app/core/services/notification.service';

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.css']
})
export class SignupComponent {

  signupForm = new FormGroup({
    email: new FormControl('', [Validators.required, Validators.email]),
    password: new FormControl('', [Validators.required, Validators.minLength(4)]),
    name: new FormControl('', [Validators.required, Validators.minLength(1)]),
  })

  constructor(private authService: AuthService, private notificationService: NotificationService) { }

  async onSignup() {
    try {
      const response = await this.authService.signup(this.signupForm.value)
      this.notificationService.notify(response.message)
      this.signupForm.reset({})
    } catch (err) {
      console.log(err);

      this.notificationService.notify((err.error || {}).message || 'Something went wrong!')
    }
  }

  get name(): AbstractControl {
    return this.signupForm.get('name')!
  }
  get email(): AbstractControl {
    return this.signupForm.get('email')!
  }
  get password(): AbstractControl {
    return this.signupForm.get('password')!
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

  getNameErrorMessage() {
    if (this.name.hasError('required')) {
      return 'You must enter a value';
    }
    return 'Name should be greater than 1 char in length';
  }

}
