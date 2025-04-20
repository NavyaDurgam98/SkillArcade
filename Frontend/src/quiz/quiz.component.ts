import { Component, OnInit, OnDestroy } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { CommonModule } from '@angular/common';
import { QuizService } from './quiz.service';
import { ActiveComponentService } from '../app/active.component.service';
import { Subscription } from 'rxjs';

interface QuizTopic {
  quiz_topic_id: string;
  quiz_topic_name: string;
  quizImgPath: string;
}

@Component({
  selector: 'app-quiz',
  standalone: true,
  templateUrl: './quiz.component.html',
  styleUrls: ['./quiz.component.css'],
  imports: [CommonModule]
})
export class QuizComponent implements OnInit, OnDestroy {
  categoryName: string = '';
  subCategoryName: string = '';
  quizTopics: QuizTopic[] = [];
  searchText: string = ''; 
  searchSubscription: Subscription | null = null; 

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private quizService: QuizService,
    private activeComponentService: ActiveComponentService
  ) {}

  ngOnInit() {
    this.categoryName = this.route.snapshot.paramMap.get('category') || '';
    this.subCategoryName = this.route.snapshot.paramMap.get('sub_category') || '';

    this.searchSubscription = this.activeComponentService.getSearchText().subscribe((searchText: string) => {
      this.searchText = searchText; 
      this.loadQuizTopics(); 
    });

    this.loadQuizTopics();
  }

  loadQuizTopics() {
    this.quizService.getQuizTopics(this.categoryName, this.subCategoryName, this.searchText).subscribe({
      next: (response: any) => {
        if (response && response.quiz_topics) {
          this.quizTopics = response.quiz_topics;
        } else {
          console.warn(`No quiz topics found for ${this.subCategoryName}`);
        }
      },
      error: (error) => {
        console.error('Error fetching quiz topics:', error);
      },
      complete: () => {
        console.log('Quiz topics fetch completed');
      }
    });
  }

  takeQuiz(quiz_topic_name : string,quiz_topic_id:string) {
    sessionStorage.setItem('currentQuizId', quiz_topic_id);
    this.router.navigate([`/${this.categoryName}/${this.subCategoryName}/${quiz_topic_name}/takequiz`]);
  }

  ngOnDestroy() {
    if (this.searchSubscription) {
      this.searchSubscription.unsubscribe();
    }
  }
}

