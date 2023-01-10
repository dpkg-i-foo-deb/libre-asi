import { Component } from '@angular/core';
import { FormControl, Validators } from '@angular/forms';
@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss'],
})
export class LoginComponent {
  iconName: string = 'eye_off';
  email = new FormControl('', [Validators.required]);
  password = new FormControl('', [Validators.required]);
  hidePassword: boolean = true;

  getEmailErrorMessage(): string {
    if (this.email.hasError('required')) {
      return 'This field is required';
    }
    return '';
  }

  getPasswordErrorMessage(): string {
    if (this.password.hasError('required')) {
      return 'This field is required';
    }
    return '';
  }

  toggleHidePassword() {
    if (this.hidePassword) {
      this.iconName = 'eye_on';
    } else {
      this.iconName = 'eye_off';
    }
    this.hidePassword = !this.hidePassword;
  }
}
