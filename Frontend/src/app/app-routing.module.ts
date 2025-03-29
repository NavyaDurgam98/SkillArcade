import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { SignupComponent } from '../signup/signup.component';
import {DashboardComponent} from '../dashboard/dashboard.component'
import { CategoryComponent } from '../category/category.component';
import { QuizComponent } from '../quiz/quiz.component';
import { TakeQuizComponent } from '../takequiz/takequiz.component';
import {ForgotPasswordComponent} from '../forgot-password/forgot-password.component';
import { LeaderboardComponent } from '../leaderboard/leaderboard.component';



const routes: Routes = [
  { path: 'login', loadChildren: () => import('../login/login.module').then(m => m.LoginModule) },
  { path: 'signup', component: SignupComponent },
  { path: 'dashboard', component: DashboardComponent },
  { path: 'leaderboard', component: LeaderboardComponent },
  { path: ':category', component: CategoryComponent } ,
  { path: ':category/:sub_category', component: QuizComponent },
  { path: ':category/:subcategory/:quizTopic/takequiz', component: TakeQuizComponent },
  { path: 'forgot-password', component: ForgotPasswordComponent },
  { path: '', redirectTo: '/login', pathMatch: 'full' },
];

@NgModule({
  imports: [RouterModule.forRoot(routes, { enableTracing: true })],
  exports: [RouterModule],
})
export class AppRoutingModule { }

