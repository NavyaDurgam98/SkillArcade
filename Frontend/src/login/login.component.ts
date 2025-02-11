import { Component } from '@angular/core';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule]
})
export class LoginComponent {
  //re-do API integration once CORS issue is fixed.
  loginForm: FormGroup;
  isForgotPassword = false;
  errorMessage = '';

  constructor(private fb: FormBuilder, private router: Router, private http: HttpClient) {
    this.loginForm = this.fb.group({
      username: ['', Validators.required],
      password: ['', Validators.required]
    });
  }

  navigateToSignup() {
    this.router.navigate(['/signup']);
  }

  forgotPassword() {
    this.isForgotPassword = true;
    this.loginForm.get('password')?.disable();
  }

  cancelForgotPassword() {
    this.isForgotPassword = false;
    this.loginForm.get('password')?.enable();
  }

  onSubmit() {
    Object.keys(this.loginForm.controls).forEach(field => {
      const control = this.loginForm.get(field);
      control?.markAsTouched();
    });
    if (this.isForgotPassword) {
      const username = this.loginForm.get('username')?.value;
      if (!username) {
        this.errorMessage = 'Please enter your username to reset the password.';
        return;
      }
      this.http.post('http://localhost:8080/api/forgot-password', { username })
        .subscribe({
          next: () => {
            alert('Password reset instructions have been sent to your email.');
            this.isForgotPassword = false;
            this.loginForm.get('password')?.enable();
          },
          error: () => {
            this.errorMessage = 'Error resetting password. Please try again later.';
          }
        });

    } else {
      if (this.loginForm.valid) {
        const { username, password } = this.loginForm.value;
        const payload = { username: username, password: password }; 
        this.http.post<{ token: string }>('http://localhost:8080/signin', payload).subscribe({
          next: (response) => {
            localStorage.setItem('authToken', response.token); 
            this.router.navigate(['/dashboard']); 
          },
          error: () => {
            this.errorMessage = 'Invalid username or password. Please try again.';
          }
        });
      }
    }
  }
}
