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
  loginForm: FormGroup;
  isForgotPassword = false;
  errorMessage = '';

  constructor(private fb: FormBuilder, private router: Router, private http: HttpClient) {
    this.loginForm = this.fb.group({
      username: ['', Validators.required],
      password: ['', Validators.required],
      email: ['', [Validators.required, Validators.email]] 
    });
    this.loginForm.get('email')?.disable();
  }

  navigateToSignup() {
    this.router.navigate(['/signup']);
  }

  forgotPassword() {
    this.isForgotPassword = true;
    this.loginForm.get('password')?.disable();
    this.loginForm.get('email')?.enable();  
  }

  cancelForgotPassword() {
    this.isForgotPassword = false;
    this.loginForm.get('password')?.enable();
    this.loginForm.get('email')?.disable();  
  }

  onSubmit() {
    Object.keys(this.loginForm.controls).forEach(field => {
      const control = this.loginForm.get(field);
      control?.markAsTouched();
    });
    if (this.isForgotPassword) {
      const email = this.loginForm.get('email')?.value;
      this.http.post('http://localhost:8080/forgotpassword', { email })
        .subscribe({
          next: () => {
            alert('Password reset instructions have been sent to your email.');
            this.isForgotPassword = false;
            // redirect to reset password instead of the below method
            this.cancelForgotPassword();
          },
          error: () => {
            this.errorMessage = 'Error resetting password. Please try again later.';
          }
        });

    } else {
      if (this.loginForm.valid) {
        const { username, password } = this.loginForm.value;
        const payload = { username: username, password: password }; 
        this.http.post<{ token: string, user_id: string }>('http://localhost:8080/signin', payload).subscribe({
          next: (response) => {
            localStorage.setItem('authToken', response.token); 
            localStorage.setItem('userId', response.user_id); 
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
