import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { SignupComponent } from '../signup/signup.component';
import {DashboardComponent} from '../dashboard/dashboard.component'
import { CategoryComponent } from '../category/category.component';
import { QuizComponent } from '../quiz/quiz.component';
import { TakeQuizComponent } from '../takequiz/takequiz.component';
import {ForgotPasswordComponent} from '../forgot-password/forgot-password.component';
import { LeaderboardComponent } from '../leaderboard/leaderboard.component';
import {UserProfileComponent} from '../userProfile/userProfile.component'
import { AuthGuard } from './auth.gaurd';


const routes: Routes = [
  { path: 'login', loadChildren: () => import('../login/login.module').then(m => m.LoginModule) },
  { path: 'signup', component: SignupComponent },
  { path: 'profile', component: UserProfileComponent, canActivate:[AuthGuard] },
  { path: 'dashboard', component: DashboardComponent,canActivate:[AuthGuard] },
  { path: 'leaderboard', component: LeaderboardComponent,canActivate:[AuthGuard] },
  { path: ':category', component: CategoryComponent,canActivate:[AuthGuard] } ,
  { path: ':category/:sub_category', component: QuizComponent,canActivate:[AuthGuard] },
  { path: ':category/:subcategory/:quizTopic/takequiz', component: TakeQuizComponent,canActivate:[AuthGuard] },
  { path: 'forgot-password', component: ForgotPasswordComponent },
  { path: '', redirectTo: '/login', pathMatch: 'full' },
];

@NgModule({
  imports: [RouterModule.forRoot(routes, { enableTracing: true })],
  exports: [RouterModule],
})
export class AppRoutingModule { }

