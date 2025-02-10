import { NgModule } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { RouterModule } from '@angular/router';
import { LoginComponent } from './login.component';
import { CommonModule } from '@angular/common';
const routes = [
  { path: '', component: LoginComponent },
];

@NgModule({
  declarations: [
    
  ],
  imports: [
    FormsModule,
    ReactiveFormsModule,
    CommonModule,
    RouterModule.forChild(routes),

  ],
  providers: [],
})
export class LoginModule { }
