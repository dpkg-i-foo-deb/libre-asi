import { Component } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss'],
})
export class LoginComponent {
  iconName: string = 'eye_off';
  loginForm = new FormGroup({
    email: new FormControl('', [Validators.required]),
    password: new FormControl('', [Validators.required]),
    role: new FormControl('', [Validators.required]),
  });

  hidePassword: boolean = true;
  selectedRole: boolean = true;

  isEmailValid(): boolean {
    return this.loginForm.get('email')?.valid ?? false;
  }

  getEmailErrorMessage(): string {
    if (this.loginForm.get('email')!.hasError('required')) {
      return 'This field is required';
    }
    return '';
  }

  isPasswordValid(): boolean {
    return this.loginForm.get('password')?.valid ?? false;
  }

  getPasswordErrorMessage(): string {
    if (this.loginForm.get('password')!.hasError('required')) {
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

  submit() {
    if (this.loginForm.valid) {
      this.selectedRole = true;
    } else {
      if (!this.loginForm.get('role')?.valid) {
        this.selectedRole = false;
      } else {
        this.selectedRole = true;
      }
    }
  }
}
